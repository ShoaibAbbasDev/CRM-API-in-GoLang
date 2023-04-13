package controllers

import (
	"net/http"

	"github.com/ShoaibDevsinc/go-project/initializers"
	model "github.com/ShoaibDevsinc/go-project/models"
	"github.com/gin-gonic/gin"
)

type Stage int

// implementing Stage Enum

const (
	Diligence Stage = iota + 1
	Lead
	Rejected
	Closed
)

func (s Stage) String() string {
	return [...]string{"Diligence", "Lead", "Rejected", "Closed"}[s-1]
}

func (s Stage) EnumIndex() int {
	return int(s)
}

var stage = Rejected
var enum_value = stage.EnumIndex()

// new company

func Create_company(c *gin.Context) {

	company := model.Company{}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company_data := model.Company{Name: company.Name}
	val_err := company_data.Validate()

	if val_err == nil {
		result := initializers.DB.Create(&company_data)
		if result != nil {
			c.JSON(400, gin.H{
				"company": company_data,
			})

		} else {
			c.JSON(200, gin.H{
				"company": company_data,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"status": val_err,
		})
	}
}

// new employee

func Create_employee(c *gin.Context) {

	employee := model.Employee{}
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var company_obj model.Company
	id, _ := c.Params.Get("id")

	initializers.DB.First(&company_obj, id)

	if company_obj.ID != 0 {

		initializers.DB.Model(&company_obj).Update(
			"Employees", []model.Employee{{First_name: employee.First_name, Last_name: employee.Last_name, Email: employee.Email, Phone: employee.Phone, Stage: employee.Stage, Probability: employee.Probability}},
		)
		val_err_emp := employee.Validate()

		if val_err_emp == nil {
			result := initializers.DB.Create(&company_obj)

			if result != nil {
				c.JSON(400, gin.H{
					"employee": company_obj,
				})
			} else {
				c.JSON(200, gin.H{
					"status": "company not found",
				})
			}
		} else {
			c.JSON(200, gin.H{
				"status": val_err_emp,
			})
		}
	}
}

// select company

func Get_company(c *gin.Context) {

	var company_obj model.Company
	id, _ := c.Params.Get("id")

	result := initializers.DB.First(&company_obj, id)

	if result != nil {

		if company_obj.ID != 0 {
			c.JSON(400, gin.H{
				"company": &company_obj,
			})

		} else {
			c.JSON(404, gin.H{
				"status": "Record Not Found",
			})
		}
	}
}

// select all comapnies in database

func Get_companies(c *gin.Context) {

	var companies []model.Company

	result := initializers.DB.Find(&companies)

	if result != nil {
		c.JSON(400, gin.H{
			"company": &companies,
		})
	}
}

// select employee

func Get_employee(c *gin.Context) {

	var emp model.Employee
	id, _ := c.Params.Get("id")
	result := initializers.DB.First(&emp, id)

	if result != nil {
		if emp.ID != 0 {
			c.JSON(400, gin.H{
				"employee": &emp,
			})

		} else {
			c.JSON(404, gin.H{
				"status": "Record Not Found",
			})
		}
	}
}

// select all employees in database

func Get_employees(c *gin.Context) {

	var emp []model.Employee
	result := initializers.DB.Find(&emp)

	if result != nil {
		c.JSON(400, gin.H{
			"employees": &emp,
		})
	}
}

// select employee which is associated to given company ID

func Get_company_employees(c *gin.Context) {

	var emp []model.Employee
	var company_obj model.Company
	id, _ := c.Params.Get("id")
	result_c := initializers.DB.First(&company_obj, id)
	result_e := initializers.DB.Where("company_id = ?", id).Find(&emp)

	if result_c != nil {
		if result_e != nil {
			if company_obj.ID != 0 {
				c.JSON(400, gin.H{
					"company_employess": &emp,
				})
			} else {
				c.JSON(404, gin.H{
					"status": "Record Not Found",
				})
			}
		}
	}
}

// Update company detail

func Update_company(c *gin.Context) {

	var company_obj model.Company
	id, _ := c.Params.Get("id")
	company := model.Company{}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := initializers.DB.First(&company_obj, id)

	if result != nil {
		if company_obj.ID != 0 {
			initializers.DB.Model(&company_obj).Update("Name", company.Name)
			c.JSON(400, gin.H{
				"company": &company_obj,
			})
		} else {
			c.JSON(404, gin.H{
				"status": "Record Not Found",
			})
		}
	}
}

// Update employee detail

func Update_employee(c *gin.Context) {

	var emp model.Employee
	employee := model.Employee{}
	id, _ := c.Params.Get("id")
	result := initializers.DB.First(&emp, id)

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result != nil {
		if emp.ID != 0 {
			initializers.DB.Model(&emp).Updates(model.Employee{
				First_name: employee.First_name, Last_name: employee.Last_name, Email: employee.Email,
				Phone: employee.Phone, Stage: employee.Stage, Probability: employee.Probability,
			},
			)
			c.JSON(400, gin.H{
				"employee": &emp,
			})
		} else {
			c.JSON(404, gin.H{
				"status": "Record Not Found",
			})
		}
	}
}

// delete employee

func Delete_employee(c *gin.Context) {

	var emp model.Employee
	id, _ := c.Params.Get("id")
	result := initializers.DB.First(&emp, id)

	if result != nil {
		if emp.ID != 0 && &emp.DeletedAt != nil {
			initializers.DB.Delete(&emp)
			c.JSON(404, gin.H{
				"status": "Record Deleted",
			})
		} else {
			c.JSON(404, gin.H{
				"status": "Record Not Found",
			})
		}
	}
}

// delete company details

func Delete_company(c *gin.Context) {

	var emp []model.Employee
	var company_obj model.Company
	id, _ := c.Params.Get("id")
	result_c := initializers.DB.First(&company_obj, id)
	result_e := initializers.DB.Where("company_id = ?", id).Find(&emp)

	if result_c != nil {
		if result_e != nil {
			if company_obj.ID != 0 && &company_obj.DeletedAt != nil {
				initializers.DB.Delete(&emp)
				initializers.DB.Delete(&company_obj)
				c.JSON(400, gin.H{
					"status": "Record Deleted",
				})
			} else {
				c.JSON(404, gin.H{
					"status": "Record Not Found",
				})
			}
		}
	}
}
