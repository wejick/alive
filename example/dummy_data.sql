.open alive.db

INSERT INTO agent(location,geohash,ISP,status) VALUES('Jakarta', 'qqguzgberuhd1','Indiehome',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Jakarta', 'qqguzgberuhd1','Indosat',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Jakarta', 'qqguzgberuhd1','Telkomsel',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Jakarta', 'qqguzgberuhd1','XL',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Surabaya', 'qw8ntwzd4j4mj','Indiehome',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Surabaya', 'qw8ntwzd4j4mj','Indosat',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Surabaya', 'qw8ntwzd4j4mj','Telkomsel',1);
INSERT INTO agent(location,geohash,ISP,status) VALUES('Surabaya', 'qw8ntwzd4j4mj','XL',1);

INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES('Testing google','Get Google.com','google.com','/','GET','https','@every 10s','','',
    '1',200,1);

INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES('Testing yahoo','Get yahoo.com','yahoo.com','/','GET','https','@every 10s','','',
    '1',200,1);

INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES('Testing google','Get Google.com','google.com','/','GET','https','@every 10s','','',
    '2',200,1);

INSERT INTO test(desc,name,domain,endpoint,method,protocol,period_in_cron,body,header,agent,expected_status_code,status)
    VALUES('Testing yahoo','Get yahoo.com','yahoo.com','/','GET','https','@every 10s','','',
    '2',200,1);

INSERT INTO agent_ping(agent_id,last_ping_time) values(1,1665127654);
INSERT INTO agent_ping(agent_id,last_ping_time) values(2,1665127660);