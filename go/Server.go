package main

import (
	//"html"
	"io"
	"log"
	//"time"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var Address string = "127.0.0.1"
var Port uint = 8080

type router struct {
	Url    string `json:"url"`
	Module string `json:"module"`
	Func   string `json:"func"`
}

type config struct {
	AppName string   `json:"app"`
	Version string   `json:"version"`
	Routers []router `json:"router"`
}

var commonConf string = "conf.json"
var routerConf string = "router.conf"

type controllerFunc func(params map[string]interface{}) (ret map[string]interface{})

type RouteTable map[string]controllerFunc

func testControllerFunc(params map[string]interface{}) (ret map[string]interface{}) {
	fmt.Println("Calling testControllerFunc")
	fmt.Printf("Params: %v", params)
	return nil
}
func testControllerFunc2(params map[string]interface{}) (ret map[string]interface{}) {
	fmt.Println("Calling testControllerFunc2")
	fmt.Printf("Params: %v", params)
	return nil
}

type Controller struct {
	rt RouteTable
}
var c Controller

func (c Controller) Regist(url string, function controllerFunc) {
	if ok := c.Match(url); ok == nil {
		c.rt[url] = function
	} else {
		fmt.Println("Already Registered for url: " + url)
	}
}

func (c Controller) Match(url string) (function controllerFunc) {
	if c.rt == nil {
		return nil
	} else {
		return c.rt[url]
	}
}

// common handle function
// handles all requests
func handle(w http.ResponseWriter, r *http.Request) {
	// parse request
	// construct response

	// figure out the Controller/Action

	host := r.URL.Host
	path := r.URL.Path
	query := r.URL.RawQuery
	params, _ := url.ParseQuery(query)

	// execute Controller->Action and get result
	fmt.Printf("%v", params)
	controllerFunc := c.Match(path)
	if controllerFunc != nil {
		controllerFunc(nil)
	}

	// write response
	io.WriteString(w, "hello host: " + host + " path: " + path + " query: " + "<br>")
	io.WriteString(w, "result: <br>")
	//io.WriteString(w, res)

	// calc execute time
	// log
}

func Run() {
	address := fmt.Sprintf("%s:%d", Address, Port)
	s := &http.Server{
		Addr:		    address,
		Handler:		http.HandlerFunc(handle),
		//ReadTimeout:	10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func LoadConfig() {
	fileBytes, err := ioutil.ReadFile(commonConf)
	if  err == nil {
		conf := new(config)
		err := json.Unmarshal(fileBytes, &conf);
		fmt.Printf("%v", conf)
		fmt.Println(err)
	} else {
		fmt.Println("Error in read file ", err)
	}
	//c := Controller {}
	c.rt = make(map[string]controllerFunc)
	c.Regist("/foo", testControllerFunc)
	c.Regist("/bar", testControllerFunc2)
	fmt.Printf("%v", c.rt)
}

func Init() {
	// do initialize
	// load config
	LoadConfig()
	// regist handle function
	//for i, elem := range Config.Routers {
	//	RouteTable[name] = 
	//}
	// run startup hooks
	// etc
}

func main() {
	Init()
	Run()
}

