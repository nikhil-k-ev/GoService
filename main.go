package main

import (
	"fmt"
	"practice/mod/controller"
	"practice/mod/dblayer"
	"practice/mod/service"

	_ "practice/mod/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Practice Service
// @version 1.0
// @description Service to practice making of new API

// @contact.name Nikhil Kumar

// @contact.email nikhil.k@eagleview.com

// @BasePath /test

func main() {
	e := echo.New()

	apiBasePath := "/test"

	dbRepo, err := initDBClient()
	if err != nil {
		fmt.Print("", "", "Unable to Initialize DB due to"+err.Error())
	}
	empService := service.EmployeeService{
		DbClient: dbRepo,
	}
	empCon := controller.EmployeeController{
		EmployeeService: empService,
	}

	e.GET(apiBasePath+"/v1/practice/printName", empCon.PrintName)
	e.GET(apiBasePath+"/v1/employee/allEmployeeDetails", empCon.GetAllEmployees)
	e.GET(apiBasePath+"/v1/employee/employeeDetails/:id", empCon.EmployeeDetails)
	e.POST(apiBasePath+"/v1/employee/create", empCon.CreateEmployee)
	e.DELETE(apiBasePath+"/v1/employee/deleteDetails/:id", empCon.DeleteEmployee)
	e.PUT(apiBasePath+"/v1/employee/update/:id", empCon.UpdateEmployee)

	// e.GET("/", func(c echo.Context) (err error) {
	// 	return c.HTML(http.StatusOK, "<head><meta http-equiv=\"refresh\" content=\"0; url=/swagger/index.html\" /></head><body>p><a href=\"/swagger/index.html\">Redirect</a></p></body>")
	// })

	e.GET(apiBasePath+"/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8000"))
}

func initDBClient() (dblayer.DBRepo, error) {
	_dbrepo := dblayer.DBRepo{}
	err := _dbrepo.DBConnect()
	if err != nil {
		return _dbrepo, err
	}

	fmt.Print("connection to the database was successful")
	return _dbrepo, nil
}
