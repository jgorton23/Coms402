
go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}
go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/database/models --target ./pkg/database/ent

