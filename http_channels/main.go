package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {

	// calling http.Get() which gives a response and an error value.
	resp, err := http.Get(url)

	// error handling
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s // sending into the channel
	} else {
		// a good practice is to close the Response Body if there is one
		// If you forget to close it there will be a resource leak.
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d  \n", url, resp.StatusCode)
		c <- s

		// fetching the page if HTTP Status Code is 200 (success)
		if resp.StatusCode == 200 {

			// The resp.Body implements the io.Reader interface
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			// Creating the file's name
			file := strings.Split(url, "//")[1]
			file += ".txt" // and I concatenate .txt to file value

			s := fmt.Sprintf("Writing response Body to %s\n", file)

			// Writing the response Body to File
			// If the file doesn't exist WriteFile() creates it and if it already exists
			// the function will truncate it before writing to file.
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				//log.Fatal(err)
				s += "Error writing file \n"
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}

}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com", "https://www.wikipedia.com"}


	c := make(chan string)


	// Iterating over the URLs and call the function for each URL
	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No of goroutines", runtime.NumGoroutine())

	for i:=0; i<len(urls);i++{
		fmt.Println(<-c)
	}

}