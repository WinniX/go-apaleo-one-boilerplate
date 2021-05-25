## How to update an existing API client

### Booking

1. Download the latest version of `swagger.json` here https://api.apaleo.com/swagger/booking-v1/swagger.json and save it as `booking-v1.json`
2. Generate the client

```bash
docker run --rm \
  -v ${PWD}:/local \
  -u 1000:1000 \
  openapitools/openapi-generator-cli generate \
  -i /local/booking-v1.json \
  -g go \
  -o /local/bookingclient \
  --additional-properties=packageName=bookingclient \
  --git-repo-id go-apaleo-one-boilerplate/api/bookingclient \
  --git-user-id winnix
```

## Integration

1. Download the latest version of `swagger.json` here https://integration.apaleo.com/swagger/integration-v1/swagger.json and save it as `integration-v1.json`
2. Generate the client

```bash
docker run --rm \
  -v ${PWD}:/local \
  -u 1000:1000 \
  openapitools/openapi-generator-cli generate \
  -i /local/integration-v1.json \
  -g go \
  -o /local/integrationclient \
  --additional-properties=packageName=integrationclient \
  --git-repo-id go-apaleo-one-boilerplate/api/integrationclient \
  --git-user-id winnix
```

## How to add a new apaleo API client:

1. Generate the client (see above)
2. Add newly generated client to the project

```bash
go mod edit -replace github.com/winnix/go-apaleo-one-boilerplate/api/{API_NAME}client=./clients/{API_NAME}client
go get github.com/winnix/go-apaleo-one-boilerplate/api/{API_NAME}client
```
