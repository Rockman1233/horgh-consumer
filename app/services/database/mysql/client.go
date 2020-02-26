package mysql

import (
	"context"
	"fmt"
	"horgh-consumer/app/entities"
	"strings"
)

type Implementation struct {
}

func (i Implementation) Insert(ctx context.Context, queryObj entities.Query) error {

	var keyString strings.Builder
	var valueString strings.Builder
	table := "user" // TODO:: Where is this from?
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

	conn, _ := Client.Connect("mysql", "root:my-secret-pw@(0.0.0.0:3306)/dc_image")
	_, _ = conn.Exec(insertingQuery)

	return i.client.Insert(ctx)
}

func (i Implementation) Delete(ctx context.Context) error {
	return nil
}

func (i Implementation) Update(ctx context.Context) error {
	return nil
}

func New(conf Config) Implementation {
	return Implementation{}
}
