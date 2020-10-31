package spontit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Push contains attrubutes for a push notification
type Push struct {
	Content            string   `json:"content,omitempty"`
	PushContent        string   `json:"pushContent,omitempty"`
	PushTitle          string   `json:"pushTitle,omitempty"`
	ChannelName        string   `json:"channelName,omitempty"`
	Subtitle           string   `json:"subtitle,omitempty"`
	Link               string   `json:"link,omitempty"`
	PushToFollowers    []string `json:"pushToFollowers,omitempty"`
	PushToPhoneNumbers []string `json:"pushToPhoneNumbers,omitempty"`
	PushToEmails       []string `json:"pushToEmails,omitempty"`
	Schedule           int64    `json:"schedule,omitempty"`
	ExpirationStamp    int64    `json:"expirationStamp,omitempty"`
	OpenLinkInApp      bool     `json:"openLinkInApp,omitempty"`
	OpenInHomeFeed     bool     `json:"openInHomeFeed,omitempty"`
	IOSDeepLink        string   `json:"iOSDeepLink,omitempty"`
}

// PushResult contains the results of a push notification
type PushResult struct {
	Data pushResultData `json:"data"`
}

type pushResultData struct {
	Message string `json:"message"`
}

// Push sends a push notification
func (c *Client) Push(push *Push) (*PushResult, error) {
	if push.Content == "" && push.PushContent == "" {
		return nil, fmt.Errorf("Either Content, PushContent, or both must be provided")
	}

	reqJSON, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/push", baseURL), bytes.NewBuffer(reqJSON))
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
	log.Println(string(content))

	pushResult := new(PushResult)
	json.Unmarshal(content, pushResult)

	return pushResult, nil
}
