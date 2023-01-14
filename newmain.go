package main

// concurrented code

/*

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
)

var redditSession *geddit.LoginSession

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
	inPutProRedditCh := make(chan Story, 10)
	inPutITRedditCh := make(chan Story, 10)

	outConsoleCh := make(chan Story, 10)
	outFileCh := make(chan Story, 10)

	go loadReddit("programming", inPutProRedditCh)
	go loadReddit("IT", inPutITRedditCh)

	file, err := os.Create("stories.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	go outConsole(outConsoleCh)
	go outFile(outFileCh, file)

	// flags
	proRedditOpen := true
	itRedditOpen := true

	for proRedditOpen || itRedditOpen {
		select {
		case story, open := <-inPutProRedditCh:
			if open {
				outConsoleCh <- story
				outFileCh <- story
			} else {
				proRedditOpen = false
			}
		case story, open := <-inPutITRedditCh:
			if open {
				outConsoleCh <- story
				outFileCh <- story
			} else {
				itRedditOpen = false
			}
		}
	}

}

// To use concurrence in the code, we need to enter any data that is received into a channel to be managed. So we put the final data in a channel.
func loadReddit(filter string, storiesCh chan<- Story) {
	defer close(storiesCh)
	// stories are channel by type Story

	var listOption geddit.ListingOptions
	submission, err := redditSession.SubredditSubmissions(filter, "", listOption)
	if err != nil {
		fmt.Println(err)
	}

	for _, s := range submission {
		newStory := Story{
			title: s.Title,
			url:   s.URL,
		}
		storiesCh <- newStory
	}

}

func outConsole(storyCh <-chan Story) {
	for {
		s := <-storyCh
		fmt.Printf("%s \n \t %s \n\n", s.title, s.url)
	}
}

func outFile(storyCh <-chan Story, file *os.File) {
	for {
		s := <-storyCh
		fmt.Fprintf(file, "%s \n \t %s \n\n", s.title, s.url)
	}
}

*/
