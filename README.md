# API staring point 

## Getting Started 

Required software

- Golang 
- Docker 

All other (go) tools are installed automatically when required or use docker

Using the included tools 

```go run mage.go ```

Will print out a helpful list of tools that can be ran. Running the 
Server is as simple as 

```go run mage.go server:run```  

The tools are namespaced into similar catagories

for example all the linters are under ```linter``` ad 
all the database tools are under ```database```



## Included features 
- http (Chi)
- http_authentication (authboss)
- authorization (casbin)
- linting (see .golangci.yaml)


## General TODO's
- Move over testing to magefiles
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