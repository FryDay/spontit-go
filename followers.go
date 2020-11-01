package spontit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Followers contains a list of followers
type Followers struct {
	Data []string `json:"data"`
}

type followersRequest struct {
	ChannelName string `json:"channelName,omitempty"`
}

// Followers gets the list of usernames that follow a specific channel.
// Pass an empty string to return followers of your main channel.
func (c *Client) Followers(channelName string) (*Followers, error) {
	reqJSON, err := json.Marshal(&followersRequest{ChannelName: channelName})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/followers", baseURL), bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Authorization", c.config.key)
	req.Header.Add("X-UserId", c.config.userID)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	followers := new(Followers)
	json.Unmarshal(content, followers)

	return followers, nil
}
