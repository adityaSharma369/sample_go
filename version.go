package main
import (
    "github.com/joho/godotenv"
    "fmt"
    "os"
"os/exec"
)

func main() {
	
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
    version := os.Getenv("INPUT_VERSION")
    fmt.Println("version----",version)
    app := "go"
    arg0 := "run"
    arg1 := "-ldflags"
    arg2 := "-X main.Version="+version
    fmt.Println("arg2=",arg2)
	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Start()
	fmt.Println("done")
}


