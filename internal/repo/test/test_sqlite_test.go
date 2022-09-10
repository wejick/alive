package test

import (
	"database/sql"
	"reflect"
	"strconv"
	"testing"

	model "github.com/wejick/alive/internal/model"
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
		agent  string
		offset int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantTestlist []model.Test
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
			wantTestlist: []model.Test{
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
			wantTestlist: []model.Test{
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
			wantTestlist: []model.Test{
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
			wantTestlist: []model.Test{
				{ID: 3, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
				{ID: 4, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
		},
		{
			name: "select agent00 rows = 1",
			fields: fields{
				db: sqldb,
			},
			args: args{
				rows:  1,
				agent: "AGENT00",
			},
			wantTestlist: []model.Test{
				{ID: 5, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT00", ExpectedStatusCode: 200, Status: 1},
			},
		},
		{
			name: "select agent00 rows = 1 offset=1",
			fields: fields{
				db: sqldb,
			},
			args: args{
				rows:   1,
				agent:  "AGENT00",
				offset: 1,
			},
			wantTestlist: []model.Test{
				{ID: 6, Name: "NAME", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT00", ExpectedStatusCode: 200, Status: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T := &TestSqlite{
				db: tt.fields.db,
			}
			gotTestlist, err := T.GetTest(tt.args.IDs, tt.args.agent, tt.args.rows, tt.args.offset)
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

func TestTestSqlite_GetTotalTest(t *testing.T) {
	sqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return
	}
	defer sqldb.Close()

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name      string
		fields    fields
		wantTotal int64
		wantErr   bool
	}{
		{
			name:      "get total",
			fields:    fields{db: sqldb},
			wantTotal: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T := &TestSqlite{
				db: tt.fields.db,
			}
			gotTotal, err := T.GetTotalTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TestSqlite.GetTotalTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("TestSqlite.GetTotalTest() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func TestTestSqlite_AddTest(t *testing.T) {
	sqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return
	}
	defer sqldb.Close()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		test model.Test
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []model.Test
		wantErr    bool
	}{
		{
			name:   "add simple",
			fields: fields{db: sqldb},
			args: args{
				test: model.Test{Name: "Simple test", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
					Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1},
			},
			wantResult: []model.Test{{Name: "Simple test", Desc: "DESC", Domain: "DOMAIN", Endpoint: "ENDPOINT", Method: "METHOD", Protocol: "PROTOCOL", PeriodInCron: "PERIOD_IN_CRON",
				Body: "BODY", Header: "HEADER", Agent: "AGENT", ExpectedStatusCode: 200, Status: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T := &TestSqlite{
				db: tt.fields.db,
			}
			if err := T.AddTest(tt.args.test); (err != nil) != tt.wantErr {
				t.Errorf("TestSqlite.AddTest() error = %v, wantErr %v", err, tt.wantErr)
			}

			lastID, _ := T.GetTotalTest()
			ID := strconv.Itoa(int(lastID))
			result, _ := T.GetTest([]string{ID}, "", 0, 0)
			tt.wantResult[0].ID = lastID
			if !reflect.DeepEqual(result, tt.wantResult) {
				t.Errorf("TestSqlite.GetTest() = %v, want %v", result, tt.wantResult)
			}
		})
	}
}
