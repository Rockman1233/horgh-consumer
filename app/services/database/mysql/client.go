package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"horgh-consumer/app/entities"
	"strings"
)

type Implementation struct {
	db *sql.DB
}

type insertingQuery struct {
	Command string
	Table   string
	Keys    string
	Values  string
}

func (i Implementation) Insert(ctx context.Context, queryObj entities.Query) error {

	var keyString strings.Builder
	var valueString strings.Builder
	table := queryObj.Table
	lengthOfMap := len(queryObj.Data)

	// Building strings
	for i, value := range queryObj.Data {
		separator := ","
		if i == (lengthOfMap - 1) {
			separator = ""
		}
		keyString.WriteString(value.Name + separator)
		valueString.WriteString("\"" + fmt.Sprint(value.Value) + "\"" + separator)
	}

	insertingQuery := "INSERT INTO " + table + " (" + keyString.String() + ") VALUES (" + valueString.String() + ")"
	_, err := i.db.Exec(insertingQuery)

	return err
}

func (i Implementation) Delete(ctx context.Context) error {
	return nil
}

func (i Implementation) Update(ctx context.Context) error {
	return nil
}

func New(conf Config) Implementation {
	dataSource := fmt.Sprintf(conf.User, ":", conf.Password, "@(", conf.Host, ":", conf.Port, ")")
	db, _ := sql.Open("mysql", dataSource)
	return Implementation{db}
}
