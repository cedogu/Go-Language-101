package restful

import (
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
