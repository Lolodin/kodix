package route

import (
	"encoding/json"
	"fmt"
	"github.com/Lolodin/kodix.git/internal/store"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var statusMap = map[string]string{
	"0": "В пути",
	"1": "На складе",
	"2": "Продан",
	"3": "Снят с продажи",
}
type requestError struct {
	Error string `json:"error"`
	Status string `json:"status"`
}



//Get
func GetProduct (m  store.Memory) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		a := strings.Split(request.URL.Path, "/")[2]
		p, e := m.Get(a)
		if e!= nil {
			req:= requestError{}
			req.Error = e.Error()
			req.Status = "Error"
			b, _:=json.Marshal(req)
			writer.Write(b)
			return
		}
		writer.Write(p.ToJson())
	}
}
//Update
func UpdateProduct (m  store.Memory) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		req:= requestError{}
		fmt.Println("Save Product")
		p:= store.Product{}
		buffBody, e:= ioutil.ReadAll(request.Body)
		if e!= nil {
			fmt.Println(e)
		}
		e=json.Unmarshal(buffBody,&p)
		fmt.Println(p)
		if e!= nil {
			req:= requestError{}
			req.Status = "Error"
			req.Error = fmt.Sprint(e)

			b, e:= json.Marshal(req)
			if e != nil {
				fmt.Println(e)
			}
			writer.Write(b)
			return
		}
		p.Status = statusMap[p.Status]
		err:=m.Update(p)
		req.Status = "Ok"
		if err != nil {
			req.Status = "error"
			req.Error = fmt.Sprint(err)
			b, e:= json.Marshal(req)
			if e != nil {
				fmt.Println(e)
			}
			writer.Write(b)
			return
		}
		b, e:= json.Marshal(req)
		if e != nil {
			fmt.Println(e)
		}
		writer.Write(b)

	}
}
//Delete
func DeleteProduct (m  store.Memory) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		a := strings.Split(request.URL.Path, "/")[2]
		m.Remove(a)
	}
}
//PUT
func CreateProduct (m  store.Memory) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		req:= requestError{}
		fmt.Println("Save Product")
		p:= store.Product{}
		buffBody, e:= ioutil.ReadAll(request.Body)
		if e!= nil {
			fmt.Println(e)
		}
		e=json.Unmarshal(buffBody,&p)
		if e!= nil {
			req:= requestError{}
			req.Status = "Error"
			req.Error = fmt.Sprint(e)

			b, e:= json.Marshal(req)
			if e != nil {
				fmt.Println(e)
			}
			writer.Write(b)
			return
		}
		req.Status = "Ok"
		b, e:= json.Marshal(req)
		if e != nil {
			fmt.Println(e)
		}
		p.Status = statusMap[p.Status]
		m.Set(p)
		writer.Write(b)


	}
}
func ProductList (m  store.Memory) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
	list := store.NewList(m)
	b, e := json.Marshal(list)
	if e!=nil {
		fmt.Println(e)
	}
	writer.Write(b)
	}
}
func IndexHandler () func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("kodix-app/build/index.html")
		err := t.Execute(writer, "index")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}