package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUES ('eko', 'EKO')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id , name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("==================================")
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id , name, email , balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		//var email string
		var email sql.NullString
		var balance int32
		var rating float64
		//var birthDate time.Time
		var birthDate sql.NullTime
		var createAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============================================================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate)
		}
		fmt.Println("Married:", married)
		fmt.Println("Create At:", createAt)
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login dengan Username", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login dengan Username", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
func TestExecSQLParameterSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ilham'; DROP TABLE user; #"
	password := "ilham"

	script := "INSERT INTO user (username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "nadhif@gmail.com"
	comment := "Halo nama saya Nadhif"

	script := "INSERT INTO comments  (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new user with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments  (email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "ilham" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id", id)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments  (email, comment) VALUES (?, ?)"
	stmt, err := tx.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "ilham" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id", id)

	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
