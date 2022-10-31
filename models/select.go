package models

import (
	"github.com/rbsilmann/api-estudo/db"
)

func SelectOne(id int64) (todo ToDo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `SELECT * FROM todos WHERE id = $1`
	query := conn.QueryRow(sql, id)
	err = query.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	return
}

func SelectAll() (todos []ToDo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `SELECT * FROM todos`
	query, err := conn.Query(sql)
	if err != nil {
		return
	}
	for query.Next() {
		var todo ToDo
		err = query.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
