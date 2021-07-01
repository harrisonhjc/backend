package main

import (
	"backend/router"
	_ "backend/docs"

)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name harrison
// @contact.url https://github.com/harrisonhjc/go-vue-project

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// schemes http
func main() {
	router.Init()
}
