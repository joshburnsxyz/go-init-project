package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	projectName string
	pkgDecl     string
)

func initFlags() {
	flag.StringVar(&projectName, "name", "", "Name for your project")
	flag.StringVar(&pkgDecl, "pkg", "", "The package declaration (github.com/username/mypackage)")
}

func main() {

	// Parse CLI flags
	initFlags()
	flag.Parse()

	// Create project directory
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		os.MkdirAll(projectName, 0777)
	}

	// Generate project files
	err := os.Chdir(projectName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=> Creating main.go")
		mainDotGoContent := []byte(fmt.Sprintf("package %s", projectName))
		writeFileErr := os.WriteFile("./main.go", mainDotGoContent, 0777)
		if writeFileErr != nil {
			// FIXME: Handle writeFileErr
			panic(writeFileErr)
		}
		mainDotGoContent = nil
		writeFileErr = nil

		fmt.Println("=> Initializing go modules")
		cmdOut, cmdErr := exec.Command("go", "mod", "init", pkgDecl).Output()
		if cmdErr != nil {
			// FIXME: Handle cmdErr properly
			panic(cmdErr)
		}
		fmt.Println(cmdOut)

		fmt.Println("=> Tidying modules")
 		cmdOut, cmdErr = exec.Command("go", "mod", "tidy").Output()
		if cmdErr != nil {
			// FIXME: Handle cmdErr properly
			panic(cmdErr)
		}
		fmt.Println(cmdOut)
    cmdOut = nil
    cmdErr = nil
   
		fmt.Println("=> Done")
	}
}
