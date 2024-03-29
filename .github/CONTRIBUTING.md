# Contribution guide

## Basics

Prior to create a PR please create an issue that will describe a problem.

## Project structure

Every API part should be implemented in its separate package.

Any package which implements methods to work with API uses the
following structure:

```
selectel_craas_api_object/ # Name of the directory should desrcibe an API object (cluster/...)
├── doc.go               # Documentation that will be available at the godoc.org
├── requests.go          # Methods to work with the API
├── requests_opts.go     # Structures that contain options for a POST/PATCH requests
├── schemas.go           # Structures that contain unmarshalled responses
└── testing
    ├── fixtures.go      # Tests fixtures
    └── requests_test.go # Tests for all implemented requests
```

## Tests

Please implement tests for all methods that you're creating.
You can use the fake ServiceClient, TokenID and testing HTTPServer from the `testing`
package.
