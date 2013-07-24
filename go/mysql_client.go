package main

import (
	"fmt"
	"database/sql"
	"reflect"
)

import _ "github.com/go-sql-driver/mysql"

type ResultRow map[string]interface{}

type const_master struct {
	const_name string
	event_id int
	const_value float64
}

func getDriver() string {
	return "mysql";
}

func getConf() string {
	// TODO: 
	// read user:pass@/db
	// from config file
	return "root:@/test" 
}

func GetHandle() (*sql.DB) {
	// TODO:
	// get handle from handle pool
	// if there is not a available handle
	// then create one
	dbh, err := sql.Open(getDriver(), getConf())
	if err != nil {
		panic(err.Error())
	}
	return dbh
}


func QueryExec(dbh *sql.DB, stmt string, args ...interface{}) (*sql.Rows) {
	sth, err := dbh.Prepare(stmt)
	//fmt.Println(stmt, args)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := sth.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	return rows
}

func GetResult(rows *sql.Rows) ([]interface{}, []string){
	if rows == nil {
		panic("rows is null")
	}
	columns, err := rows.Columns()
	
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	values := make([]interface{}, len(columns))
	scan_args := make([]interface{}, len(columns))
	for i := range values {
		scan_args[i] = &values[i]
	}

	var ret []interface{}
	for rows.Next() {
		err = rows.Scan(scan_args...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		//TODO
		//row := makeRow(columns, scan_args...)
		//row := make(ResultRow, len(columns))
		//var row ResultRow
		row := ResultRow{}
		for i, col_name := range columns {
			row[col_name] = values[i];
		}
		ret = append(ret, row);
	}
	return ret, columns
}

//func makeRow() {

//}

func convert(v interface{}) {
}

func Dump(results []interface{}, columns []string) {
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()

	for _, row := range results {
		rr := row.(ResultRow)
		//fmt.Println(reflect.TypeOf(rr), " : ", reflect.ValueOf(rr))
		//fmt.Println(reflect.ValueOf(rr).Kind())
		for _, field := range rr {
			switch val := field.(type) {
				case []byte:
					fmt.Printf("%s\t", string(val))
				default:
					fmt.Printf("%+v\t", val)
			}
		}
		fmt.Println();
	}
}


func main() {
	dbh  := GetHandle();
	// TODO: replace this statment 
	defer dbh.Close()
/*
	stmt := `
	SELECT	squareNumber 
	  FROM	squarenum 
	  WHERE	number = ?
`*/
	stmt := `
	SELECT	* 
	  FROM	const_master
`

	//rows := QueryExec(dbh, stmt, 12)
	rows := QueryExec(dbh, stmt)

	results, columns := GetResult(rows)

	//Dump(results, columns)
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()

	for _, row := range results {
		rr, ok := row.(const_master)
		if ok {
			fmt.Println(rr);
		} else {
			//panic("const_master convertion failed")
		}
	}


	// set result rows into correspond data structures
	// Target row struct
	// type tgt_row struct {
	// 		some defination ...
	// }
	// var tgt_row[100] tr
	// rows.Scan(&tgt_row)

}
