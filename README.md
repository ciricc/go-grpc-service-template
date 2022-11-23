# Go GRPC Service template with VS Code

## Makefile

`build` - Build docker image
`genproto` - Generate code from protos int `protos/*` folder

## Configuration

Config file - `config.yml` (YAML)

Use `config` package to change configuration propterties and add auto-filling it with values from ENVIRONMENT ([cleaneanv](https://github.com/ilyakaznacheev/cleanenv) implementation).

## Logging

Using logger from [logrus](https://github.com/sirupsen/logrus)

## Pattern

This template uses dependency injection with [do](https://github.com/samber/do) implementation. Use package `di` for creating new constructors (providers) for any structures which you want to reuse many times with same initialization.

## Run

This template created for developing ing VS Code Remote Containers. So, you can run it int dev mode:

```sh
go run cmd/printer/main.go
```

Or you can fristly build your Dockerfile and after this run it.

```sh
make build && docker run ciricc/printer
```