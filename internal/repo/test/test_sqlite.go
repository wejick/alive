package test

import (
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

func (T *TestSqlite) AddTest(modelTest.Test) (err error) {

	return
}

func (T *TestSqlite) GetTotalTest() (total int64, err error) {
	return
}
