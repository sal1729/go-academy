package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	bbAge "github.com/bearbin/go-age"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

type Datasource interface {
	GetStudentData() []StudentDetails
}

type LocalDatasource struct {
	Filepath string
}

func (d LocalDatasource) GetStudentData() ([]StudentDetails, error) {
	// Read data from filepath to []StudentDetails
	localData, err := os.Open(d.Filepath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return nil, err
	}

	defer func() {
		if err := localData.Close(); err != nil {
			log.Printf("Warning: failed to close file %s: %v", d.Filepath, err)
		}
	}()

	var studentDetails []StudentDetails
	byteData, err := io.ReadAll(localData)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(byteData, &studentDetails)
	if err != nil {
		log.Printf("Error parsing json: %s\n", err)
		return nil, err
	}

	return studentDetails, nil
}

type ByFirstNameAsc []StudentDetails

func (a ByFirstNameAsc) Len() int           { return len(a) }
func (a ByFirstNameAsc) Less(i, j int) bool { return a[i].Name.FirstName < a[j].Name.FirstName }
func (a ByFirstNameAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByFirstNameDesc []StudentDetails

func (a ByFirstNameDesc) Len() int           { return len(a) }
func (a ByFirstNameDesc) Less(i, j int) bool { return a[i].Name.FirstName > a[j].Name.FirstName }
func (a ByFirstNameDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLastNameAsc []StudentDetails

func (a ByLastNameAsc) Len() int           { return len(a) }
func (a ByLastNameAsc) Less(i, j int) bool { return a[i].Name.LastName < a[j].Name.LastName }
func (a ByLastNameAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLastNameDesc []StudentDetails

func (a ByLastNameDesc) Len() int           { return len(a) }
func (a ByLastNameDesc) Less(i, j int) bool { return a[i].Name.LastName > a[j].Name.LastName }
func (a ByLastNameDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByDateOfBirthAsc []StudentDetails        // Equivalent to AgeDesc
func (a ByDateOfBirthAsc) Len() int           { return len(a) }
func (a ByDateOfBirthAsc) Less(i, j int) bool { return a[i].Dob < a[j].Dob }
func (a ByDateOfBirthAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByDateOfBirthDesc []StudentDetails        // Equivalent to AgeAsc
func (a ByDateOfBirthDesc) Len() int           { return len(a) }
func (a ByDateOfBirthDesc) Less(i, j int) bool { return a[i].Dob > a[j].Dob }
func (a ByDateOfBirthDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (s StudentDetails) getFullName() string {
	var middle string
	if s.Name.MiddleName == "" {
		middle = " "
	} else {
		middle = " " + s.Name.MiddleName + " "
	}
	return s.Name.FirstName + middle + s.Name.LastName
}
func (s StudentDetails) getAge() int {
	dob := s.Dob.ToDobType()
	birthday := time.Date(dob.Year, dob.Month, dob.Day, 0, 0, 0, 0, time.UTC)
	return bbAge.Age(birthday)
}

// Here we define the custom display for Student Details
func (s StudentDetails) String() string {
	return fmt.Sprintf("Student: %s, Age: %d", s.getFullName(), s.getAge())
}

func Register(students []StudentDetails, sortValue string, sortOrder string, display bool) ([]StudentDetails, error) {
	switch fmt.Sprintf("%s_%s", sortValue, sortOrder) {
	case "first_name_asc":
		sort.Sort(ByFirstNameAsc(students))
	case "first_name_desc":
		sort.Sort(ByFirstNameDesc(students))
	case "last_name_asc":
		sort.Sort(ByLastNameAsc(students))
	case "last_name_desc":
		sort.Sort(ByLastNameDesc(students))
	case "dob_asc", "age_desc":
		sort.Sort(ByDateOfBirthAsc(students))
	case "dob_desc", "age_asc":
		sort.Sort(ByDateOfBirthDesc(students))
	default:
		return nil, errors.New("invalid sort order/sort value pair")
	}

	if display {
		for _, student := range students {
			fmt.Println(student)
		}
	}
	return students, nil
}
