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
