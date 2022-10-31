package models

import (
	"github.com/rbsilmann/api-estudo/db"
)

func Update(id int64, todo ToDo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	sql := `UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`
	query, err := conn.Exec(sql, todo.Id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}
	return query.RowsAffected()
}
