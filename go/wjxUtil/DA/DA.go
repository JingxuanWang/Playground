package DA

import (
	"fmt"
	"database/sql"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)

type ResultRow map[string]interface{}

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
	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
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
			panic(err.Error())
		}

		row := ResultRow{}
		for i, col_name := range columns {
			row[col_name] = values[i];
		}
		ret = append(ret, row);
	}
	return ret, columns
}

func ConvertResult(src interface{}, dst interface{}) {
	rr, ok := src.(ResultRow)
	if !ok {
		panic("result isn't match")
	}
	v := reflect.ValueOf(dst)
	e := v.Elem()
	for name, elem := range rr {
		if e.Kind() == reflect.Struct {
			f := e.FieldByName("Field_" + name)
			if f.IsValid() {
				if f.CanSet() {
					k := f.Kind()
					if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
						elemInt, _ := elem.(int64)
						f.SetInt(int64(elemInt))
					} else if k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 {
						elemUint, _ := elem.(uint64)
						f.SetUint(uint64(elemUint))
					} else if k == reflect.Float32 || k == reflect.Float64 {
						elemFloat, _ := elem.(float64)
						f.SetFloat(float64(elemFloat))
					} else if k == reflect.Bool {
						elemBool, _ := elem.(bool)
						f.SetBool(bool(elemBool))
					} else if k == reflect.String {
						elemStr, _ := elem.([]byte)
						f.SetString(string(elemStr))
					} else {
						fmt.Println("type mismatch: ", k)
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
