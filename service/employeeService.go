package service

import (
	"practice/mod/dblayer"
	"practice/mod/model"
)

type EmployeeService struct {
	DbClient dblayer.DBRepo
}

func (es EmployeeService) AddDetails(reqDetails model.EmpDetailsReq) (err error) {
	err = es.DbClient.AddDetailsinDB(reqDetails)
	return err
}

func (es EmployeeService) GetEmployeeDetailsbyId(id int) (retModel model.EmployeeDetails, err error) {
	retModel, err = es.DbClient.GetDetailsfromDBforId(id)
	return retModel, err
}

func (es EmployeeService) GetEmployeeDetails() (retModel model.AllDetailsResp, err error) {
	retModel, err = es.DbClient.GetDetailsfromDB()
	return retModel, err
}
