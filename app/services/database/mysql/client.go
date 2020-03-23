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
	Keys    strings.Builder
	Values  strings.Builder
}

func (i Implementation) Insert(ctx context.Context, queryObj entities.Query) error {

	var keyString strings.Builder
	var valueString strings.Builder

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

	instanceOfQuery := insertingQuery{
		Command: "INSERT INTO", // TODO:: Does is good?
		Table:   queryObj.Table,
		Keys:    keyString,
		Values:  valueString,
	}

	execute := fmt.Sprintf("%s %s (%s) VALUES (%s)", instanceOfQuery.Command, instanceOfQuery.Table, instanceOfQuery.Keys.String(), instanceOfQuery.Values.String())

	_, err := i.db.Exec(execute)

	return err
}

func (i Implementation) Delete(ctx context.Context) error {
	return nil
}

func (i Implementation) Update(ctx context.Context) error {
	return nil
}

func New(conf Config) Implementation {
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)", conf.User, conf.Password, conf.Host, conf.Port)
	db, _ := sql.Open("mysql", dataSource)
	return Implementation{db}
}
