package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var supabaseUrl = os.Getenv("SUPABASE_URL")
var supabaseKey = os.Getenv("SUPABASE_ANON_KEY")

func FetchShortLinks() ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/v1/shortsy", supabaseUrl), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", supabaseKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var result []map[string]interface{}
	json.Unmarshal(body, &result)

	return result, nil
}
