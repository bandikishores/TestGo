# TestGo
Sample project to test out go features

* Setting Up Project
   1) Clone the repo into $GOPATH/bandi.com/
   2) Run `make` command inside $GOPATH/bandi.com/TestGo Folder
   3) To RUN application
        `go run main.go`
   4) To RUN grpc get user client use
        `go run examples/main/main_grpc_client.go `
   5) To RUN tests
        `make test`
   6) To RUN Benchmark Tests
        `make bench` or manually run `go test -bench=.`
   7) Running Swagger Document
        `swagger serve data/swagger/apidocs.swagger.json`


* Sample Visual Studio Launch.json for Debugging code. (Setup Visual Studio by adding folder bandi.com in workspace)
  1) Paste this in launch.json

        ```json
        {
            // Use IntelliSense to learn about possible attributes.
            // Hover to view descriptions of existing attributes.
            // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
            "version": "0.2.0",
            "configurations": [
                {
                    "name": "Launch",
                    "type": "go",
                    "request": "launch",
                    "mode": "auto",
                    "program": "${workspaceFolder}/TestGo/main.go",
                    "env": {},
                    "args": []
                }
            ]
        }
        ```


* After Launching Application
  1) Create User
      ```
        curl -X POST \
        http://localhost:18081/organizations/bandi/users \
        -H 'Content-Type: application/json' \
        -d '{
                "name": "kishore",
                "displayName": "bandi kishore",
                "email": "bandikishores@gmail.com"
            }'
      ```
  2) Get User
      `curl -X GET http://localhost:18081/organizations/bandi/users/kishore?queryParam1=123 -H 'Host: localhost:18081'`
  3) For Streaming Responses
      `curl -X GET http://localhost:18081/organizations/bandi/streamusers/kishore?queryParam1=123 -H 'Host: localhost:18081'`

            OR

      `curl -X GET http://localhost:18081/streamusers/kishore?queryParam1=123 -H 'Host: localhost:18081'`
  4) Update User
      ```
        curl -X PUT \
        http://localhost:18081/organizations/bandi/users/kishore \
        -H 'Content-Type: application/json' \
        -d '{
            "user" : {
                "name": "kishore",
                "displayName": "bandi kishore s",
                "email": "bandikishores@gmail.com"
            }
        }'
      ```


* Manual Installation and Compilation of Proto/Go

  1) Go Mod Created using 
      `go mod init bandi.com/main`
  2) To Compile ProtoBuf from root directory
      `protoc --go_out=./pkg/ ./data/proto/addressbook.proto`
  3) 