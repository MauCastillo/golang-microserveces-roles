package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func TestHandleRequest(t *testing.T) {
	c := require.New(t)

	event := events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{
		"date": "2019-12-01", "day": "5",
	}}

	_, err := HandleRequest(context.Background(), event)
	c.NoError(err)
}

func TestHandleRequestError(t *testing.T) {
	c := require.New(t)

	event := events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{}}

	res, err := HandleRequest(context.Background(), event)
	c.NoError(err)

	c.Equal(res.Body, "{error: empty date parameter}")

	event = events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{"date": "2019-12-01"}}
	res, err = HandleRequest(context.Background(), event)
	c.NoError(err)

	c.Equal(res.Body, "{error: obtaining the range of days}")
}
