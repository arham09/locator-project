package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Project struct {
	gorm.Model
	Name string
	Path string
}

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
		fmt.Println(i+1, project.Name)
	}
}

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

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "project.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	fmt.Println("done")
	// Migrate the schema
	db.AutoMigrate(&Project{})
}
