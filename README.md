# Dockter

A tool to deploy docker containers for frontend and backend servers in distributed cloud environments. The types of node in Dockter are:
<br/> **Master:** Node Responsible for all the orchestration and communication.
<br/> **Worker:** Nodes where the containers are deployed
<br/> **Builder:** Converts code to image and stores to registery

[![Watch the video](https://raw.githubusercontent.com/ayushmd/Dockter/blob/main/assets/demo.mp4)]

### Features

1. Deploy containers
2. Auto heal Dead Containers
3. Multi-node setup, which are independent
   of cloud environment.
4. Health checkups for Nodes in cluster.
5. Custom load balancer based on health and stats of container.
6. Dynamic Domain resolution for containers.
7. Replica for a docker container.
8. Ability to SSH and develop in container.

### Tech Stack

1. Primarly written in **Go lang**
2. **SQLite**: As a embeded db for dns storage.
3. **Kafka**: For storing deployment requests
4. **GRPC**: For communication between nodes.
5. **Docker Engine**: Used as programatical low level interface for deploying containers.
6. **Docker Hub**: For storing images of containers.

### Steps to Start a cluster

1. Create a auth token for others to join

```
go run . --generatetoken
```

2. To spin up Master nodes

```
go run . --state MASTER
```

3. To spin up Builder node

```
go run . --state BUILDER --join <auth_token>
```

3. To spin up Worker node

```
go run . --state WORKER --join <auth_token>
```