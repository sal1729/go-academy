package main

import (
	"go-academy/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// TODO Use log more consistently across other programs
	log.SetPrefix("city-display: ")
	log.SetFlags(0)
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	dir, err := os.Getwd()
	catchError(err)

	studentDataSource := utils.LocalDatasource{Filepath: filepath.Join(dir, "student-data.json")}
	students, err := studentDataSource.GetStudentData()
	catchError(err)

	students, err = utils.Register(students, "first_name", "asc", true)
	catchError(err)
}
