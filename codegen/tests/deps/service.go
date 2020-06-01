// Code generated by sysl DO NOT EDIT.
package deps

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Deps
type Service interface {
	GetApiDocsList(ctx context.Context, req *GetApiDocsListRequest) (*ApiDoc, error)
	GetSuccessList(ctx context.Context, req *GetSuccessListRequest) (*http.Header, error)
}

// Client for Deps API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for Deps
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetApiDocsList ...
func (s *Client) GetApiDocsList(ctx context.Context, req *GetApiDocsListRequest) (*ApiDoc, error) {
	required := []string{}
	var okResponse ApiDoc

	var errorResponse Status

	u, err := url.Parse(fmt.Sprintf("%s/api-docs", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, &errorResponse)
	if err != nil {
		response, ok := err.(*restlib.HTTPResult)
		if !ok {
			return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Deps <- GET "+u.String(), err)
		}

		return nil, common.CreateDownstreamError(ctx, common.DownstreamResponseError, response.HTTPResponse, response.Body, &errorResponse)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}

	OkApiDocResponse, ok := result.Response.(*ApiDoc)
	if ok {
		valErr := validator.Validate(OkApiDocResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkApiDocResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetSuccessList ...
func (s *Client) GetSuccessList(ctx context.Context, req *GetSuccessListRequest) (*http.Header, error) {
	required := []string{}
	u, err := url.Parse(fmt.Sprintf("%s/success", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, nil, nil)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Deps <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}

	return &result.HTTPResponse.Header, nil
}
