package main

import (
	_ "ecom-project/cmd/swag/docs"
	"ecom-project/internal/initialize"
)

// @title           API Ecommerce Backend Golang Project
// @version         1.0
// @description     This is API Document for Ecommerce Backend Golang Project.
// @termsOfService  https://github.com/stephennguyen1803/ecom-golang

// @contact.name   Anh Dung Nguyen
// @contact.url    https://www.linkedin.com/in/anhdungnguyen-95262198/
// @contact.email  anhdung.phc@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8090
// @BasePath  /v1/2024
// @schemes   http
func main() {
	initialize.Run()
}
