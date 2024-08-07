# Golang Template Backend For Product and Project
[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=bugs)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=DobbyAkhmadi_Golang-Backend-Base-Project&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=DobbyAkhmadi_Golang-Backend-Base-Project)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/powerman/go-service-example)](https://pkg.go.dev/github.com/powerman/go-service-example)
[![Project Layout](https://img.shields.io/badge/Standard%20Go-Project%20Layout-informational)](https://github.com/golang-standards/project-layout)
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
    - [Structure of Go packages](#structure-of-go-packages)
    - [Features](#features)
- [Development](#development)
    - [Requirements](#requirements)
    - [Setup](#setup)
    - [Usage](#usage)
- [Run](#run)
    - [Docker](#docker)
    - [Source](#source)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Structure of Go packages

##### this application has implement service repository pattern

- `cmd/*` - main application(s)
  - `/api/*` - list api routes
  - `/documentation/` - swagger routes
- `config/` - Configuration all settings for your microservices.
- `deployments/` - Holds deployment-related files, such as Dockerfiles and Kubernetes YAML files, for deploying your microservices.
- `docs/*` - documentation swagger
- `internal/app/{AppName}`  - Contains business logic, including use cases, entities, repositories, and services.
    - `handlers/` - list handlers
    - `models/` - list models
    - `repository/` - list repository
    - `routes/` - list routes
    - `service/` - list services
- `pkg/*` - This directory contains packages that can be shared between microservices if needed. These packages should contain reusable code and functionalities.
- `platform/*` - database migration , cache
- `scripts/*` Utility scripts for tasks like database initialization, building, or deploying the microservices.
- `log/*` - logging application
- `third_party/*` - list 3rd party communication

### Features

- ✅ Using Service Repository Pattern
- ✅ Using SonarCloud For Static Code Analysis
- ✅ Using Basic Auth JWT/Oauth(Optional) RBAC (cashbin)
- ✅ Using Basic CRUD
- ✅ Using Reverse Proxy Service Discovery Nginx/Traefik/Supervisord
- ✅ Using API Resources Naming Standard https://restfulapi.net/resource-naming/
- ✅ Using Graceful shutdown support.
- ✅ Using OpenTelemetry for collection Distributed Tracing with using Jaeger and Zipkin
- ✅ Using OpenTelemetry for collection Metrics with using Prometheus and Grafana
- ✅ Using Zap and structured logging
- ✅ Using Viper for configuration management
- ✅ Using docker and docker-compose for deployment
- ✅ Using Domain Driven Design in some of services
- ✅ Using MiniKube and Helm (monitor) Kubernetes for deployment
- ✅ Using RabbitMQ / Kafka for message queueing.
- ✅ Using Testify for unit test
- ✅ Using Air for Live Reload
- ✅ Using Circuit Breaker(gobreaker,Heimdall)
- ✅ Using Google Cloud Platform/Cloudinary for uploading image/video and can do streaming
- ✅ Using Swagger for API Documentation
