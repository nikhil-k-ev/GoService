package dblayer

import "practice/mod/model"

func (dc *DBRepo) AddDetailsinDB(reqDetails model.EmpDetailsReq) (err error) {

	sqlQuery := "insert into public.Employees (name, age, salary, profile_picture) values ($1,$2,$3,$4)"
	_, err = dc.SqlDB.Exec(sqlQuery, reqDetails.Name, reqDetails.Age, reqDetails.Salary, reqDetails.ProfilePic)
	return err
}

func (dc *DBRepo) GetDetailsfromDBforId(id int) (retModel model.EmployeeDetails, err error) {

	sqlQuery := "select id, age, name, salary, profile_picture from public.Employees where id = $1"
	err = dc.SqlDB.QueryRow(sqlQuery, id).Scan(&retModel.Id, &retModel.Age, &retModel.Name, &retModel.Salary, &retModel.ProfilePic)
	return retModel, err
}

func (dc *DBRepo) GetDetailsfromDB() (retModel model.AllDetailsResp, err error) {

	sqlQuery := "select id, age, name, salary, profile_picture from public.Employees"
	rows, err := dc.SqlDB.Query(sqlQuery) //Scan(&retModel.Id, &retModel.Age, &retModel.Name, &retModel.Salary, &retModel.ProfilePic)
	if err != nil {
		return retModel, err
	}
	for rows.Next() {
		employeeDetails := model.EmployeeDetails{}
		rows.Scan(&employeeDetails.Id, &employeeDetails.Age, &employeeDetails.Name, &employeeDetails.Salary, &employeeDetails.ProfilePic)
		retModel.Data = append(retModel.Data, employeeDetails)
	}
	return retModel, err
}
