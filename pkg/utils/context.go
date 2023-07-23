package utils

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// Read request query
func ReadQueryRequest(ctx echo.Context, request interface{}) error {
	queryMap, err := queryStringToMap(ctx.QueryString())
	if err != nil {
		return err
	}
	queryByte, _ := json.Marshal(queryMap)
	if err := json.Unmarshal(queryByte, request); err != nil {
		return err
	}
	return nil
}

// Read request body
func ReadBodyRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return nil
}

// Get context with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

func queryStringToMap(query string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	for _, q := range strings.Split(query, "&") {
		qs := strings.Split(q, "=")
		if len(qs) == 2 {
			num, err := strconv.Atoi(qs[1])
			if err == nil {
				res[qs[0]] = num
				continue
			}
			res[qs[0]] = qs[1]
		}
	}
	return res, nil
}
