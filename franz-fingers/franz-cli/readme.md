# franz-cli

Command line access to the franz datastore.
Based on exercise-17.go
Decide which version you want, update the `accessMethod` variable in `cli.go` then run `go build -o franz cli.go` to construct the executable
Our app is called franz, after Franz Liszt.

## Example commands
- CREATE: `./franz -create "Clone the Car" -status "To Do"`
- READ: `./franz -list all`, `./franz -list "Clone the Car"`, `./franz -list all -status "To Do"`
- UPDATE: `./franz -update "Clone the Car" -status "In Progress"`, `./franz -update all -status "To Do*TO*Blocked"`
- DELETE: `./franz -delete "Clone the Car"`, `./franz -delete all -status "Done"`, `./franz -delete all`