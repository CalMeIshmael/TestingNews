package main

import "fmt"
import "github.com/caser/gophernews"
import "net/http"
import "io/ioutil"
import "strings"
import "strconv"

/*
**	Since the Golang interpretation of the hackernewsapi doesn't include a function to get the newest stories, I wrote my own.
**	The function makes an HTTP GET request to the hackernews api for its recent stories, and then Parses the string of keys returned into
**	a int64[] structure that can actually be used.
 */

func ParseJson(string Json) []int {
	stringSplit := strings.Split(Json, ",")
	for index, elem := range stringSplit {
		if index == 0 {
			o, err







}




func getNewStories() []int {
	var numbers []int
	resp, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	stringBody := string(body[:]) //initializes new story key values as string
	newStringBody := strings.Split(stringBody, ",")
	for index, elem := range newStringBody {
		if index == 0 {
			i, err := strconv.ParseInt(strings.Replace(elem, "[", "", -1), 10, 32)
			if err == nil {
				numbers = append(numbers, int(i))
			}
		} else if index == len(newStringBody)-1 {
			i, err := strconv.ParseInt(strings.Replace(elem, "]", "", -1), 10, 32)
			if err == nil {
				numbers = append(numbers, int(i))
			}
		} else {
			i, err := strconv.ParseInt(elem, 10, 32)
			if err == nil {
				numbers = append(numbers, int(i))
			}
		}
	}

	return numbers
}

func main() {
	client := gophernews.NewClient()
	newStories := getNewStories()
	for index := range newStories {
		story, _ := client.GetStory(newStories[index])
		fmt.Println(story.Title)
		index++
	}
}
