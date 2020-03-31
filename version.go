package main

import "os/exec"
import "os"

func main() {
    version = os.Getenv("Version")
    app := "go"
    arg0 := "run"
    arg1 := "-ldflags"
    arg2 := "-X main.Version="+version
    cmd := exec.Command(app, arg0, arg1, arg2)
}
