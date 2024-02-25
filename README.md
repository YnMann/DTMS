# DTMS
Development of a Distributed Task Management System (DTMS),

## Description of the Project

This project is a distributed task management system designed to efficiently distribute and complete tasks between different performers in a distributed network. The system is designed to simplify the task processing process, ensure their reliable execution and provide tools for monitoring and managing tasks in real time.

## Technology Stack

- **Gin:** A web framework for fast RESTful API development on Go.
- **gRPC + Protocol Buffers:** For effective interaction between microservices using binary data serialization and strong typing.
- **RabbitMQ:** A message broker to ensure reliable asynchronous messaging between system components.
- **MongoDB:** Document-oriented database for flexible data storage and query.
- **Docker:** A containerization platform to simplify application deployment and management.
- **ELK Stack (Elasticsearch, Logstash, Kibana):** To collect, store and visualize application logs in real time.

## Getting Started

### Preliminary Requirements

To work with the project, make sure that you have installed:

- Docker
- Docker Compose

### Installation and Launch

1. Clone the project repository:

```bash
git clone https://github.com/YnMann/DTMS
cd <project name>
docker-compose up --build
