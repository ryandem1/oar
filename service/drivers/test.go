package drivers

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
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
		"INSERT INTO OAR_TESTS (summary, outcome, analysis, resolution, doc) VALUES ($1, $2, $3, $4, $5)",
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
		"UPDATE OAR_TESTS SET summary=$1, outcome=$2, analysis=$3, resolution=$4, doc=$5 WHERE id=$6",
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

	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		test := &models.Test{}
		err = rows.Scan(
			&test.ID,
			&test.Summary,
			&test.Outcome,
			&test.Analysis,
			&test.Resolution,
			&test.Created,
			&test.Modified,
			&test.Doc,
		)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}

// DeleteTests will take in a slice of test IDs and attempt to delete all tests with those IDs. Will return the amount
// of rows deleted and any error that occurred. If an error occurred, it will return -1 rows deleted, which is invalid.
func DeleteTests(pgPool *pgx.ConnPool, testIDs []int64) (int64, error) {
	conn, err := pgPool.Acquire()
	if err != nil {
		return -1, err
	}
	defer pgPool.Release(conn)

	pgTestIDs := &pgtype.Int8Array{}
	err = pgTestIDs.Set(testIDs)
	if err != nil {
		return -1, err
	}

	exec, err := conn.Exec("DELETE FROM OAR_TESTS WHERE ID = ANY($1)", pgTestIDs)
	if err != nil {
		return -1, err
	}

	return exec.RowsAffected(), nil
}
