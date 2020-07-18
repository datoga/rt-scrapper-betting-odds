package main

import (
	"encoding/json"
	"log"

	"github.com/itchyny/gojq"
)

func parseJSON(resp []byte) ([]string, error) {
	query, err := gojq.Parse(`
		.markets[] | 
		select (.id == "55405316").selections[] | 
		.name + " " + (.trueOdds|tostring)
	`)

	if err != nil {
		return []string{}, err
	}

	var jsonObj map[string]interface{}

	err = json.Unmarshal(resp, &jsonObj)

	if err != nil {
		return []string{}, err
	}

	iter := query.Run(jsonObj)

	results := []string{}

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			log.Fatalln(err)
		}
		results = append(results, v.(string))
	}

	return results, nil
}
