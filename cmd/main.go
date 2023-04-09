package main

import (
	"fmt"
	"os"

	_ "github.com/0xFinish/TraderJoeBot/pkg/config"
	// "github.com/0xFinish/TraderJoeBot/pkg/controllers"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(os.Getenv("TELEGRAM_TOKEN"))
}
