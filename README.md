# craas-go: Go SDK for Container Registry Service

[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/craas-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/craas-go)](https://goreportcard.com/report/github.com/selectel/craas-go)
[![Coverage Status](https://coveralls.io/repos/github/selectel/craas-go/badge.svg?branch=master)](https://coveralls.io/github/selectel/craas-go?branch=master)

Package craas-go provides Go SDK to work with the Selectel Container Registry Service.

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/craas-go/).

## What this library is capable of

You can use this library to work with the following objects of the Selectel Container Registry Service:

* [token](https://pkg.go.dev/github.com/selectel/craas-go/pkg/v1/token)
* [registry](https://pkg.go.dev/github.com/selectel/craas-go/pkg/v1/registry)
* [repository](https://pkg.go.dev/github.com/selectel/craas-go/pkg/v1/repository)
* [garbage-collection](https://pkg.go.dev/github.com/selectel/craas-go/pkg/v1/gc)

## Getting started

### Installation

You can install needed `craas-go` packages via `go get` command:

```bash
go get github.com/selectel/craas-go/pkg/v1/registry
```

### Authentication

To work with the Selectel Container Registry API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create a project in Selectel Cloud Platform [projects](https://my.selectel.ru/vpc/projects).
* Retrieve a token for your project via API or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

### Endpoints

Selectel Container Registry Service currently has the following API endpoints:

| URL                           |
|-------------------------------|
| https://cr.selcloud.ru/api/v1 |

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/selectel/craas-go/pkg"
	"github.com/selectel/craas-go/pkg/v1/registry"
	"github.com/selectel/craas-go/pkg/v1/repository"
)

func main() {
	// Token to work with Selectel Cloud project.
	token := "gAAAAABeVNzu-..."

	// CRaaS endpoint to work with.
	endpoint := "https://cr.selcloud.ru/api/v1"

	// Create a new CRaaS client.
	crClient := v1.NewCRaaSClientV1(token, endpoint)

	// Prepare empty context.
	ctx := context.Background()

	// Create a new registry.
	createdRegistry, _, err := registry.Create(ctx, crClient, "my-registry")
	if err != nil {
		log.Fatal(err)
	}

	// Print the registry fields.
	fmt.Printf("Created registry: %+v", createdRegistry)

	// Get a list of registry repositories.
	repositories, _, err := repository.ListRepositories(ctx, crClient, createdRegistry.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Print the repository fields.
	for _, repo := range repositories {
		fmt.Printf("Repository: %+v", repo)
	}
}
```