package spontit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Push contains attributes for a push notification
type Push struct {
	// To provide text in a push, supply one of either "content" or "pushContent" (or both).
	// Limited to 2500 characters. (Required if a value for "pushContent" is not provided).
	Content string `json:"content,omitempty"`
	// If you want to control exactly what shows when the notification pops up, then provide a value for "pushContent".
	// Limited to 100 characters. The value provided for "pushContent" will appear when the notification first pops up.
	// Once the user opens the notification, they will then see the value provided for "pushContent" and for "content" (if any).
	// (Required if a value for "content" is not provided).
	PushContent string `json:"pushContent,omitempty"`
	// The title of push. Appears in bold at the top. Limited to 100 characters.
	PushTitle string `json:"pushTitle,omitempty"`
	// The name of a channel you created.
	// If you have not yet created a channel, simply don't provide this value and the push will be sent to your main account.
	ChannelName string `json:"channelName,omitempty"`
	// The subtitle of your push. Limited to 20 characters. Only appears on iOS devices.
	Subtitle string `json:"subtitle,omitempty"`
	// A link that can be attached to your push. Must be a valid URL.
	Link string `json:"link,omitempty"`
	// An array of userIds to whom to send the notification.
	// If all three attributes 'pushToFollowers', 'pushToPhoneNumbers' and 'pushToEmails' are not supplied, then everyone who follows the channel will receive the push notification.
	// If 'pushToFollowers' is supplied, only those listed in the array will receive the push notification.
	// If one of the userIds supplied does not follow the specified channel, then that userId value will be ignored.
	PushToFollowers []string `json:"pushToFollowers,omitempty"`
	// An array of phoneNumbers to whom to send the notification.
	// If all three attributes 'pushToFollowers', 'pushToPhoneNumbers' and 'pushToEmails' are not supplied, then everyone who follows the channel will receive the push notification.
	// If 'pushToPhoneNumbers' is supplied, then we will map the numbers to Spontit accounts and push accordingly.
	// The users specified by 'pushToPhoneNumbers' do not have to follow the specified channel in order to receive the push.
	// However, they can report your push as spam if they do not follow you and do not wish to receive your pushes.
	PushToPhoneNumbers []string `json:"pushToPhoneNumbers,omitempty"`
	// An array of emails to whom to send the notification.
	// If all three attributes 'pushToFollowers', 'pushToPhoneNumbers' and 'pushToEmails' are not supplied, then everyone who follows the channel will receive the push notification.
	// If 'pushToEmails' is supplied, then we will map the emails to Spontit accounts and push accordingly.
	// The users specified by 'pushToEmails' do not have to follow the specified channel in order to receive the push.
	// However, they can report your push as spam if they do not follow you and do not wish to receive your pushes.
	PushToEmails []string `json:"pushToEmails,omitempty"`
	// A Unix timestamp. Schedule a push to be sent at a later date and time.
	Schedule int64 `json:"schedule,omitempty"`
	// A Unix timestamp. When to automatically expire your push notification. The default is 10 days after pushing.
	// The push will become unaccessible within 15-30 minutes of the selected time, but will remain on all device screens until dismissed or clicked.
	ExpirationStamp int64 `json:"expirationStamp,omitempty"`
	// Whether to open the provided link within the iOS app or in Safari. Android PWA opens all links in the default web browser.
	OpenLinkInApp bool `json:"openLinkInApp,omitempty"`
	// Control whether the notification opens to the home feed or to a standalone page with the notification. The default (openInHomeFeed=False) is to open the notification on a standalone page.
	OpenInHomeFeed bool `json:"openInHomeFeed,omitempty"`
	// An iOS deep link. Use this to deep link into other apps. Alternatively, you can provide a universal link in the link attribute and set openLinkInApp to false.
	IOSDeepLink string `json:"iOSDeepLink,omitempty"`
}

// Push sends a push notification
func (c *Client) Push(push *Push) error {
	if push.Content == "" && push.PushContent == "" {
		return fmt.Errorf("Either Content, PushContent, or both must be provided")
	}

	reqJSON, err := json.Marshal(push)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/push", baseURL), bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Authorization", c.config.key)
	req.Header.Add("X-UserId", c.config.userID)

	_, err = c.httpClient.Do(req)
	return err
}
