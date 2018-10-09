package gqlapi

//StandardQuery provides wrapper for final body sent to AppSync API
type StandardQuery struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}
