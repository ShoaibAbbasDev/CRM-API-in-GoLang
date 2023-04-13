package main

import (
	"github.com/ShoaibDevsinc/go-project/initializers"
	model "github.com/ShoaibDevsinc/go-project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnToDb()
}
func main() {
	initializers.DB.AutoMigrate(&model.Company{}, &model.Employee{})
}
