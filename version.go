package main

import "os/exec"
import "os"
import "fmt"

func main() {
    version := os.Getenv("Version")
    fmt.Println("version----",version)
    app := "go"
    arg0 := "run"
    arg1 := "-ldflags"
    arg2 := "-X main.Version="+version
    fmt.Println("arg2=",arg2)
    exec.Command(app, arg0, arg1, arg2)
}


