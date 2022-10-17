# ALIVE
Periodically test endpoint in location and ISP aware

# Overview 
There are 2 part that make Alive, `alive-agent` run the test cases while `alive-server` contain information including test cases, agent information and many other relevant configurations. Alive cluster will consist of several `alive-agent` deployed on various configuration of location and ISP,  one `alive-server` coordinating all the agents. 

`alive-agent` will run test cases and expose the result as prometheus formated metrics ready to be scrapped, by default alive is using fluent-bit to scrape and store the data into prometheus via prometheus_remote_write.

![architecture overview of alive cluster](/doc/overview.png)

# TODO
- [ ] Package for agent with auto update support, probably using snap
- [ ] Automatic provisioning of agent and it's host
- [ ] Account / agent security, probably first combination of simple password and jwt
- [ ] API remove agent
- [ ] multitenancy of test
- [ ] think better test runner
- [ ] think better test result storage and reporting
- [ ] working test report dashboard
- [ ] make the test report dashboard multitenant
- [ ] think about tenant / user of test