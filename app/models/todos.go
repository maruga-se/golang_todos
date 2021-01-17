package models

import (
	"database/sql"
	"log"
	"time"
	"todo_app/config"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (content, user_id, created_at)
			values(?,?,?)`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select * from todos where id = ?`
	todo = Todo{}

	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err

}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
