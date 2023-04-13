package main

import (
	"github.com/ShoaibDevsinc/go-project/initializers"
	"github.com/ShoaibDevsinc/go-project/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnToDb()
}

func main() {
	routes.Routes()

}
