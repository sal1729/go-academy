# README

The to-do app "franz" lives in `franz-fingers`. An earlier, more object-orient-y, less client-side-javascript-y version lives in `todo-app-web`.
For the latest "franz" version, check out `feature/todo-lizt`.

## File Structure

The file structure is a little haphazard as it grew out of the exercises.
- `assignment-xyz` contains early stuff from the pre-todo app exercises
- `todo-app-exercises` contains the `main` programs for the named exercises + some files with data in
- `todo-app-functions` contains the functions written as part of the exercises in `todo-app-exercises` but these have been tweaked as part of the more general to do app work.
They wrap up a json file containing a list of tasks as a datasource, with crud operations.
It also contains the cli tool version of the franz app.
- `todo-app-web` contains the web app, which began life as exercise 18.
This has a set of functions which depend heavily on `todo-app-functions`.
It's API-ish, definitely not restful.

Then there's the more recent stuff
- `franz-brain` contains the crud operations for manipulating an in-memory list + some read/write stuff for persisting the list to a json file.
Also the data file, and the api handler (maybe misplaced tbh).
- `franz-fingers` contains the three access points
  - `franz-cli` - cli tool to make requests against the datastore
  - `franz-api` - an attempt to wrap the datastore in an api. It's not _too_ bad, but I can't work out how to wire it up with the cli tool.
  - `franz-web` - a web app which serves a web page full of client-side js for boshing the api, which gets served at the same time. 

## Branches

-  `feature/a-failed-attempt-to-rest` - An attempt to make things more restful -
   I gave this a go for a morning, but gave up because there were lot's of niggles and it probably needed completely re-writing.
   I wanted to get on to implementing some go routines so we're just going to have to accept the less good API.
- `feature/no-go` - The working app before I attempted any refactoring of the datasource logic.
  This doesn't use any go routines.
  It's undoubtedly not very performant as every CRUD operation against the data is wrapped in a read/write to the file containing the data.
- `feature/oops` - The working app after adding in some concurrency/goroutines _but_ this is too object-orient-y. I think I must have massively misunderstood the exercises somewhere 🤷
However, this version of the app has very little client-side logic, due to the separate handlers for each action.
- `feature/todo-lizt` - The latest working version of the "franz" to do app, with no extra spam.
- `main` - The latest version, with all the trimmings. A less object-orient-y re-write.
The web app has considerably more client-side logic than the previous version now that the only routes are `/franz` and `/api`.
I like this a lot less. The previous version has a much cleaner separation between frontend and backend, imho.
Still TODO:
   * Get the cli tool using the api <- not sure how to do this
   * Tests
   * Users/context logging/database backend etc (v. unlikely I'll get round to this)
   * Bosh through and tidy up some todos