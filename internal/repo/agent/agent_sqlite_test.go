package agent

import (
	"database/sql"
	"reflect"
	"testing"

	modelAgent "github.com/wejick/alive/internal/model/agent"
	_ "modernc.org/sqlite"
)

const dbpath = "../../../alive.db"

func TestAgentSqlite_GetAgents(t *testing.T) {
	sqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return
	}
	defer sqldb.Close()

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
		wantAgentList []modelAgent.Agent
	}{
		{
			name:   "select 1",
			fields: fields{db: sqldb},
			args:   args{agentIDs: []string{"1"}},
			wantAgentList: []modelAgent.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome"},
			},
		},
		{
			name:   "select 1 , 2",
			fields: fields{db: sqldb},
			args:   args{agentIDs: []string{"1", "2"}},
			wantAgentList: []modelAgent.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome"},
				{ID: 2, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indosat"},
			},
		},
		{
			name:   "select",
			fields: fields{db: sqldb},
			wantAgentList: []modelAgent.Agent{
				{ID: 1, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indiehome"},
				{ID: 2, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Indosat"},
				{ID: 3, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "Telkomsel"},
				{ID: 4, Location: "Jakarta", GeoHash: "qqguzgberuhd1", ISP: "XL"},
				{ID: 5, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Indiehome"},
				{ID: 6, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Indosat"},
				{ID: 7, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "Telkomsel"},
				{ID: 8, Location: "Surabaya", GeoHash: "qw8ntwzd4j4mj", ISP: "XL"},
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
	sqldb, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return
	}
	defer sqldb.Close()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		agent modelAgent.Agent
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
			args: args{agent: modelAgent.Agent{
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
