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
	var intJSON []int
	noComma := strings.Split(Json, ",")
	for index, elem := range noComma {
		if index == 0 {
			ParsedElem, err := strconv.ParseInt(strings.Replace(elem, "[", "", -1), 10, 32)
			if err == nil {
				intJSON = append(intJSON, int(ParsedElem))
			}
		} else if index == len(noComma)-1 {
			ParsedElem, err := strconv.ParseInt(strings.Replace(elem, "]", "", -1), 10, 32)
			if err == nil {
				intJSON = append(intJSON, int(ParsedElem))
			}
		} else {
			ParsedElem, err := strconv.ParseInt(elem, 10, 32)
			if err == nil {
				intJSON = append(intJSON, int(ParsedElem))
			}
		}
	}
	return intJSON
}

func NewsHttpGet(Link string) string {
	resp, _ := http.Get(Link)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body[:]) //initializes new story key values as string
}

func getNewStories() []int {
	var NewStories string = "https://hacker-news.firebaseio.com/v0/newstories.json")
	return ParseJSON(NewsHttpGet(NewStories))
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
