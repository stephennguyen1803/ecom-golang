package main

import (
	"ecom-project/internal/routers"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	r := routers.NewServer()

	// Listen and serve on 8081
	r.Run(":8080")
}
