// Code generated by "requestgen -method GET -url v2/orders -type GetOrdersRequest -responseType []Order"; DO NOT EDIT.

package max

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (g *GetOrdersRequest) Market(market string) *GetOrdersRequest {
	g.market = market
	return g
}

func (g *GetOrdersRequest) Side(side string) *GetOrdersRequest {
	g.side = &side
	return g
}

func (g *GetOrdersRequest) GroupID(groupID uint32) *GetOrdersRequest {
	g.groupID = &groupID
	return g
}

func (g *GetOrdersRequest) Offset(offset int) *GetOrdersRequest {
	g.offset = &offset
	return g
}

func (g *GetOrdersRequest) Limit(limit int) *GetOrdersRequest {
	g.limit = &limit
	return g
}

func (g *GetOrdersRequest) Page(page int) *GetOrdersRequest {
	g.page = &page
	return g
}

func (g *GetOrdersRequest) OrderBy(orderBy string) *GetOrdersRequest {
	g.orderBy = &orderBy
	return g
}

func (g *GetOrdersRequest) State(state []OrderState) *GetOrdersRequest {
	g.state = state
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetOrdersRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetOrdersRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check market field -> json key market
	market := g.market

	// assign parameter of market
	params["market"] = market
	// check side field -> json key side
	if g.side != nil {
		side := *g.side

		// assign parameter of side
		params["side"] = side
	} else {
	}
	// check groupID field -> json key groupID
	if g.groupID != nil {
		groupID := *g.groupID

		// assign parameter of groupID
		params["groupID"] = groupID
	} else {
	}
	// check offset field -> json key offset
	if g.offset != nil {
		offset := *g.offset

		// assign parameter of offset
		params["offset"] = offset
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}
	// check page field -> json key page
	if g.page != nil {
		page := *g.page

		// assign parameter of page
		params["page"] = page
	} else {
	}
	// check orderBy field -> json key order_by
	if g.orderBy != nil {
		orderBy := *g.orderBy

		// assign parameter of orderBy
		params["order_by"] = orderBy
	} else {
		orderBy := "desc"

		// assign parameter of orderBy
		params["order_by"] = orderBy
	}
	// check state field -> json key state
	state := g.state

	// assign parameter of state
	params["state"] = state

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetOrdersRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetOrdersRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetOrdersRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetOrdersRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetOrdersRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetOrdersRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetOrdersRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetOrdersRequest) Do(ctx context.Context) ([]Order, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "v2/orders"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []Order
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
