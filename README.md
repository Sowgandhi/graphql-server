# graphql-server
Repo created for practicing graphql in go

# Getting Started
* Initialize go mod file by running `go mod init graphql-server`
* Install gqlgen by running the command `go get github.com/99designs/gqlgen`
* Initialize new project using recommended project structure by running the command `go run github.com/99designs/gqlgen init`
* Modify the schema in schema.graphqls 
* Generate structs according to defined schema by running `go run github.com/99designs/gqlgen generate` from graph directory 
* Your logic goes in `schema.resolvers.go` file
* Command to run the server `go mod server.go`

# Sample Query Structure 
```
query getEmployees{
      employees{
            id
            name
            salary
            address{
               street
               city
               state
               pinCode
             }
         }
      }

mutation addEmployee{
      addEmployee(input:{
            name: "",
            age: 0,
            salary: 0,
            street: "",
            city: "",
            state: "",
            pinCode: ""
            }){
                id
                name
                age
                salary
                address {
                  street
                  state
                  pinCode
                }

             }
        }

```
