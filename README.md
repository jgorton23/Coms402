
go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}
go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/database/models --target ./pkg/database/ent

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

oapi-codegen --package=internal --generate types,chi-server,spec -o internal/delivery/controller/http/api/v1/api.gen.go api/_build/openapi.yaml


oapi-codegen --output-config --old-config-style --package=v1 --generate types,chi-server,spec api/_build/openapi.yaml


# Example api with user registration and management

## To Implement

### authorization through casbin 

    Implementing user authorization using casbin should not be that hard but there are some items to consider 

    - Where to store the conf file 
    - Re implement or use the existing adapter

### Migrate authboss logger to conform to our logger

Switching authbosses logger will be tricky and not trivial 

### Switch config to use go embed