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
	cmd := exec.Command("go build -ldflags=\"-X 'main.Version="+version+"'\" go_exp.go")
	cmd.Start()
	fmt.Println(cmd)
}
