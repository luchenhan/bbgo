// Code generated by "requestgen -method POST -url /api/v3/wallet/m/loans/:currency -type MarginLoanRequest -responseType .LoanRecord"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (m *MarginLoanRequest) Amount(amount string) *MarginLoanRequest {
	m.amount = amount
	return m
}

func (m *MarginLoanRequest) Currency(currency string) *MarginLoanRequest {
	m.currency = currency
	return m
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (m *MarginLoanRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (m *MarginLoanRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check amount field -> json key amount
	amount := m.amount

	// assign parameter of amount
	params["amount"] = amount

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (m *MarginLoanRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := m.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if m.isVarSlice(_v) {
			m.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (m *MarginLoanRequest) GetParametersJSON() ([]byte, error) {
	params, err := m.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (m *MarginLoanRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check currency field -> json key currency
	currency := m.currency

	// TEMPLATE check-required
	if len(currency) == 0 {
		return nil, fmt.Errorf("currency is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of currency
	params["currency"] = currency

	return params, nil
}

func (m *MarginLoanRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (m *MarginLoanRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (m *MarginLoanRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (m *MarginLoanRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := m.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (m *MarginLoanRequest) Do(ctx context.Context) (*LoanRecord, error) {

	params, err := m.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/api/v3/wallet/m/loans/:currency"
	slugs, err := m.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = m.applySlugsToUrl(apiURL, slugs)

	req, err := m.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := m.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse LoanRecord
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
