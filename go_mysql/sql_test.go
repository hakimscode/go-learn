package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer (id, name) VALUES ('setiawan', 'Setiawan')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	defer rows.Close()
}

func TestExecSqlNew(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer (id, name, email, balance, rating, birth_date, married) " +
		"VALUES ('setiawan', 'Setiawan', NULL, 90000, 85.0, '1993-10-27', false)"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySqlNew(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name, email sql.NullString
		var balance int32
		var rating float32
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=============================")
		if !email.Valid {
			fmt.Println("Email ini NULL")
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name.String)
		fmt.Println("Email:", email.String)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth Date:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "hakims; DROP TABLE user; #"
	password := "password"

	// ? can manipulate this query if using string concat
	// query := "INSERT INTO user (username, password) VALUES " +
	// 	"('" + username + "', '" + password + "')"

	// ? use query param instead to avoid sql injection
	queryParam := "INSERT INTO user (username, password) VALUES (?, ?)"

	_, err := db.ExecContext(ctx, queryParam, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new User")
}

func TestQuerySqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "heri';#"
	password := "password"

	// ? can manipulate this query if using string concat
	// query := "SELECT username FROM user " +
	// 	"WHERE username = '" + username + "' AND password = '" + password + "'" +
	// 	"LIMIT 1"
	// ? use query param instead to avoid sql injection
	queryParam := "SELECT username FROM user " +
		"WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, queryParam, username, password)

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

		fmt.Println("Sukes login:", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestLastInsertId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "hhakimsetiawan@gmail.com"
	comment := "Halo, apa kabar?"

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "hakim_" + strconv.Itoa(i) + "@gmail.com"
		comment := "Belajar golang yuk"

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment using prepare statement with id", insertId)
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

	// do transaction
	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "hakim_" + strconv.Itoa(i) + "@gmail.com"
		if i == 9 {
			query = "INSERT INTO comments (comment) VALUES (?)"
		}
		comment := "Belajar golang yuk"

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment using prepare statement with id", insertId)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
