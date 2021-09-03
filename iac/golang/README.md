
## How to run a go program
go mod init example.com/hello
touch hello.go
go mod tidy  # It locates and downloades the modules that contains the imported packages (if any)
go run . 