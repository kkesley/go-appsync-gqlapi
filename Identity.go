package gqlapi

//Identity provides custom header for request
type Identity struct {
	SourceIP              []string `json:"sourceIp"`
	UserARN               string   `json:"userArn"`
	AccountID             string   `json:"accountId"`
	User                  string   `json:"user"`
	CognitoIdentityPoolID string   `json:"cognitoIdentityPoolId"`
	CognitoIdentityID     string   `json:"cognitoIdentityId"`
}
