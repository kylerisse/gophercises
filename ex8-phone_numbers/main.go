package main

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	// mysql driver for db
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	test := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"Q49709",
		"ABCDEF9999",
		"(123)456-7892",
	}

	db, err := sql.Open("mysql", "root:password@/numbers")
	if err != nil {
		panic(err)
	}

	for _, t := range test {
		rez, err := normalize(t)
		if err != nil {
			println(err.Error())
			continue
		}
		db.Query("INSERT INTO numbers (numbers) VALUES (" + strconv.Itoa(rez) + ");")
	}

}

func normalize(number string) (int, error) {
	number = strings.ReplaceAll(number, "(", "")
	number = strings.ReplaceAll(number, ")", "")
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, " ", "")
	if len(number) != 10 {
		return -1, errors.New(number + " bad length ")
	}
	ret, err := strconv.Atoi(number)
	return ret, err
}
