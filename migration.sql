.open alive.db
CREATE TABLE IF NOT EXISTS agent (
    id INTEGER PRIMARY KEY,
    location TEXT NOT NULL,
    geohash TEXT NOT NULL,
    ISP TEXT NOT NULL,
    status INTEGER
);

CREATE TABLE IF NOT EXISTS test (
    id INTEGER PRIMARY KEY,
    desc TEXT NOT NULL,
    name TEXT NOT NULL,
    domain TEXT NOT NULL,
    endpoint TEXT NOT NULL,
    method TEXT NOT NULL,
    protocol TEXT NOT NULL,
    period_in_cron TEXT NOT NULL,
    body TEXT NOT NULL,
    header TEXT NOT NULL,
    agent TEXT NOT NULL,
    expected_status_code INTEGER NOT NULL,
    status INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS agent_ping (
    id INTEGER PRIMARY KEY,
    agent_id INTEGER UNIQUE,
    last_ping_time INTEGER,
    FOREIGN KEY (agent_id) REFERENCES agent(id)
);