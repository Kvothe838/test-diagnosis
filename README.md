# Diagnosis Platform

 This is a technical test for the hiring process of a backend developer position. Requirements are no longer available.

 ## Prerequisites

 - Install Golang: https://go.dev/doc/install
 
 - Install Makefile: https://medium.com/@samsorrahman/how-to-run-a-makefile-in-windows-b4d115d7c516
 
 - Install golangci-lint: https://golangci-lint.run/

 - Create local-env/config.yaml, add same env variables than local-env/config-example.yaml and set their values according to your Postgres connection.

 - Execute backup for Postgres db to have all the schema and examples already saved. Use file under internal/database/scripts.

 ## Run

 To run the backend, just execute on the root of the project the following command:

 ```
make run
```

## Test

To run the unit tests, just execute on the root of the project the following command:

```
make test
```

## Lint

To run lint (it help us to write better code), just execute on the root of the project the following command:

```
make lint
```

## Documentation

Postman and Swagger documentation is under /docs.

Also, you can check Swagger API here: https://app.swaggerhub.com/apis/Kvothe838/TestDiagnosis/1.0.0

