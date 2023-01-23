package main

import (
  "github.com/bitfield/script"
  "flag"
  "os"
  "os/exec"
)

var (
  projectName string
)

func initFlags() {
  flag.Var(&projectName, "name", "Name for your project")
}

func main() {

  // Parse CLI flags
  initFlags()
  flag.Parse()

  // Create project directory
  dirExists,err := script.IfExists(projectName)
  if err != nil {
    fmt.Println("> ERROR: %s", err)
  }
  if dirExists == false {
    fmt.Println("> ERROR: Project directory already exists...exiting")
    os.Exit(3)
  } else {
    os.MkdirAll(projectName, 0644)
  }

  // Generate project files
  err := os.Chdir(projectName)
  if err != nil {
    fmt.Println("> ERROR: Could not cd into new project directory, wtf?")
  } else {
    createMainFileCmd := exec.Command("touch main.go")
    initGoModuleCmd := exec.Command(fmt.Sprintf("go mod init %s", projectName))

    err := createMainFileCmd.Run()
    if err == nil {
      script.File("main.go").WriteFile("package main")
      err := initGoModuleCmd.Run()
      if err == nil {
        fmt.Println("done")
      } else {
        fmt.Println("> ERROR: Could not initialize go module")
        os.Exit(2)
      }
    } else {
        fmt.Println("> ERROR: Could not create main.go")
        os.Exit(2)
    }
  }
}
