package belajargolangdatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsertExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer (id, name) VALUES ('2','Si Tampan')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Customer")
}
func TestQueryExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id , name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
	defer rows.Close()
}

func TestQueryExecSqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married , created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name, email, birthDate string
		var balance int64
		var rating float32
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
	defer rows.Close()
}
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	script := "SELECT username FROM user WHERE username = '" + username + "' AND PASSWORD = '" + password + "' LIMIT 1"
	fmt.Println(script)
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
		fmt.Println("Success Login : ", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlSelectWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	script := "SELECT username FROM user WHERE username = ? 	AND PASSWORD = ? LIMIT 1"
	fmt.Println(script)
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
		fmt.Println("Success Login : ", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlInsertWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "petrus"
	password := "petrus"
	script := "INSERT INTO user (username , password) VALUES (? , ? )"
	fmt.Println(script)
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

}
func TestSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	email := "anggiatpangaribuan12@gmail.com"
	comment := "coba bang"
	script := "INSERT INTO comments (email , comment) VALUES (? , ? )"
	fmt.Println(script)
	res, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Lasted Id Inserted ", insertedId)

}
func TestSqlPreparedStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "INSERT INTO comments (email , comment) VALUES (? , ? )"
	fmt.Println(script)
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for i := 0; i < 12; i++ {
		email := "anggiatpangaribuan" + strconv.Itoa(i) + "@gmail.com"
		komentar := "Komentar ke " + strconv.Itoa(i)

		res, err := stmt.ExecContext(ctx, email, komentar)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("success insert", id)
	}

}
func TestSqlTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comments (email , comment) VALUES (? , ? )"
	for i := 0; i < 12; i++ {
		email := "anggiatpangaribuan" + strconv.Itoa(i) + "@gmail.com"
		komentar := "Komentar ke " + strconv.Itoa(i)

		res, err := tx.ExecContext(ctx, script, email, komentar)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("success insert", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}
