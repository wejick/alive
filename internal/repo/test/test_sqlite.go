package test

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	modelTest "github.com/wejick/alive/internal/model/test"
	_ "modernc.org/sqlite"
)

// TestSqlite ...
type TestSqlite struct {
	db *sql.DB
}

func NewSqlite(sqliteDB *sql.DB) (test *TestSqlite) {
	test = &TestSqlite{db: sqliteDB}

	return
}

// GetTest get test configuration by id or by page.
// if IDs are provided, rows and offset will be ignored
func (T *TestSqlite) GetTest(IDs []string, rows, offset int) (testlist []modelTest.Test, err error) {
	query := ""
	if len(IDs) > 0 && IDs[0] != "" {
		ids := strings.Join(IDs[:], ",")
		query = fmt.Sprintf("SELECT * FROM test WHERE id IN (%s)", ids)
	} else {
		query = fmt.Sprintf("SELECT * FROM test LIMIT %d OFFSET %d", rows, offset)
	}

	dbRows, err := T.db.Query(query)
	if err != nil {
		return
	}
	defer dbRows.Close()

	for dbRows.Next() {
		item := modelTest.Test{}
		dbRows.Scan(&item.ID, &item.Desc, &item.Name, &item.Domain, &item.Endpoint, &item.Method, &item.Protocol, &item.PeriodInCron, &item.Body,
			&item.Header, &item.Agent, &item.ExpectedStatusCode, &item.Status)

		testlist = append(testlist, item)
	}

	return
}

func (T *TestSqlite) AddTest(test modelTest.Test) (err error) {
	query := "INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
	tx, err := T.db.BeginTx(context.Background(), nil)
	if err != nil {
		return
	}
	_, err = tx.Exec(query, test.Desc, test.Name, test.Domain, test.Endpoint, test.Method, test.Protocol, test.PeriodInCron, test.Body,
		test.Header, test.Agent, test.ExpectedStatusCode, test.Status)
	tx.Commit()

	return
}

func (T *TestSqlite) GetTotalTest() (total int64, err error) {
	row := T.db.QueryRow("SELECT COUNT(id) FROM test")
	err = row.Scan(&total)

	return
}