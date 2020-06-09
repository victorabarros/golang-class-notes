# Advanced Cloud Native Go

https://learning.oreilly.com/videos/advanced-cloud-native/9781787286238

> Published br Packat
> By Mario Leanerder Reimer

## Section 1

### Frameworks

> Anatomy of a Cloud Native Application
> Basic building blocks
> key functions/techonologies.

Basic:
Microsservices Chassis         |
Service Client                 |
Service Discovery              | consul, dns, coredns, eureka, netflix
Configuration And Coordination | Consul, etcd
Diagnosability and Monitoring  | grafana, prometheus, kibana
Api gateway                    | ngnix, fabio, kong

github.com/cncf/landscape

### Go framework and libaries

Individual Libraries:

- grpc/grpc-go
- afex/hytrix-og Netflix's latency and fault tolerance
- armon/go-metrics
- spacemonkeygo/monkit process data collection, metrucs, monitorins, instrumentation and tracing
- sirpsen/logrus
- go-kit/kit standar libary for microsservices
- micro/go-micro pluggaberpc framework
- nytimes/gizmo toolkit
- gorillatookit web tool
- gin-gonic/gin http web framwork

### Go with Gin Gonic Framework

./gin-web/main.go

### Containerization and composition with Docker

Write Dockerfile and docker-compose

### Orchestration With Kubernetes

./kubernets/*.yml

## Section 2

### Service Discovery and Configuration

Using Consul

- Start and Running
- Registering services
- Lookup services

./Discovery/Consul/*

You can use consul as a rest api to add services and set dns to them localy.
The Consul monitor the services healths.

### Consul as a Central configuration module

### Implement Microservice Registration With Consul

### Implement Microservice Lookup With Consul

### Implement Microservice Discovery and Configuration with Kubernetes

> .yamls and kuctl comands. See logs on minikube ui.

## Section 3

### Microservice Communication Patterns

Communication System:

- **Request/Response**: Client wait the response from server.
- **Push**            : One direction only from server to client.
- **Peer-to-Peer**    :

Messaging:

- **Message Passing**       : tradicional 1 to 1
- **WorkQueue**             : 1 to n sharing queue
- **Remote Procedure Call** : 1 to n with n queues
- **Publish/Subscribe**     : consumer and publihser communicate each other with tow queues

Payload formats:

- XML
- JSON
- Binary

### Implement Sync RPC calls with Binary Protocols

### Using Circuit Breaker for Resilient Communication
