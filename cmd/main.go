package main

import (
	"fmt"
	"github.com/felipebool/bloom/filter"
	"github.com/felipebool/bloom/function"
	"log"
)

func getStringList() []string {
	return []string{
		"It’s been a long day, and the last thing you want to do is start",
		"studying your German vocabulary and grammar.Still, you don’t want",
		"to break your streak.After all, learning every day is a good habit",
		"to have.If you could find a way to absorb some of the language without",
		"breaking open your flash cards? Music to your ears! By which we mean,",
		"literally, music to your ears, because we’re recommending German Spotify",
		"playlists for you to listen to as you learn German.",
		"While it’s difficult to learn a language solely through music, there are",
		"benefits to listening to songs in your target language as a supplement to",
		"your usual studying.With music, you can pick up some pronunciation,",
		"discover new vocab and learn about German musical culture.",
		"Jumping right in blind can be a bit difficult, however, because some",
		"songs are better for learning than others.We compiled a list of five",
		"German Spotify playlists that range from easy to difficult so you can",
		"choose what’s best for you based on your experience.",
	}
}

// 15 strings
// 8 MUST NOT be in the filter
// 7 MAY BE in the filter : )
func getTestStringList() []string {
	return []string{
		"It’s been a long day, thing you want to do is start",						// X
		"studying your German vocabulary and grammar.Still, you don’t want",		//
		"to break your streak. learning every day is a good habit",					// X
		"to have.If you could find a way to absorb some of the language without",	//
		"breaking open your ears! By which we mean,",								// X
		"literally, music to your ears, because we’re recommending German Spotify", //
		"playlists for to as you learn German.",									// X
		"While it’s difficult to learn a language solely through music, there are",	//
		"benefits to as a supplement to",											// X
		"your usual studying.With music, you can pick up some pronunciation,",		//
		"discover new vocab and learn culture.",									// X
		"Jumping right in blind can be a bit difficult, however, because some",		//
		"songs are others.We compiled a list of five",								// X
		"German Spotify playlists that range from easy to difficult so you can",	//
		"choose what’s you based on your experience.",								// X
	}
}

func main() {
	var present, notPresent int

	hashFunctions := []func(string) (uint32, error){
		function.HashFNV(),
		function.HashFNVa(),
	}

	bFilter := filter.NewBloom(
		200,
		hashFunctions,
	)

	for _, str := range getStringList() {
		if err := bFilter.AddString(str); err != nil {
			log.Fatal(err)
		}
	}

	for _, str := range getTestStringList() {
		isPresent, _ := bFilter.CheckString(str)
		if isPresent {
			present++
			continue
		}

		notPresent++
	}

	fmt.Printf("Total of present strings: %d\n", present)
	fmt.Printf("Total of not present strings: %d\n", notPresent)
}
