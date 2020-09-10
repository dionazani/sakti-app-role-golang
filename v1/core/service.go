package core

import (
	"fmt"

	_ "github.com/lib/pq"
)

func Insert(id int32, code string, description string) string {

	var result string = ""

	db, err := connect()
	if err != nil {
		result = err.Error()
		return result
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO app_role (id, code, description) VALUES ($1, $2, $3)", id, code, description)
	if err != nil {
		result = err.Error()
	}

	return result

}

func GetById(id int64) *appRoleModel {

	var entity = appRole{}
	var model = appRoleModel{}

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return &model
	}
	defer db.Close()

	// select from app_role
	stmt, err := db.Prepare("SELECT id, code, description, is_active, created_at FROM app_role WHERE id = $1")
	if err != nil {
		fmt.Println(err.Error())
		return &model
	}
	defer db.Close()

	stmt.QueryRow(id).Scan(&entity.Id, &entity.Code, &entity.Description, &entity.IsActive, &entity.CreatedAt)

	model.Id = entity.Id
	model.Code = entity.Code
	model.Description = entity.Description
	model.IsActive = entity.IsActive
	t := entity.CreatedAt
	model.CreatedAt = t.Format("2006-01-02 15:04:05")

	return &model
}
