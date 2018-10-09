# go-appsync-gqlapi
GraphQL API Wrapper for AWS Appsync

This repository is experimental and used by internal team only.

## Generating clients
Client constructor requires 2 arguments:
1. AppSync URL
2. Region

This is how you use it

`client := gqlapi.New(os.Getenv("APPSYNC_URL"), os.Getenv("CHATTING_REGION"))`

## Mutation
Mutation sends a post request to AppSync. It requires 2 arguments:
1. GraphQL query string
1. Variables for the query string

example:

```golang
query := `
    mutation ($message: String!) {
        someMutation(
            message: $message
        ){
            message
        }
    }
`
variables := Response{
    Message:    "Hello World!",
}

_, err := client.Post(query, variables)
if err != nil {
    return Response{}, err
}
```