package twitter

// UserByUsernameResponse represents a response from the
// /users/by/username/:username endpoint.
type UserByUsernameResponse struct {
	User User `json:"data"`
}

// TweetsResponse represents a response from the /users/:id/tweets
// endpoint.
type TweetsResponse struct {
	Tweets []Tweet `json:"data"`
	Meta   struct {
		Count    int    `json:"result_count"`
		Next     string `json:"next_token"`
		Previous string `json:"previous_token"`
	}
}
