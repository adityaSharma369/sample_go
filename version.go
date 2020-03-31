package main

import "os/exec"
import "os"
import "fmt"
import "github.com/joho/godotenv"

func main() {
    err := godotenv.Load()
      if err != nil {
        fmt.Println("Error loading .env file")
      }
    
    
    var myEnv map[string]string
    myEnv, err2 := godotenv.Read()
    fmt.Println(err2)
    version := myEnv["VERSION"]

    
//     version := os.Getenv("INPUT_VERSION")
    fmt.Println("version----",version)
    app := "go"
    arg0 := "run"
    arg1 := "-ldflags"
    arg2 := "-X main.Version="+version
    fmt.Println("arg2=",arg2)
    exec.Command(app, arg0, arg1, arg2)
}


