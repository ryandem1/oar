package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"net/http"
	"strconv"
	"strings"
)

// TestController will maintain a database pool for all test controllers
type TestController struct {
	DBPool *pgx.ConnPool
}

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func (tc *TestController) CreateTest(c *gin.Context) {
	test, err := DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	test.Clean()
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	testID, err := InsertTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	c.JSON(http.StatusCreated, testID)
}

// PatchTest will perform a patch (partial update) operation on an existing test if it exists. Because of the nature of
// the Test enrichment process, I imagine this will be used more than a PUT would be.
func (tc *TestController) PatchTest(c *gin.Context) {
	testPatch, err := DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if testPatch.ID == 0 {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(
			errors.New("must define an id of an existing testPatch to update")),
		)
		return
	}

	// Perform partial update on copy of existing testPatch and doc merge
	tests, err := SelectTests(tc.DBPool, "SELECT * FROM OAR_TESTS WHERE id=$1", testPatch.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	test := tests[0]
	test.Merge(testPatch)

	// Validate after update to ensure testPatch is still okay
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	// Update in DB
	err = UpdateTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	c.Status(http.StatusOK)
}

// DeleteTests is a bulk endpoint that can mark tests as deleted. These tests will no longer be visible from the UI, but
// will still remain in the DB unless manually removed just-in-case they need to be recovered.
// DeleteTests will silently ignore if the caller passes in test IDs that already don't exist.
// DeleteTests will respond with a http.StatusNotModified (304) status code if it does not delete a single test.
// DeleteTests will respond with a http.StatusOK (200) status code if it deletes at least 1 test.
func (tc *TestController) DeleteTests(c *gin.Context) {
	var testsToDelete []Test

	if err := c.BindJSON(&testsToDelete); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}
	var testIDsToDelete []uint64
	for _, testToDelete := range testsToDelete {
		testIDsToDelete = append(testIDsToDelete, testToDelete.ID)
	}

	testsDeleted, err := DeleteTests(tc.DBPool, testIDsToDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if testsDeleted < 0 {
		c.Status(http.StatusInternalServerError)
	} else if testsDeleted == 0 {
		c.Status(http.StatusNotModified)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetTests is the primary way to query for tests. It looks for a TestQuery request body and limit and offset URL
// params. Passing multiple values within an array will be treated as a logical 'OR' for querying that field. Multiple
// attributes passed in the query will be treated as logical 'AND'.
//
// Additionally, the unstructured Doc can be queried, it will partially match with the Postgres "contains (@>)"
// operator. For more information, see: https://www.postgresql.org/docs/current/functions-json.html
func (tc *TestController) GetTests(c *gin.Context) {
	var query TestQuery
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "250"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if limit > 1000 { // Maximum limit
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(errors.New("maximum allowed limit is 1000")))
		return
	}

	if err := c.BindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	// Build query
	var wheres []string // Will contain all the "WHERE" clauses
	var params []any    // These are the parameters to pass for the SQL prepared statement

	SQL := "SELECT * FROM OAR_TESTS"
	if len(query.IDs) > 0 {
		wheres = append(wheres, "ID = ANY($)")
		params = append(params, query.IDs)
	}

	if len(query.Summaries) > 0 {
		wheres = append(wheres, "SUMMARY ~* "+"'"+strings.Join(query.Summaries, "|")+"'")
	}

	if len(query.Outcomes) > 0 {
		wheres = append(wheres, "OUTCOME = ANY($)")
		params = append(params, query.Outcomes)
	}

	if len(query.Analyses) > 0 {
		wheres = append(wheres, "ANALYSIS = ANY($)")
		params = append(params, query.Analyses)
	}

	if len(query.Resolutions) > 0 {
		wheres = append(wheres, "RESOLUTION = ANY($)")
		params = append(params, query.Resolutions)
	}

	if query.CreatedBefore != nil {
		wheres = append(wheres, "CREATED < $")
		params = append(params, query.CreatedBefore)
	}

	if query.CreatedAfter != nil {
		wheres = append(wheres, "CREATED > $")
		params = append(params, query.CreatedAfter)
	}

	if query.ModifiedBefore != nil {
		wheres = append(wheres, "MODIFIED < $")
		params = append(params, query.ModifiedBefore)
	}

	if query.ModifiedAfter != nil {
		wheres = append(wheres, "MODIFIED > $")
		params = append(params, query.ModifiedAfter)
	}

	if len(query.Docs) > 0 {
		docQuery := "("
		for i, doc := range query.Docs {
			if i == 0 {
				docQuery += "doc @> '$'"
			} else {
				docQuery += " " + "OR" + " " + "doc @> '$'"
			}
			strDoc, err := json.Marshal(doc)
			if err != nil {
				c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
				return
			}
			docQuery = strings.Replace(docQuery, "$", string(strDoc), 1)
		}
		docQuery += ")"
		wheres = append(wheres, docQuery)
	}

	for i, where := range wheres {
		where = strings.Replace(where, "$", "$"+strconv.Itoa(i+1), 1) // formats string replacement params
		if i == 0 {
			SQL += " " + "WHERE" + " " + where
		} else {
			SQL += " " + "AND" + " " + where
		}
	}

	// Orders by the most recently modified tests being first
	SQL += " " + "ORDER BY MODIFIED DESC"

	// Add offset and limit
	SQL += " " + "OFFSET " + strconv.Itoa(offset) + " " + "LIMIT" + " " + strconv.Itoa(limit)

	tests, err := SelectTests(tc.DBPool, SQL, params...)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if tests == nil {
		tests = []*Test{}
	}

	response := TestQueryResponse{Count: uint64(len(tests)), Tests: tests}
	c.JSON(200, response)
}

// EncodeSearchQuery will take a TestQuery as a body and encode it in base64 to send to the GET endpoint.
// This is an intermediate step for 2 reasons:
//
// 1.) By encoding the request, it allows for a complex search query language without complex GET URL params.
// I had originally made the GET endpoint take a request body, but that is not by HTTP 1.1 spec.
//
// 2.) This will also allow search queries to be sent via url which will be helpful.
func EncodeSearchQuery(c *gin.Context) {
	var query TestQuery
	if err := c.BindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	encodedString, err := encodeToBase64(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}
	c.JSON(200, encodedString)
}
