# Golang Template Backend

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

- `cmd/*` - main application(s)
  - `/api/*` - list api routes
  - `/doc/` - swagger routes
- `config/` - configuration(s) (default values, env, flags) for application(s) subcommands and tests
- `deployments/` - IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).
- `docs/*` - documentation swagger
- `internal/appfolder/`  -define interfaces and implements business-logic
    - `handlers/` - list handlers controller
    - `models/` - list models
    - `repository/` - list repository
    - `routes/` - list routes
- `pkg/*` - helper packages 3rd party library, not related to architecture and business-logic
- `test/*` - list unit testing
- `third_party/*` - list 3rd party

### Features

- [X] Project structure (mostly) follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- [X] Easily testable code (thanks to The Clean Architecture).
- [X] Avoids (and resists to) using global objects (to make it possible to
  embed such microservices into modular monolith).
- [X] CLI subcommands support using [cobra](https://github.com/spf13/cobra).
- [X] Graceful shutdown support.
- [X] Configuration defaults can be overwritten by env vars and flags.