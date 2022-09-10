.open alive.db

INSERT INTO agent(location,geohash,ISP) VALUES("Jakarta", "qqguzgberuhd1","Indiehome");
INSERT INTO agent(location,geohash,ISP) VALUES("Jakarta", "qqguzgberuhd1","Indosat");
INSERT INTO agent(location,geohash,ISP) VALUES("Jakarta", "qqguzgberuhd1","Telkomsel");
INSERT INTO agent(location,geohash,ISP) VALUES("Jakarta", "qqguzgberuhd1","XL");

INSERT INTO agent(location,geohash,ISP) VALUES("Surabaya", "qw8ntwzd4j4mj","Indiehome");
INSERT INTO agent(location,geohash,ISP) VALUES("Surabaya", "qw8ntwzd4j4mj","Indosat");
INSERT INTO agent(location,geohash,ISP) VALUES("Surabaya", "qw8ntwzd4j4mj","Telkomsel");
INSERT INTO agent(location,geohash,ISP) VALUES("Surabaya", "qw8ntwzd4j4mj","XL");

INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT",200,1);
INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT",200,1);
INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT",200,1);
INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT",200,1);
INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT00",200,1);
INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES("DESC","NAME","DOMAIN","ENDPOINT","METHOD","PROTOCOL","PERIOD_IN_CRON","BODY","HEADER",
    "AGENT00",200,1);