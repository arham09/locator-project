package handler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"go-cli/internal/model"
)

func OpenHandler(name string) {
	project := model.GetProjects(&name)

	cmd := exec.Command("bash", "-c", "code .") // or whatever the program is
	cmd.Dir = project.Path                      // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("error", err)
	} else {
		fmt.Printf("%s", out)
	}
	fmt.Println("Opening Project")
}

func GetHandler(name string) {
	project := model.GetProjects(&name)

	fmt.Println(project.Name)
	fmt.Println(project.Path)
}

func ReadHandler() {
	model.AllProjects()
}

func AddHandler(text string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	model.AddProject(&text, &dir)
}
