package main
 
import (
        "fmt"
        r "reflect"
)
 
type S struct {
        I   int
        Str string
}
 
type error string
 
func (e error) String() string { return string(e) }
 
func main() {
        s := S{5, "Bunnies"}
        m, _ := StructToMap(s)
 
        s1 := S{}
        MapToStruct(m, &s1)
 
        fmt.Printf("%+v\n%v\n%+v\n", s, m, s1)
}
 
func StructToMap(val interface{}) (mapVal map[string]interface{}, ok bool) {
        // indirect so function works with both structs and pointers to them
        structVal, ok := r.Indirect(r.ValueOf(val)).(*r.StructValue)
        if !ok {
                return
        }
        
        typ := structVal.Type().(*r.StructType)
        mapVal = make(map[string]interface{})
 
        for i := 0; i < typ.NumField(); i++ {
                field := structVal.Field(i)
                if field.CanSet() {
                        mapVal[typ.Field(i).Name] = field.Interface()
                }
        }
 
        return
}
 
func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
        structVal, ok := r.Indirect(r.ValueOf(val)).(*r.StructValue)
        if !ok {
                return
        }
 
        for name, elem := range mapVal {
                structVal.FieldByName(name).SetValue(r.ValueOf(elem))
        }
 
        return
}
