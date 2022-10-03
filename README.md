
go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}
go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/database/models --target ./pkg/database/ent

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

oapi-codegen --package=internal --generate types,chi-server,spec -o internal/delivery/controller/http/api/v1/api.gen.go api/_build/openapi.yaml


oapi-codegen --output-config --old-config-style --package=v1 --generate types,chi-server,spec api/_build/openapi.yaml