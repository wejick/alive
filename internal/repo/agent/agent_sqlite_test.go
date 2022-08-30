package agent

import (
	"database/sql"
	"reflect"
	"testing"

	modelAgent "github.com/wejick/alive/internal/model/agent"
	_ "modernc.org/sqlite"
)

func TestAgentSqlite_GetAgents(t *testing.T) {
	sqldb, err := sql.Open("sqlite", "./alive.db")
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
		},
		{
			name:   "select 1 , 2",
			fields: fields{db: sqldb},
			args:   args{agentIDs: []string{"1", "2"}},
		},
		{
			name:   "select",
			fields: fields{db: sqldb},
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
	sqldb, err := sql.Open("sqlite", "./alive.db")
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
