# README

The to-do app "franz" lives in `todo-app-web`.

The file structure is a little haphazard as it grew out of the exercises.
- `todo-app-exercises` contains the `main` programs for the named exercises + some files with data in
- `todo-app-functions` contains the functions written as part of the exercises in `todo-app-exercises` but these have been tweaked as part of the more general to do app work.
They wrap up a json file containing a list of tasks as a datasource, with crud operations.
It also contains the cli tool version of the franz app.
- `todo-app-web` contains the web app, which began life as exercise 18.
This has a set of functions which depend heavily on `todo-app-functions`.
It's API-ish, definitely not restful. See the branch `feature/a-failed-attempt-to-rest` for an attempt to make things more restful -
I gave this a go for a morning, but gave up because there were lot's of niggles and it probably needed completely re-writing.
I wanted to get on to implementing some go routines so we're just going to have to accept the less good API.