package main

import (
	"encoding/json"
	"fmt"
	"github"
	"log"
	"omdb"
	"os"
	"time"
	"xkcd"
)

func JSON() {
	type Movie struct {
		Title string
		Year  int `json:"released"`
		// Tag:json(xxx)对应encoding/json(xxx)
		// omitempty:空或零值时不生成
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}

const oneMonthDuration = time.Hour * 24 * 30
const oneYearDuration = time.Hour * 24 * 365

func compareTime(t time.Time) string {
	now := time.Now()

	parsedTime, err := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	if err != nil {
		fmt.Println("无法解析时间：", err)
		return "o"
	}

	onMonthAgo := now.Add(-oneMonthDuration)
	oneYearAgo := now.Add(-oneYearDuration)

	if parsedTime.Before(onMonthAgo) && parsedTime.After(oneYearAgo) {
		return "iy"
	} else if parsedTime.Before(oneYearAgo) {
		return "y"
	} else if parsedTime.After(onMonthAgo) {
		return "m"
	}

	return "o"
}

func getIssues() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	timeMap := make(map[string][]github.Issue, 3)

	for _, issue := range result.Items {
		timeMap[compareTime(issue.CreateAt)] = append(timeMap[compareTime(issue.CreateAt)], *issue)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("less than a month:")
	for _, item := range timeMap["m"] {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("less than a year:")
	for _, item := range timeMap["iy"] {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("more than a year:")
	for _, item := range timeMap["y"] {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	// }
}

func getMoviePosterByTitle() {
	omdb.DownLoadPoster(os.Args[1:])

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(result.Poster)
}

func main() {
	// JSON()
	// getIssues()
	// getMoviePosterByTitle()
	// xkcd.GenerateOfflineComicIndex()
	xkcd.SearchComic(os.Args[1:])
}
