﻿# TopDoctors Backend Challenge

 Solution for Top Doctors Backend Challenge: https://bitbucket.org/topdoctors/backend-developer/src/master/challenge.md.

 ## Prerequisites

 - Install Golang: https://go.dev/doc/install
 
 - Install Makefile: https://medium.com/@samsorrahman/how-to-run-a-makefile-in-windows-b4d115d7c516
 
 - Install golangci-lint: https://golangci-lint.run/

 - Create local-env/config.yaml, add same env variables than local-env/config-example.yaml and set their values according to your Postgres connection.

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

Also, you can check Swagger API here: https://app.swaggerhub.com/apis/Kvothe838/TopDoctorsBackendChallenge/1.0.0

