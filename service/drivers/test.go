package drivers

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/models"
)

// InsertTest will insert a new models.Test object into the postgres DB
func InsertTest(pgPool *pgx.ConnPool, test *models.Test) error {
	conn, err := pgPool.Acquire()
	if err != nil {
		return err
	}
	defer pgPool.Release(conn)

	exec, err := conn.Exec(
		"INSERT INTO tests (summary, outcome, analysis, resolution, doc) VALUES ($1, $2, $3, $4, $5)",
		test.Summary,
		test.Outcome,
		test.Analysis,
		test.Resolution,
		test.Doc,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return fmt.Errorf("rows affected: %d != 1", exec.RowsAffected())
	}

	return nil
}

// UpdateTest will update an existing test in the postgres DB by ID
func UpdateTest(pgPool *pgx.ConnPool, test *models.Test) error {
	conn, err := pgPool.Acquire()
	if err != nil {
		return err
	}
	defer pgPool.Release(conn)

	if err != nil {
		return err
	}
	exec, err := conn.Exec(
		"UPDATE tests SET summary=$1, outcome=$2, analysis=$3, resolution=$4, doc=$5 WHERE id=$6",
		test.Summary,
		test.Outcome,
		test.Analysis,
		test.Resolution,
		test.Doc,
		test.ID,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return fmt.Errorf("rows affected: %d != 1", exec.RowsAffected())
	}

	return nil
}

// SelectTests will take in a query that returns rows that are in the models.Test schema, deserialize them, and return
// pointers to the models.
// args will be passed down to Conn.query
func SelectTests(pgPool *pgx.ConnPool, query string, args ...any) ([]*models.Test, error) {
	conn, err := pgPool.Acquire()
	if err != nil {
		return nil, err
	}
	defer pgPool.Release(conn)
	var tests []*models.Test

	rows, err := conn.Query(query, args)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		test := &models.Test{}
		err = rows.Scan(&test.ID, &test.Summary, &test.Outcome, &test.Analysis, &test.Resolution, &test.Doc)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}
