package models

import (
	"database/sql"
	"log"
	"time"
	"todo_app/config"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID       int
	UUID     string
	Email    string
	UserId   int
	CreateAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (uuid,name,email,password,created_at) values (?, ?, ?, ?, ?)`

	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := "update users set name = ?, email = ? where id = ?"
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select * from users where email = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (uuid, email, user_id, created_at) values (?,?,?,?)`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select * from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreateAt,
	)

	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select * from sessions where uuid = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	err = Db.QueryRow(cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserId,
		&s.CreateAt,
	)

	if err != nil {
		valid = false
		return
	}
	if s.ID != 0 {
		valid = true
	}
	return valid, err
}

func (s *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	_, err = Db.Exec(cmd, s.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	Db, _ = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	err = Db.QueryRow(cmd, s.UserId).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}
