package main

/*

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
)

// recieve data from Reddit by an object named redditSession
// Its type is a type of pointer to LoginSession
// This object must be set in order to connect to the site
var redditSession *geddit.LoginSession

// story is news ,at the moment we just need title and url
type Story struct {
	title string
	url   string
}

func init() {
	var err error
	redditSession, err = geddit.NewLoginSession("g_d_bot", "K417k4FTua52", "gdAgent v0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	// print stories by function loadReddit
	stories := loadReddit()
	fmt.Printf("%v \n\n", stories)

}

// function for calling Reddit site and recive data from it
func loadReddit() []Story {
	var stories []Story

	var listOption geddit.ListingOptions
	submission, err := redditSession.SubredditSubmissions("programming", "", listOption)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, s := range submission {
		story := Story{
			title: s.Title,
			url:   s.URL,
		}
		stories = append(stories, story)
	}

	return stories
}

*/
