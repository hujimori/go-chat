package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	name      string
	email     int
	password  string
	CreatedAt time.Time
}

func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users ORDER BY created_at DESC")

	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.name, &user.email, &user.password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func UserByEmail(email string) (user User, err error) {
	rows, err := Db.Query("SELECT id FROM users LIMIT 1")

	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Uuid, &user.name, &user.email, &user.password, &user.CreatedAt); err != nil {
			return
		}
	}
	rows.Close()
	return

}
