# franz-cli

Command line access to the franz datastore.
Based on exercise-17.go
Run `go build -o franz main.go` to construct the executable
Our app is called franz, after Franz Liszt.

## Example commands
- CREATE: `./franz -create "New Task Name" -status "In Progress"`
- READ: `./franz -list all`, `./franz -list "New Task Name"`, `./franz -list all -status "Status"`
- UPDATE: `./franz -update "New Task Name" -status "New Status"`, `./franz -update all -status "To Do*TO*Blocked"`
- DELETE: `./franz -delete all`,`./franz -delete "New Task Name"`, `./franz -delete all -status "Done"`