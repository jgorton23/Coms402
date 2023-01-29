# API staring point 

## Included features 
- http (Chi)
- http_authentication (authboss)
- authorization (casbin)
- linting (see .golangci.yaml)


## General TODO's
- Move over testing to magefiles
- Move over code generation to magefiles
    - add "go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}"
- Create build binary actions in magefile
- Create build Dockerfile actions in magefile
- Add go mod tidy / create cleanup actions in magefiles
- 
- Create getting started Section in the readme
- 
- Enable more golangci linters
- Unit testing
- Mock testing
- Have the request id on connection close and the only way to do that will be to modify middleware.RequestLogger
- start documentation