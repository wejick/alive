package test

import (
	"database/sql"
	"reflect"
	"testing"

	modelTest "github.com/wejick/alive/internal/model/test"
	_ "modernc.org/sqlite"
)

const dbpath = "../../../alive.db"

func TestTestSqlite_GetTest(t *testing.T) {

	sqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return
	}
	defer sqldb.Close()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		IDs    []string
		rows   int
		offset int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantTestlist []modelTest.Test
		wantErr      bool
	}{
		{
			name: "select 1",
			fields: fields{
				db: sqldb,
			},
			args: args{
				IDs: []string{"1"},
			},
			wantTestlist: []modelTest.Test{
				{ID: 1, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
		},
		{
			name: "select 1, 2",
			fields: fields{
				db: sqldb,
			},
			args: args{
				IDs: []string{"1", "2"},
			},
			wantTestlist: []modelTest.Test{
				{ID: 1, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
				{ID: 2, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
		},
		{
			name: "select rows = 2, offset 0",
			fields: fields{
				db: sqldb,
			},
			args: args{
				rows:   2,
				offset: 0,
			},
			wantTestlist: []modelTest.Test{
				{ID: 1, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
				{ID: 2, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
		},
		{
			name: "select rows = 2, offset 2",
			fields: fields{
				db: sqldb,
			},
			args: args{
				rows:   2,
				offset: 2,
			},
			wantTestlist: []modelTest.Test{
				{ID: 3, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
				{ID: 4, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T := &TestSqlite{
				db: tt.fields.db,
			}
			gotTestlist, err := T.GetTest(tt.args.IDs, tt.args.rows, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestSqlite.GetTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTestlist, tt.wantTestlist) {
				t.Errorf("TestSqlite.GetTest() = %v, want %v", gotTestlist, tt.wantTestlist)
			}
		})
	}
}
