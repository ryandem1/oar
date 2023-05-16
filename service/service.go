package main

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

// QueryTest will take a DB connection pool to the OAR DB, parse the query, apply the limit and offset and return
// the query response.
// See GetTests for more info
func QueryTest(dbPool *pgx.ConnPool, query *TestQuery, limit int, offset int) (*TestQueryResponse, error) {
	// Build query
	var wheres []string // Will contain all the "WHERE" clauses
	var params []any    // These are the parameters to pass for the SQL prepared statement

	SQL := "SELECT * FROM OAR_TESTS"
	if len(query.IDs) > 0 {
		wheres = append(wheres, "ID = ANY($)")
		params = append(params, query.IDs)
	}

	if query != nil && len(query.Summaries) > 0 {
		wheres = append(wheres, "SUMMARY ~* "+"'"+strings.Join(query.Summaries, "|")+"'")
	}

	if query != nil && len(query.Outcomes) > 0 {
		wheres = append(wheres, "OUTCOME = ANY($)")
		params = append(params, query.Outcomes)
	}

	if query != nil && len(query.Analyses) > 0 {
		wheres = append(wheres, "ANALYSIS = ANY($)")
		params = append(params, query.Analyses)
	}

	if query != nil && len(query.Resolutions) > 0 {
		wheres = append(wheres, "RESOLUTION = ANY($)")
		params = append(params, query.Resolutions)
	}

	if query != nil && query.CreatedBefore != nil {
		wheres = append(wheres, "CREATED < $")
		params = append(params, query.CreatedBefore)
	}

	if query != nil && query.CreatedAfter != nil {
		wheres = append(wheres, "CREATED > $")
		params = append(params, query.CreatedAfter)
	}

	if query != nil && query.ModifiedBefore != nil {
		wheres = append(wheres, "MODIFIED < $")
		params = append(params, query.ModifiedBefore)
	}

	if query != nil && query.ModifiedAfter != nil {
		wheres = append(wheres, "MODIFIED > $")
		params = append(params, query.ModifiedAfter)
	}

	if query != nil && len(query.Docs) > 0 {
		docQuery := "("
		for i, doc := range query.Docs {
			if i == 0 {
				docQuery += "doc @> '$'"
			} else {
				docQuery += " " + "OR" + " " + "doc @> '$'"
			}
			strDoc, err := json.Marshal(doc)
			if err != nil {
				return nil, err
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
	SQL += " " + "ORDER BY CREATED DESC"

	// Add offset and limit
	SQL += " " + "OFFSET " + strconv.Itoa(offset) + " " + "LIMIT" + " " + strconv.Itoa(limit)

	tests, err := SelectTests(dbPool, SQL, params...)
	if err != nil {
		return nil, err
	}

	if tests == nil {
		tests = []*Test{}
	}

	response := &TestQueryResponse{Count: uint64(len(tests)), Tests: tests}
	return response, nil
}
