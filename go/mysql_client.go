package main

import (
	"fmt"
	"database/sql"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)


type ResultRow map[string]interface{}

type const_master struct {
	Field_const_name string
	Field_event_id int
	Field_const_value float64
}

type squareNum struct {
	Field_number int64
	Field_squareNumber int64
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

// should make a dst of equal length to src
func ConvertResult(src interface{}, dst interface{}) {
	rr, ok := src.(ResultRow)
	if !ok {
		panic("result isn't match")
	}
	v := reflect.ValueOf(dst)
	e := v.Elem()
	for name, elem := range rr {
		elemInt, _ := elem.(int64)
		if e.Kind() == reflect.Struct {
			f := e.FieldByName("Field_" + name)
			if f.IsValid() {
				if f.CanSet() {
					if f.Kind() == reflect.Int64 {
						f.SetInt(int64(elemInt))
					} else {
						fmt.Println("field kind is not int")
					}
				} else {
					fmt.Println("field can not set")
				}
			} else {
				fmt.Println("field is not valid")
			}
		} else {
			fmt.Println("elem kind is not struct")
		}
	}
}


func Dump(results []interface{}, columns []string) {
	return
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()

	for _, row := range results {
		rr := row.(ResultRow)
		fmt.Println(reflect.TypeOf(rr), " : ", reflect.ValueOf(rr))
		fmt.Println(reflect.ValueOf(rr).Kind())
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
`
*/
	stmt := `
	SELECT	* 
	  FROM	squareNum
	  WHERE number < 3
`
	//rows := QueryExec(dbh, stmt, 3)
	rows := QueryExec(dbh, stmt)

	results, columns := GetResult(rows)

	Dump(results, columns)
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()

	res := make([]squareNum, len(results))
	for i, row := range results {
		ConvertResult(row, &res[i])
	}
	fmt.Println(res)
}
