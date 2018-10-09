package gqlapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

//Client hold shared information
type Client struct {
	client     *http.Client
	appsyncURL string
	region     string
	signer     *v4.Signer
}

//New generate new client
func (c Client) New(url string, region string) {
	c.client = new(http.Client)
	c.appsyncURL = url
	c.region = region
	config := aws.Config{
		Region: aws.String("ap-southeast-1"),
	}
	sess := session.Must(session.NewSession(&config))
	c.signer = v4.NewSigner(sess.Config.Credentials)
}

//Post make a POST request to AppSync for mutation
func (c Client) Post(query string, variables interface{}) (*http.Response, error) {
	queryObj := StandardQuery{
		Query:     query,
		Variables: variables,
	}
	queryBytes, err := json.Marshal(&queryObj)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	queryReader := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", c.appsyncURL, queryReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.signer.Sign(req, queryReader, "appsync", c.region, time.Now())
	return c.client.Do(req)
}
