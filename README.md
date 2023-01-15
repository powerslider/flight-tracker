# Flight Tracker API

## Overview

There are over 100,000 flights a day, with millions of people and cargo being transferred around the world.
With so many people and different carrier/agency groups, it can be hard to track where a person might be.
In order to determine the flight path of a person, we must sort through all of their flight records.

Given a list of flight routes this service determines the start and end airport of the entire flight path.

Examples:

```
[["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
[["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

```

## Getting Started

### Local Setup

**Step 0.** Install [pre-commit](https://pre-commit.com/):

```shell
pip install pre-commit

# For macOS users.
brew install pre-commit
```

Then run `pre-commit install` to setup git hook scripts.
Used hooks can be found [here](.pre-commit-config.yaml).

______________________________________________________________________

NOTE

> `pre-commit` aids in running checks (end of file fixing,
> markdown linting, go linting, runs go tests, json validation, etc.)
> before you perform your git commits.

______________________________________________________________________

**Step 1.** Install external tooling (golangci-lint, etc.):

```shell script
make install
```

**Step 2.** Setup project for local testing (code lint, runs tests, builds all needed binaries):

```shell script
make all
```

______________________________________________________________________

NOTE

> All binaries can be found in `<project_root>/bin` directory.
> Use `make clean` to delete old binaries.

______________________________________________________________________

**Step 3.** Run server:

```shell
make run-server
```

and open http://localhost:8080/swagger/index.html to open Swagger UI.

**Step 4.** Run following example query or just explore the API via Swagger UI

```shell
curl -X 'POST' \
  'http://localhost:8080/calculate' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  [
    "A","B"
  ],
  [
    "B","C"
  ]
]'
```

______________________________________________________________________

NOTE

> Check [Makefile](Makefile) for other useful commands.

______________________________________________________________________

### Docker-compose Setup

**Step 1.** Run `docker-compose` to build and run the application as a Docker container:

```shell script
docker-compose up -d
```

and open http://0.0.0.0:8080/swagger/index.html to open Swagger UI.

## API Specification

```shell
@Description Trace start and end airport given a list of flight routes.

POST /calculate

Content-Type: application/json
Accept: application/json
```

For more info check [Swagger UI](http://localhost:8080/swagger/index.html) after
running the server.

______________________________________________________________________

NOTE

> After modification of the API remember to run `make docs` is order to re-generate
> the API documentation so that Swagger UI visualizes the up-to-date changes.

______________________________________________________________________

## License

[MIT](LICENSE)
