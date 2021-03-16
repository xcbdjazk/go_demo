package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5/esapi"
	"log"
)

func main() {
	e, _ := es.NewDefaultClient()

	q := map[string]interface{}{
		"name": "Geo-point as an object",
		"location": map[string]interface{}{
			"lat": "41.12",
			"lon": -71.34,
		},
	}
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(q)
	fmt.Println(b.String())
	res, _ := e.Create("sss", "1", &b, e.Create.WithDocumentType("restaurant"))
	get_errMsg(res)
}

func Query() {
	e, _ := es.NewDefaultClient()
	query := map[string]interface{}{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(query)
	res, err := e.Search(
		e.Search.WithContext(context.Background()),
		e.Search.WithIndex("attractions"),
		e.Search.WithBody(&buf),
		e.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	get_errMsg(res)
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(r)
}

func get_errMsg(res *esapi.Response) {
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
}
