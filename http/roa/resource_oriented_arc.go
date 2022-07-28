// resource oriented architecture(ROA) is a style of software architecture
// 面向资源的架构（ROA）是一种软件架构风格
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["mike"] = &Employee{"e-1", "Mike", 20}
	employeeDB["rose"] = &Employee{"e-2", "Rose", 20}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)

	if info, ok = employeeDB[qName]; !ok {
		w.Write([]byte("{\"error\":\"Not found\"}"))
		return
	}
	// 使用 reflect 解析
	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
		return
	}
	w.Write(infoJson)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)
	log.Fatal(http.ListenAndServe(":8080", router))
}
