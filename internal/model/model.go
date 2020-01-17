package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Project Model for the orm
type Project struct {
	gorm.Model
	Name string
	Path string
}

// AllProjects - To get all the data stored in the db
func AllProjects() {
	db, err := gorm.Open("sqlite3", "project.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var projects []Project
	db.Find(&projects)

	if len(projects) == 0 {
		fmt.Println("There is no project saved")
	}

	for i, project := range projects {
		fmt.Println(i+1, project.Name, project.Path)
	}
}

// AddProject - To Add data to the database
func AddProject(name *string, path *string) {
	db, err := gorm.Open("sqlite3", "project.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&Project{Name: *name, Path: *path})
	fmt.Println("Data successfully inserted")
}

// func GetProjects(name *string) Project {
// 	return Project
// }

// InitialMigration - To initialize migration
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "project.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	fmt.Println("done")

	db.AutoMigrate(&Project{})
}
