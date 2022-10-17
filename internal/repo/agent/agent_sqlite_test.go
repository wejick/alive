package agent

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"

	model "github.com/wejick/alive/internal/model"
	_ "modernc.org/sqlite"
)

const dbpath = "../../../alive.db?_pragma=foreign_keys(1)&_pragma=busy_timeout(1000)"

var sqldb *sql.DB

func init() {
	localSqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		fmt.Print(err)
		return
	}
	sqldb = localSqldb
}

func TestMain(m *testing.M) {
	defer sqldb.Close()

	flag.Parse()

	os.Exit(m.Run())
}

func TestAgentSqlite_GetAgents(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		agentIDs []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantAgentList []model.Agent
	}{
		{
			name:   "select 1",
			fields: fields{db: sqldb},
			args:   args{agentIDs: []string{"1"}},
			wantAgentList: []model.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome", StatusText: "ACTIVE", Status: 1},
			},
		},
		{
			name:   "select 1 , 2",
			fields: fields{db: sqldb},
			args:   args{agentIDs: []string{"1", "2"}},
			wantAgentList: []model.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome", StatusText: "ACTIVE", Status: 1},
				{ID: 2, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indosat", StatusText: "ACTIVE", Status: 1},
			},
		},
		{
			name:   "select",
			fields: fields{db: sqldb},
			wantAgentList: []model.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome", StatusText: "ACTIVE", Status: 1},
				{ID: 2, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indosat", StatusText: "ACTIVE", Status: 1},
				{ID: 3, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Telkomsel", StatusText: "ACTIVE", Status: 1},
				{ID: 4, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "XL", StatusText: "ACTIVE", Status: 1},
				{ID: 5, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Indiehome", StatusText: "ACTIVE", Status: 1},
				{ID: 6, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Indosat", StatusText: "ACTIVE", Status: 1},
				{ID: 7, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Telkomsel", StatusText: "ACTIVE", Status: 1},
				{ID: 8, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "XL", StatusText: "ACTIVE", Status: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &AgentSqlite{
				db: tt.fields.db,
			}
			if gotAgentList := A.GetAgents(tt.args.agentIDs); !reflect.DeepEqual(gotAgentList, tt.wantAgentList) {
				t.Errorf("AgentSqlite.GetAgents() = %v, want %v", gotAgentList, tt.wantAgentList)
			}
		})
	}
}

func TestAgentSqlite_AddAgent(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		agent model.Agent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "select 1",
			fields: fields{db: sqldb},
			args: args{agent: model.Agent{
				Location: "kediri",
				GeoHash:  "yyyykdr",
				ISP:      "kediri net",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &AgentSqlite{
				db: tt.fields.db,
			}
			if err := A.AddAgent(tt.args.agent); (err != nil) != tt.wantErr {
				t.Errorf("AgentSqlite.AddAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
