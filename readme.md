# README - `feature/todo-liszt`

The to-do app "franz" lives in `franz-fingers`. An earlier, more object-orient-y, less client-side-javascript-y version lives in `todo-app-web`, but you'll need to checkout `main` for that.

## Running the app
To run the web app, first ensure the `datastore` constant is set correctly for you. That is, update line 13 of `franz-brain/readWriteFunctions.go`.
Then run `go run .` from `franz-fingers/franz-web`.

## File Structure
- `franz-brain` contains the crud operations for manipulating an in-memory list + some read/write stuff for persisting the list to a json file.
Also the data file, and the api handler (maybe misplaced tbh).
- `franz-fingers` contains the three access points
  - `franz-cli` - cli tool to make requests against the datastore
  - `franz-api` - an attempt to wrap the datastore in an api. It's not _too_ bad, but I can't work out how to wire it up with the cli tool.
  - `franz-web` - a web app which serves a web page full of client-side js for boshing the api, which gets served at the same time. 

## Branch Status

This is the latest working version of the "franz" to do app, with no extra spam. It's a less object-orient-y re-write of what came before.
The web app has considerably more client-side logic than the previous version now that the only routes are `/franz` and `/api`.
I like this a lot less. The previous version has a much cleaner separation between frontend and backend, imho.
Still TODO:
   * Get the cli tool using the api <- not sure how to do this
   * Tests
   * Users/context logging/database backend etc (v. unlikely I'll get round to this)
   * Bosh through and tidy up some todos