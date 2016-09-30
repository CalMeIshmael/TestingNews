package main

import "net/http"
import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

//import "log"

func main() {
	var numbers[]int64
	resp, _ := http.Get("https://hacker-news.firebaseio.com/v0/newstories")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	
	stringBody := string(body[:])	//initializes new story key values as string 
	newStringBody := strings.Split(stringBody, ",")
	for index, elem := range newStringBody {
		if index == 0 {
			i, err := strconv.ParseInt(strings.Replace(elem, "[", "", -1), 10, 64)
			if err == nil {
				numbers = append(numbers, i)
			}
		} else if index == len(newStringBody)-1 {
			i, err := strconv.ParseInt(strings.Replace(elem, "]", "", -1), 10, 64)
			if err == nil {
				numbers = append(numbers, i)
			}
		} else {
			i, err := strconv.ParseInt(elem, 10, 64)
			if err == nil {
				numbers = append(numbers, i)
			}
		}
	}
	resp.Body.Close()
	fmt.Println(numbers)


}
