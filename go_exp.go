package main
import (
//         "github.com/joho/godotenv"
        "fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"os/exec"
)


func main() {
//   err := godotenv.Load()
//   if err != nil {
//     fmt.Println("Error loading .env file")
//   }
//     version := os.Getenv("INPUT_VERSION")
//     fmt.Println("version----",version)
// 	cmd := exec.Command("go build -ldflags=\"-X 'main.Version="+version+"'\" go_exp.go")
// 	cmd.Start()
// 	fmt.Println(cmd)
	jsonFile, err := os.Open("package.json")
	if err != nil {
        fmt.Println(err)
}
// 	fmt.Println("!!!!!!!",jsonFile)
	byteValue,   _ := ioutil.ReadAll(jsonFile)
	var fileContents map[string]interface{}

	json.Unmarshal([]byte(byteValue), &fileContents)
	var version=fmt.Sprintf("%v",fileContents["version"])
// 	fmt.Println(version)
	fmt.Println("version=", version)
	fmt.Println("its jenkins world")
	
	//declaring a integer variable x
	var x int
	x = 3                //assigning x the value 3
	fmt.Println("x:", x) //prints 3

	//declaring a integer variable y with value 20 in a single statement and prints it
	var y int = 80
	fmt.Println("y:", y)

	//declaring a variable z with value 50 and prints it
	//Here type int is not explicitly mentioned
	var z = 51
	fmt.Println("z:", z)
	fmt.Println("hi hello guys")

	//Multiple variables are assigned in single line- i with an integer and j with a string
	var i, j = 68, "hello"
	fmt.Println("i and j:", i, j)
	
	cmd := exec.Command("go build -ldflags=\"-X 'main.Version="+version)
	cmd.Start()
	fmt.Println(cmd)
}
