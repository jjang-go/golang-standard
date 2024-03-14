package main

import (
	"flag"

	"github.com/jjang-go/golang-standard/init/cmd"
)

var configPathFlag = flag.String("config", "./cofnig.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd("./config.toml")
}

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	// http를 사용
// 	http.HandleFunc("/", helloWorld)

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("error")
// 		panic(err)
// 	}
// }

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("hello World")
// }
