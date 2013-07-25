package main

import (
	"fmt"
	"wjxUtil/DA"
)

// Fields in sturct should be defined to be able to exported
// So first letter should be Capitalized
// It is recommended that using Field_{col_name} style
type const_master struct {
	Field_const_name string
	Field_event_id int64
	Field_const_value float64
}

type test_client struct {
	Field_uint_col uint64
	Field_int_col int64
	Field_tiny_col int8
	Field_bigint_col int64
	Field_float_col float64
	Field_bool_col bool
	Field_varchar_col string
}

type squareNum struct {
	Field_number int64
	Field_squareNumber int64
}

func main() {
	dbh  := DA.GetHandle();
	// TODO: replace this statment 
	defer dbh.Close()
/*
	stmt := `
	SELECT	squareNumber 
	  FROM	squarenum 
	  WHERE	number = ?
`
*/
	stmt := `
	SELECT	*
	  FROM	test_client
`
	//rows := QueryExec(dbh, stmt, 9)
	rows := DA.QueryExec(dbh, stmt)

	results, columns := DA.GetResult(rows)

	DA.Dump(results, columns)
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()

	//res := make([]squareNum, len(results))
	res := make([]test_client, len(results))
	for i, row := range results {
		DA.ConvertResult(row, &res[i])
	}
	fmt.Println(res)
}
