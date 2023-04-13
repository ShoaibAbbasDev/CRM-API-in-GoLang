package routes
import (
	"github.com/ShoaibDevsinc/go-project/controllers"
	"github.com/gin-gonic/gin"
)
func Routes(){
r := gin.Default()

// Create
r.POST("/create/company", controllers.Create_company)
r.POST("/company/:id/create/employee", controllers.Create_employee)

// Read
r.GET("/show/company/:id", controllers.Get_company)
r.GET("/companies", controllers.Get_companies)

r.GET("/employees", controllers.Get_employees)
r.GET("/show/employee/:id", controllers.Get_employee)
r.GET("/show/company/:id/employee", controllers.Get_company_employees)

// Update 
r.POST("/update/company/:id", controllers.Update_company)
r.POST("/update/employee/:id", controllers.Update_employee)

// Delete
r.DELETE("/remove/employee/:id", controllers.Delete_employee)
r.DELETE("/remove/company/:id", controllers.Delete_company)

r.Run()

}
