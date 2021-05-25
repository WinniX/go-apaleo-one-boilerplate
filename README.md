# [Go] apaleo One boilerplate

Quickstart Go project template for developing an "apaleo One" app.
More info on what "apaleo One" is here https://apaleo.dev/guides/apaleo-store/

### Prerequisites

1. Based on the content of `api/env-config.yaml` add a new file e.g. `dev-config.yaml` and set your values.
2. (optional) If you made a change to the DB connection string on the previous step, make sure to adjust the environment variables inside `docker-compose.yml` accordingly.

### How to run it locally

1. Create a database

```bash
docker-compose up -d
```

2. Navigate to the `api` folder and run

```bash
go get -d ./...
go build .
```

3. Run it

```bash
./api --config=dev-config.yaml
```

or (on Windows)

```bash
api.exe --config=dev-config.yaml
```

4. Open up http://localhost:8080

## Contribution

Feel free to open an issue if you found any error or to create a pull request if want to add additional content.
