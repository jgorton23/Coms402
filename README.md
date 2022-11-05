
# Future Work

## Casbin

Casbin will provide authorization for the API routes

## Migrate middleware.RequestLogger 

I want to have the request id on connection close and the only way to do that will be to modify middleware.RequestLogger

## Migrate authboss logger

I want to include the request_id on every log message 


go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}