# Microservices Architecture for Code Analysis and Documentation Platform

## Overview
This repository contains the microservices architecture for a platform designed to analyze codebases and generate interactive documentation. The platform is built using Docker to facilitate a microservices architecture, allowing each component to be scaled and maintained independently.

## Services
- **Auth**: Handles authentication and user management.
- **Documentation**: Manages and serves dynamically generated documentation from code analysis data.
- **Code Analysis**: Processes codebases to generate ASTs and populates a graph database for further analysis.

## Common Libraries
The `/common` directory contains shared libraries and utilities used across different services. These are packaged and versioned, allowing services to update independently while sharing core functionality.

## Architecture Decisions
- **Docker Containers**: Each service runs in its own Docker container, ensuring isolation and ease of deployment.
- **Docker Compose**: Used for defining and running multi-container Docker applications.
- **AWS EC2**: All containers are hosted on a single EC2 instance initially, allowing for easy management and cost-effectiveness.
- **AWS ECS**: Utilized for managing containers, handling deployment, scaling, and load balancing.
- **API Gateway**: Manages incoming requests and routes them to the appropriate microservice.
- **Nginx**: Acts as a reverse proxy and load balancer, handling and distributing incoming traffic to various services.

## Scalability
- Services are designed to be easily split into separate repositories if necessary, supporting both a monolithic and microservices deployment approach.
- The use of Docker and AWS ECS facilitates scaling as the load increases.

## Security
- All traffic is routed through AWS API Gateway, providing an additional layer of security and the ability to monitor and manage traffic efficiently.

## Future Considerations
- As the platform grows, each microservice can be moved to its own EC2 instance or managed via Kubernetes to handle increased load and provide high availability.

## Getting Started
To get started with this setup, clone the repository and navigate to the root directory. Use the following command to start all services:

```z
docker-compose up --build
```

For detailed service configuration, refer to individual Dockerfiles and the `docker-compose.yml` file in the root directory.
