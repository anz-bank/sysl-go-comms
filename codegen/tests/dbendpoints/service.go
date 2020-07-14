// Code generated by sysl DO NOT EDIT.
package dbendpoints

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for DbEndpoints
type Service interface {
	GetCompanyLocationList(ctx context.Context, req *GetCompanyLocationListRequest) (*GetCompanyLocationResponse, error)
}

// Client for DbEndpoints API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for DbEndpoints
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetCompanyLocationList ...
func (s *Client) GetCompanyLocationList(ctx context.Context, req *GetCompanyLocationListRequest) (*GetCompanyLocationResponse, error) {
	required := []string{}
	var okResponse GetCompanyLocationResponse
	u, err := url.Parse(fmt.Sprintf("%s/company/location", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	q.Add("deptLoc", req.DeptLoc)

	if req.CompanyName != nil {
		q.Add("companyName", *req.CompanyName)
	}

	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, nil)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: DbEndpoints <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkGetCompanyLocationResponseResponse, ok := result.Response.(*GetCompanyLocationResponse)
	if ok {
		valErr := validator.Validate(OkGetCompanyLocationResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkGetCompanyLocationResponseResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}
