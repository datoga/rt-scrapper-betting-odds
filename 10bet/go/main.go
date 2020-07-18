package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	log.Println("Calling URL")

	resp, err := getURL()

	if err != nil {
		panic(err)
	}

	log.Println("Parsing json")

	results, err := parseJSON(resp)

	if err != nil {
		panic(err)
	}

	log.Println("JSON parsed")

	strResults := ""

	for _, v := range results {
		strResults += v + "\n"
	}

	log.Println(strResults)
}
