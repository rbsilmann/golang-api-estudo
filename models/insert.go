package models

import (
	"github.com/rbsilmann/api-estudo/db"
)

func Insert(newRegister ToDo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, newRegister.Title, newRegister.Description, newRegister.Done).Scan(&id)
	return
}
