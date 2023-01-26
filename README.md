# API staring point 

## Included features 
- http (Chi)
- http_authentication (authboss)
- authorization (casbin)
- linting (see .golangci.yaml)


## General TODO's
- Enable more golangci linters
- Unit testing
- Mock testing
- Have the request id on connection close and the only way to do that will be to modify middleware.RequestLogger
- add "go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}" to makefile 
- cleanup makefile 
- start documentation