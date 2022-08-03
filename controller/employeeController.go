package controller

import (
	"net/http"
	"practice/mod/model"
	"practice/mod/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

// printName godoc
// @summary printName
// @description It will print Name
// @accept  json
// @produce json
// @success 200 {string} json
// @router /v1/practice/printName [Get]
// @tags Practice
func (ec EmployeeController) PrintName(context echo.Context) error {
	return context.JSON(http.StatusOK, "print Name working")
}

// allEmployeeDetails godoc
// @summary allEmployeeDetails
// @description It will return all employee details
// @accept  json
// @produce json
// @success 200 {string} json
// @router /v1/employee/allEmployeeDetails [Get]
// @tags employee
func (ec EmployeeController) GetAllEmployees(context echo.Context) error {

	respModel, err := ec.EmployeeService.GetEmployeeDetails()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "getting error in getting info from DB : "+err.Error())
	}
	respModel.Status = "success"
	return context.JSON(http.StatusOK, respModel)
}

// employeeDetails godoc
// @summary employeeDetails
// @description It will return employee details for a particular id
// @accept  json
// @produce json
// @param   id     path    string     true        "id corresponding to which you need emp details"
// @success 200 {string} json
// @success 404 {string} string	"No data found"
// @failure 400 {string} string "Bad Request."
// @failure 500 {string} string "An error ocurred while getting emp details"
// @router /v1/employee/employeeDetails/{id} [Get]
// @tags employee
func (ec EmployeeController) EmployeeDetails(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "provide valid id "+err.Error())
	}
	respModel, err := ec.EmployeeService.GetEmployeeDetailsbyId(idInt)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "getting error in getting info from DB : "+err.Error())
	}
	return context.JSON(http.StatusOK, respModel)
}

// createEmployee godoc
// @summary create
// @description It will create new employee details.
// @accept  json
// @produce json
// @param   args     body     model.EmpDetailsReq   true        "new employee details"
// @success 200 {string} json
// @success 404 {string} string	"No data found"
// @failure 400 {string} string "Bad Request."
// @failure 500 {string} string "An error ocurred while getting emp details"
// @router /v1/employee/create [Post]
// @tags employee
func (ec EmployeeController) CreateEmployee(context echo.Context) error {

	reqModel := &model.EmpDetailsReq{}
	err := context.Bind(reqModel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "provide valid entries "+err.Error())
	}

	err = ec.EmployeeService.AddDetails(*reqModel)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "getting error in adding in DB : "+err.Error())
	}

	return context.JSON(http.StatusOK, "success")
}

// deleteDetails godoc
// @summary deleteDetails
// @description It will delete employee details for a particular id
// @accept  json
// @produce json
// @param   id     path    string     true        "id corresponding to which you need to delete emp details"
// @success 200 {string} json
// @success 404 {string} string	"No data found"
// @failure 400 {string} string "Bad Request."
// @failure 500 {string} string "An error ocurred while getting emp details"
// @router /v1/employee/deleteDetails/{id} [Delete]
// @tags employee
func (ec EmployeeController) DeleteEmployee(context echo.Context) error {
	return context.JSON(http.StatusOK, "print Name working")
}

// updateDetails godoc
// @summary updateDetails
// @description It will update employee details for a particular id
// @accept  json
// @produce json
// @param   id     path    string     true        "id corresponding to which you need to update emp details"
// @param   args     body     model.EmpDetailsReq   true       "employee details to be updated"
// @success 200 {string} json
// @success 404 {string} string	"No data found"
// @failure 400 {string} string "Bad Request."
// @failure 500 {string} string "An error ocurred while getting emp details"
// @router /v1/employee/update/{id} [Put]
// @tags employee
func (ec EmployeeController) UpdateEmployee(context echo.Context) error {
	return context.JSON(http.StatusOK, "print Name working")
}
