package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	layoutISO       = "2006-01-02"
	fieldDate       = "date"
	fieldDay        = "day"
	baseURL         = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/%s"
	statisticsFails = "formatting the statistics"
	emptyDate       = "empty date parameter"
	layoutError     = "{error: %s}"
)

var (
	errConvertDay = errors.New("obtaining the range of days")
)

type Sales []struct {
	ClientID int       `json:"clientId"`
	Nombre   string    `json:"nombre"`
	Compro   bool      `json:"compro"`
	Date     time.Time `json:"date"`
	TDC      string    `json:"tdc"`
	Monto    float32   `json:"monto"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if _, ok := request.QueryStringParameters[fieldDate]; !ok {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(layoutError, emptyDate), StatusCode: 400}, nil
	}

	response, err := rangeQuery(request)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(layoutError, err.Error()), StatusCode: 500}, nil
	}

	statistics, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(layoutError, statisticsFails), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(statistics), StatusCode: 200}, nil
}

func rangeQuery(request events.APIGatewayProxyRequest) (*Response, error) {
	date := request.QueryStringParameters[fieldDate]
	days := request.QueryStringParameters[fieldDay]
	res := NewResponse()

	startDate, err := time.Parse(layoutISO, date)
	if err != nil {
		return res, err
	}

	rangeDays, err := strconv.Atoi(days)
	if err != nil {
		return res, errConvertDay
	}

	for i := 0; i < rangeDays; i++ {
		sales, err := requestGet(startDate.Format(layoutISO))
		if err != nil {
			continue
		}

		res.analizer(sales)

		startDate = startDate.AddDate(0, 0, 1)
	}

	return res, nil
}

func requestGet(dateQuery string) (Sales, error) {
	apiRecruit := fmt.Sprintf(baseURL, dateQuery)

	resp, err := http.Get(apiRecruit)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sales := Sales{}
	err = json.Unmarshal(body, &sales)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func main() {
	lambda.Start(HandleRequest)
}
