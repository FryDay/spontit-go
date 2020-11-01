package spontit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Category contains a single category
type Category struct {
	Code  float64 `json:"categoryCode"`
	Title string  `json:"categoryTitle"`
}

type categories struct {
	Data []*Category `json:"data"`
}

// Categories returns available categories
func (c *Client) Categories() ([]*Category, error) {
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

	categories := new(categories)
	json.Unmarshal(content, categories)

	return categories.Data, nil
}
