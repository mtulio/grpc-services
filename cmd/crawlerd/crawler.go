package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Crawler(url string) (string, bool) {
	status := false
	bodyString := ""
	response, err := http.Get(url)
	var bodyBytes []byte

	if err != nil {
		// log.Fatal(err)
		return fmt.Sprintf("error, statusCode: %w", err), status
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			bodyString = fmt.Sprintf("error, statusCode: %s", response.StatusCode)
		} else {
			bodyString = string(bodyBytes)
			status = true
		}
	} else {
		bodyString = fmt.Sprintf("error, statusCode: %s", response.StatusCode)
	}

	return bodyString, status
}

// func crawlerService() {
// 	for {
// 		select {
// 		case url <- chCrawlerReq:
// 			chCrawlerResp <- crawler(url)
// 		}
// 	}
// }
