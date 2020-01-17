package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"go-cli/internal/model"
)

func OpenHandler() {
	cmd := exec.Command("bash", "-c", "code .")       // or whatever the program is
	cmd.Dir = "/home/arham/Documents/Project/golang/" // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("error", err)
	} else {
		fmt.Printf("%s", out)
	}
	fmt.Println("Opening Project")
}

func ReadHandler() {
	model.AllProjects()
}

func AddHandler(text *string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	model.AddProject(&text, &dir)
}