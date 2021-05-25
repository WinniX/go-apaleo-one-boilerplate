module github.com/winnix/go-apaleo-one-boilerplate/api

go 1.16

require (
	github.com/winnix/go-apaleo-one-boilerplate/api/clients/bookingclient v0.0.0-00010101000000-000000000000 // indirect
	github.com/winnix/go-apaleo-one-boilerplate/api/clients/integrationclient v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sessions v0.0.3 // indirect
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-openapi/errors v0.20.0 // indirect
	github.com/go-openapi/runtime v0.19.28 // indirect
	github.com/go-openapi/strfmt v0.20.1 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-openapi/validate v0.20.2 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/lib/pq v1.6.0 // indirect
	github.com/markbates/goth v1.67.1 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.10 // indirect
)

replace github.com/winnix/go-apaleo-one-boilerplate/api/clients/bookingclient => ./clients/bookingclient

replace github.com/winnix/go-apaleo-one-boilerplate/api/clients/integrationclient => ./clients/integrationclient
