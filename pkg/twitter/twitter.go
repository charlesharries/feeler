package twitter

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// User is a Twitter user.
type User struct {
	ID       string
	Name     string
	Username string
}

// Tweet represents a single tweet belonging to a user.
type Tweet struct {
	ID   string
	Text string
}

// Client is a Twitter API client.
type Client struct {
	BaseURL string
	Token   string
	Cache   *CacheClient
}

// NewClient instantiates a new Twitter client.
func NewClient(token string) *Client {
	return &Client{
		BaseURL: "https://api.twitter.com/2",
		Token:   token,
		Cache:   newCache("./data/cache"),
	}
}

// GetTweetsForUserID fetches a list of tweets for the given user ID.
func (c *Client) GetTweetsForUserID(userID string, limit int) ([]Tweet, error) {
	// Get first page.
	tweets, next, err := c.getTweets(userID, "")

	// Get ensuing pages.
	for {
		if len(tweets) > limit {
			break
		}

		ts, n, err := c.getTweets(userID, next)
		if err != nil {
			return []Tweet{}, err
		}

		next = n
		tweets = append(tweets, ts...)
	}

	return tweets, err
}

// GetTweetsForUsername fetches a list of tweets for the given username.
func (c *Client) GetTweetsForUsername(username string, limit int) ([]Tweet, error) {
	user, err := c.GetUserByUsername(username)
	if err != nil {
		return []Tweet{}, err
	}

	// Get first page.
	tweets, next, err := c.getTweets(user.ID, "")

	// Get ensuing pages.
	for {
		if len(tweets) >= limit {
			break
		}

		ts, n, err := c.getTweets(user.ID, next)
		if err != nil {
			return []Tweet{}, err
		}

		next = n
		for _, t := range ts {
			if len(tweets) >= limit {
				break
			}

			tweets = append(tweets, t)
		}
	}

	return tweets, err
}

func (c *Client) getTweets(userID, token string) ([]Tweet, string, error) {
	// Check that userID is numeric
	_, err := strconv.Atoi(userID)
	if err != nil {
		return []Tweet{}, "", errors.New("userID must be numeric")
	}

	url := "/users/" + userID + "/tweets"
	url = url + "?max_results=100"
	if len(token) > 0 {
		url = url + "&pagination_token=" + token
	}

	body, err := c.get(url)
	if err != nil {
		return []Tweet{}, "", err
	}

	var res TweetsResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return []Tweet{}, "", err
	}

	return res.Tweets, res.Meta.Next, nil
}

// GetUserByUsername fetches a user from the Twitter API by their username.
func (c *Client) GetUserByUsername(username string) (User, error) {
	body, err := c.get("/users/by/username/" + username)
	if err != nil {
		return User{}, err
	}

	var res UserByUsernameResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return User{}, err
	}

	return res.User, nil
}

// get runs an actual request against the Twitter API.
func (c *Client) get(endpoint string) ([]byte, error) {
	// If it exists in cache, return the cached one.
	val, found := c.Cache.Get(endpoint)
	if found {
		// []byte is json.Marshalled to a base64-encoded string, so we need
		// to decode it here to make it useful.
		return base64.StdEncoding.DecodeString(val.(string))
	}

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	c.Cache.Set(endpoint, body)

	return body, nil
}
