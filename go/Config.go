package main;

import (
	"fmt"
	//"strings"
	"io/ioutil"
	"encoding/json"
)

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

func LoadRouter() {
	fileBytes, err := ioutil.ReadFile(routerConf)
	if  err == nil {
		conf := new(config)
		err := json.Unmarshal(fileBytes, &conf);
		fmt.Printf("%v", conf)
		fmt.Println(err)
	} else {
		fmt.Println("Error in read file ", err)
	}
}

func LoadConfig() {
	fileBytes, err := ioutil.ReadFile(commonConf)
	if  err == nil {
		//s := string(fileBytes)
		//strs := strings.Split(s, "\n")
		//for i := 0; i < len(strs); i++ {
		//	fmt.Println(i + 1, " ", strs[i])
		//}
		//var conf interface{}
		conf := new(config)
		err := json.Unmarshal(fileBytes, &conf);
		fmt.Printf("%v", conf)
		fmt.Println(err)
	} else {
		fmt.Println("Error in read file ", err)
	}
}

func main() {
	//LoadConfig()
	LoadRouter()
}
