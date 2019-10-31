# ![RealWorld Example App](https://raw.githubusercontent.com/gothinkster/golang-gin-realworld-example-app/master/logo.png)

> ### GraphQL API server in golang.
# How it works
```
.
├── common
|   |── constant
|   |   |── constant.go      
│   ├── utils           
│   |   |── utils.go     //small tools function
├── container
|   |── constainer.go
├── gql       //code graphql
|   |── mutations       
|   |   |── fields
|   |   |  |── todo.go
|   |   |  |── ...
│   |   |── mutations.go
│   ├── queries     
|   |   |── fields
|   |   |  |── todo.go
|   |   |  |── ...
│   |   |── queries.go  
│   ├── types
│   |   |── todo.go
│   |   |── ...
├── model
|   |── todo.go
├── server
|   |── server.go
├── config.toml     //config file
├── main.go
```
# Getting started

## Install the Golang
https://golang.org/doc/install
## Environment Config
make sure your ~/.*shrc have those varible:
```
➜  echo $GOPATH
/path/test/
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/path/test//bin:/usr/local/go//bin
```

## Start
```
➜  go run main.go
```
