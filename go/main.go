package main

import (
	"fmt"
	"strings"
//	"math"
//	"net/http"
)

//Exercise for Slice
func Pic(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		
		//对于数组元素赋值用 = 而不是 :=
		array[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			array[i][j] = uint8(i^j)
			fmt.Printf("i=%d, j=%d, array[i][j]=%d\n", i,j,array[i][j])
		}
	}
	//return array
	return  array
}

//Exercies for Map
func WordCount(s string) map[string]int {
	
	array := strings.Fields(s);
	m := make(map[string]int);
	for _, v := range array {
		m[v]++;
	}
	//return map[string]int{"x": 1}
	return m;
}

// Exercise for Clousure
func fibonacci () func() int {
	pre := 0
	cur := 1
	
	return func() int {
		cur = pre + cur	// cur = cur + pre = next
		pre = cur - pre // pre = next - pre = cur + pre - pre = cur
		return pre
	}
}
/*
type Hello struct {}

type Struct string

type Struct struct {
	Greeting 	string
	Punct		string
	Who			string
}

func (h Hello) ServeHTTP (w http.ResponseWriter, r*http.Request) {
	fmt.Fprint(w, "hello!")
}
*/

func main() {
	//pic.Show(Pic)
	//Pic(5, 3)
	/*
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	*/
	//http.Handle("/string", String("I'm a frayed knot."))
	//http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//var h Hello
	//http.ListenAndServe("localhost:4000", h);
	var a, b, c int = 1, 2, 3
	a, b, c = c, b, a
	fmt.Println("Result", a,b,c)
}

