package main

import (
	"os"
	"context"
	"log"
	"cloud.google.com/go/spanner"
	"encoding/csv"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: ./spanner2csv <DSN> <SQL>")
		return
	}

	dsn := os.Args[1]
	sql := os.Args[2]

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	stmt := spanner.NewStatement(sql)
	iter := client.Single().Query(ctx, stmt)

	rowCount := 0
	w := csv.NewWriter(os.Stdout)
	if err := iter.Do(func (row *spanner.Row) error {
		colNames := row.ColumnNames()
		if rowCount == 0 {
			if err := w.Write(colNames); err != nil {
				return err
			}
		}
		values := make([]spanner.GenericColumnValue, len(colNames))
		valuePtrs := make([]interface{}, len(colNames))
		for i := 0; i < len(colNames); i++ {
			valuePtrs[i] = &values[i]
		}
		err := row.Columns(valuePtrs...)
		if err != nil {
			return err
		}
		strValues := make([]string, len(values))
		for i := 0; i < len(values); i++ {
			strValues[i] = values[i].Value.GetStringValue()
		}
		if err := w.Write(strValues); err != nil {
			return err
		}
		rowCount++
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
	w.Flush()
}
