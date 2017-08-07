package users

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"passwordHash"`
	Fname        string    `json:"fname"`
	Lname        string    `json:"lname"`
}

// NewUser struct
type NewUser struct {
	Name            string
	Email           string
	Fname           string
	Lname           string
	Password        string
	PasswordConfirm string
}

// Public inferface
type Public interface {
	Public() interface{}
}

// Public thing
func (u *User) Public() interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
	}
}

func (u *User) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT username FROM users WHERE id=$1;",
		u.ID).Scan(&u.Username)
}

// CheckPass func
func (u *User) CheckPass(db *sql.DB) error {
	return db.QueryRow("SELECT passwordHash, email, id FROM users WHERE username=$1;", u.Username).Scan(&u.PasswordHash, &u.Email, &u.ID)
}

// CreateUser creates a new users
func (u *User) CreateUser(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO users(username, passwordHash, email, fname, lname, id) VALUES($1, $2, $3, $4, $5, $6)",
		u.Username, u.PasswordHash, u.Email, u.Fname, u.Lname, u.ID,
	)
	return err
}

// UpdateUser for
func (u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users SET password=$1;", u.PasswordHash)
	return err
}

// CheckID looks for the id and returns a user.
func (u *User) CheckID(db *sql.DB) error {
	return db.QueryRow("SELECT username, email FROM users WHERE id=$1",
		u.ID).Scan(&u.Username, &u.Email)
}

func (u *User) getUsers(db *sql.DB, start, count int) ([]User, error) {
	fmt.Println("we are in users/models")
	rows, err := db.Query(
		"SELECT username FROM users LIMIT $1 OFFSET $2;",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fmt.Println("rows", rows)

	users := []User{}

	for rows.Next() {
		fmt.Println("whats in here")
		var u User
		if err := rows.Scan(&u.ID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
