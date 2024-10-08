# readme

## Basic commands

`go mod init example.com/filename` to initialise a module

`go mod tidy` to synchronise dependencies, and generate a pseudo version number

`go mod edit -replace example.com/filename=../filename` to update with local location

`go run .` to run a `main()`

`go build` to compile into an executable

`./filename` to execute the executable (before installing)

`go install` to install the executable (the executable file will vanish to the bin 🪄)

`filename` to run the installed program

`go test` to run the test files in the active directory

`go test -v` to run the test files in the active directory, with verbose logging

// TODO How do we get rid of installed programs?