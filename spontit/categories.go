package spontit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Categories contains a list of categories
type Categories struct {
	Data []*Category `json:"data"`
}

// Category contains a single category
type Category struct {
	Code  float64 `json:"categoryCode"`
	Title string  `json:"categoryTitle"`
}

// Categories returns available categories
func (c *Client) Categories() (*Categories, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories", baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	categories := new(Categories)
	json.Unmarshal(content, categories)

	return categories, nil
}
