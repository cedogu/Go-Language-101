package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBtyes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBtyes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBtyes, &todo)
	fmt.Println(todo)

}

func Demo2() {

	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBtyes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBtyes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBtyes, &todoResponse)
	fmt.Println(todoResponse)
}
