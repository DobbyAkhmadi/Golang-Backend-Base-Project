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
- `config/` - configuration(s) (default values, env, flags) for application(s) subcommands and tests
- `deployments/` - IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).
- `docs/*` - documentation swagger
- `internal/app/`  -define interfaces and implements business-logic
    - `handlers/` - list handlers controller
    - `models/` - list models
    - `repository/` - list repository
    - `routes/` - list routes
    -  `service/` - list services
- `pkg/*` - helper packages 3rd party library, not related to architecture and business-logic
- `platform/*` - database migration , cache
- `log/*` - logging application
- `test/*` - list unit testing
- `third_party/*` - list 3rd party

### Features

- ✅ Service Repository Pattern
- ✅ Using Basic Auth JWT/Oauth(Optional)
- ✅ Using Basic CRUD
- ✅ Using Reverse Proxy Nginx/Traefik
- ✅ Using API Resources Naming Standard https://restfulapi.net/resource-naming/
- ✅ Using Graceful shutdown support.
- ✅ Using OpenTelemetry for collection Distributed Tracing with using Jaeger and Zipkin
- ✅ Using OpenTelemetry for collection Metrics with using Prometheus and Grafana
- ✅ Using Zap and structured logging
- ✅ Using Viper for configuration management
- ✅ Using docker and docker-compose for deployment
- ✅ Using Domain Driven Design in some of services
- ✅ Using MiniKube and Helm (monitor) Kubernetes for deployment
- ✅ Using Air for Live Reload
- ✅ Using Swagger for API Documentation
