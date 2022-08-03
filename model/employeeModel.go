package model

type EmployeeDetails struct {
	Id         int    `json:"employee_id"`
	Name       string `json:"employee_name"`
	Salary     int    `json:"employee_salary"`
	Age        int    `json:"employee_age"`
	ProfilePic string `json:"profile_picture"`
}

type AllDetailsResp struct {
	Status string            `json:"status"`
	Data   []EmployeeDetails `json:"data"`
}

type EmpDetailsReq struct {
	Name       string `json:"name"`
	Salary     int    `json:"salary"`
	Age        int    `json:"age"`
	ProfilePic string `json:"profile_picture"`
}
