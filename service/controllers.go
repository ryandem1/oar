package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"net/http"
	"strconv"
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

// PatchTests will perform a patch (partial update) operation on a batch of tests identified by a base64 test query
// string obtained from the /query endpoint.
// PatchTests will respond with a http.StatusNotModified (304) status code if it does not modify a single test.
// PatchTests will respond with a http.StatusOK (200) status code if it modifies at least 1 test.
//
// Note that if an error occurs in the middle of updating a batch of tests, it will result in some tests in the batch
// being updated, while others are not.
func (tc *TestController) PatchTests(c *gin.Context) {
	var query TestQuery

	encodedQuery := c.DefaultQuery("query", "null")
	if encodedQuery == "null" {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(errors.New("must pass a query parameter")))
		return
	}

	err := decodeFromBase64(&query, encodedQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	queryResult, err := QueryTest(tc.DBPool, &query, 250, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	testPatch, err := DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if queryResult.Count == 0 {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}

	for _, test := range queryResult.Tests {
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
	}
	c.Status(http.StatusOK)
}

// DeleteTests takes in a TestQuery and will delete all the query results
// DeleteTests will respond with a http.StatusNotModified (304) status code if it does not delete a single test.
// DeleteTests will respond with a http.StatusOK (200) status code if it deletes at least 1 test.
func (tc *TestController) DeleteTests(c *gin.Context) {
	var query TestQuery

	encodedQuery := c.DefaultQuery("query", "null")
	if encodedQuery == "null" {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(errors.New("must pass a query parameter")))
		return
	}

	err := decodeFromBase64(&query, encodedQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	queryResult, err := QueryTest(tc.DBPool, &query, 250, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	var testIDsToDelete []uint64
	for _, testToDelete := range queryResult.Tests {
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

	encodedQuery := c.DefaultQuery("query", "null")
	if encodedQuery != "null" {
		err = decodeFromBase64(&query, encodedQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
			return
		}
	}

	queryResult, err := QueryTest(tc.DBPool, &query, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}
	c.JSON(200, queryResult)
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
