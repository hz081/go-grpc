
# Go GRPC Client & Server

If you want to compile *.proto, you need to [download](https://developers.google.com/protocol-buffers/docs/downloads) it first.
Set the proto binary to your path, example: `export PATH=/doc/proto/bin:$PATH` in your .bash_profile or .bashrc or .zshrc

While you see the `compile.sh` file, you will realize that there is _$PROTOPATH_ variable, here is the solution: add `export PROTOPATH=/doc/proto/include:$PATH` to your environment variable.

### Run the server
    `go run server/main.go`

### Run the client
    `go run client/main.go`

### Client End Point
    `http://localhost:18080/add/10/10`
    
You will see the result as: `client result for add 20`
