# Demo API Gateway
-----------

        _______________________________________
       /                                       \
      |              API Gateway                |
      |                (Gin)                    |
      |                Port: 8080                |
      |                                         |
      |      /           |            \         |
      |     /            |             \        |
      |    v             v              v       |
      | +--------+   +----------+   +----------+|
      | |Service1|   | Service2 |   | Service3 ||
      | |(Golang)|   | (Express)|   |(Laravel)||
      | | Port: 9001| | Port: 9002| | Port: 9003||
      | +--------+   +----------+   +----------+|
       \_______________________________________/
                \          |          /
                 \         |         /
                  \        |        /
                   \       |       /
                    \      |      /
                     \     |     /
                      \    |    /
                       \   |   /
                        \  |  /
                         \ | /
                          \|/
                       +-------+
                       | Client|
                       +-------+

##### Client:
Client is a user or application that sends a request to your system.

##### API Gateway (Gin):
API Gateway acts as a single gateway that accepts all requests from Clients.
Built using Gin in Golang and runs on Port 8080.
This gateway will forward requests to one of three services based on the requested path.

##### Service1 (Golang):
The first service is built with Golang and runs on Port 9001.
Example endpoint: http://localhost:8080/service1 will be forwarded to http://localhost:9001.

##### Service2 (Express.js):
The second service is built with Express.js and runs on Port 9002.
Example endpoint: http://localhost:8080/service2 will be forwarded to http://localhost:9002.

##### Service3 (Laravel):
The third service is built with Laravel and runs on Port 9003.
Example endpoint: http://localhost:8080/service3 will be forwarded to http://localhost:9003.
For example, http://localhost:8080/service3/categories will be forwarded to http://localhost:9003/categories.

### Benefits of Using API Gateway:
- **Centralized Entry Point:** All requests go through a single point, making it easier to manage security, logging, and monitoring.
- **Abstraction:** Clients do not need to know the implementation details or location of each service.
- **Scalability:** Makes it easy to add or remove services without changing clients.

### Current Development:
**[20-01-2025]**
- Create API gateway
- Simple CRUD Laravel Product & Category Service3
- Create/Update data to DB (MySQL) using Job
- Simple API Golang Service1
- Job and Worker Manager using Gouroutine

### Further Development:

**Authentication & Authorization:** Add security mechanism in API Gateway.
**Rate Limiting:** Limit the number of requests to prevent abuse.
**Load Balancing:** Distribute requests to multiple service instances to improve performance and availability.
**Caching:** Store frequently requested responses to reduce backend service load.
**Service Integration:** Communication from one service to another
**Multiple Database Abstraction:** Able to make abstraction database technoogy
**Multiple Database Abstraction:** Able to make abstraction database technoogy
**Implement Asynq:** distributed task queue in Go
\
\
\
\
☕️☕️☕️
*PS: Research Only*
