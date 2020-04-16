package main
import (
//     "github.com/joho/godotenv"
    "fmt"
    "os"
// "os/exec"
)

func main() {
	
//   err := godotenv.Load()
//   if err != nil {
//     fmt.Println("Error loading .env file")
//   }
// Open our jsonFile
jsonFile, err := os.Open("package.json")
// if we os.Open returns an error then handle it
if err != nil {
    fmt.Println(err)
}
	 version := jsonFile.version
//     version := os.Getenv("INPUT_VERSION")
    fmt.Println("version--0000--",version)
	cmd := exec.Command("go build -ldflags=\"-X 'main.Version="+version+"'\" go_exp.go")
	cmd.Start()
	fmt.Println(cmd)
}
