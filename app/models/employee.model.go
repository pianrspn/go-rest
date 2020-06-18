package models

import (
	"net/http"

	"github.com/pianrspn/go-rest/database"
)

// Employee is ..
type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// FetchAllEmployee is ..
func FetchAllEmployee() (Response, error) {
	var obj Employee
	var arr []Employee
	var res Response

	conn := database.Init()
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM employee")

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Phone)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arr

	return res, nil
}
