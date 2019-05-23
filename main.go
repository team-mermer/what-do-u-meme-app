package main

import (
	"encoding/json"
	"net/http"
	"what-do-u-meme-app/app"
)

func SearchByTextMock(writer http.ResponseWriter, request *http.Request) {
	json_string, _ := json.Marshal("mock search results")
	writer.Write(json_string)
}

func main() {
	// add func to handle url request
	// http.HandleFunc("/get_trending", app.GetTrending)
	http.HandleFunc("/mock/search_by_text", SearchByTextMock)
	http.HandleFunc("/mock/get_trending", app.GetTrendingMock)

	// listen and serve
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}