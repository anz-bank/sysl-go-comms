// Code generated by sysl DO NOT EDIT.
package simple

import (
	"context"
	"net/http"
	"time"

	"github.com/anz-bank/sysl-go/codegen/tests/deps"
	"github.com/anz-bank/sysl-go/codegen/tests/downstream"
	"github.com/anz-bank/sysl-go/config"
)

// DefaultSimpleImpl ...
type DefaultSimpleImpl struct {
}

// NewDefaultSimpleImpl for Simple
func NewDefaultSimpleImpl() *DefaultSimpleImpl {
	return &DefaultSimpleImpl{}
}

// GetApiDocsList Client
type GetApiDocsListClient struct {
	GetApiDocsList     func(ctx context.Context, req *deps.GetApiDocsListRequest) (*deps.ApiDoc, error)
	GetServiceDocsList func(ctx context.Context, req *downstream.GetServiceDocsListRequest) (*downstream.ServiceDoc, error)
}

// GetGetSomeBytesList Client
type GetGetSomeBytesListClient struct {
}

// GetJustOkAndJustErrorList Client
type GetJustOkAndJustErrorListClient struct {
}

// GetJustReturnErrorList Client
type GetJustReturnErrorListClient struct {
}

// GetJustReturnOkList Client
type GetJustReturnOkListClient struct {
}

// GetOkTypeAndJustErrorList Client
type GetOkTypeAndJustErrorListClient struct {
}

// GetOopsList Client
type GetOopsListClient struct {
}

// GetRawList Client
type GetRawListClient struct {
}

// GetRawIntList Client
type GetRawIntListClient struct {
}

// GetSimpleAPIDocsList Client
type GetSimpleAPIDocsListClient struct {
	GetApiDocsList func(ctx context.Context, req *deps.GetApiDocsListRequest) (*deps.ApiDoc, error)
	GetSuccessList func(ctx context.Context, req *deps.GetSuccessListRequest) (*http.Header, error)
}

// GetStuffList Client
type GetStuffListClient struct {
}

// PostStuff Client
type PostStuffClient struct {
}

// ServiceInterface for Simple
type ServiceInterface struct {
	GetApiDocsList            func(ctx context.Context, req *GetApiDocsListRequest, client GetApiDocsListClient) (*[]deps.ApiDoc, error)
	GetGetSomeBytesList       func(ctx context.Context, req *GetGetSomeBytesListRequest, client GetGetSomeBytesListClient) (*Pdf, error)
	GetJustOkAndJustErrorList func(ctx context.Context, req *GetJustOkAndJustErrorListRequest, client GetJustOkAndJustErrorListClient) error
	GetJustReturnErrorList    func(ctx context.Context, req *GetJustReturnErrorListRequest, client GetJustReturnErrorListClient) error
	GetJustReturnOkList       func(ctx context.Context, req *GetJustReturnOkListRequest, client GetJustReturnOkListClient) error
	GetOkTypeAndJustErrorList func(ctx context.Context, req *GetOkTypeAndJustErrorListRequest, client GetOkTypeAndJustErrorListClient) (*Response, error)
	GetOopsList               func(ctx context.Context, req *GetOopsListRequest, client GetOopsListClient) (*Response, error)
	GetRawList                func(ctx context.Context, req *GetRawListRequest, client GetRawListClient) (*Str, error)
	GetRawIntList             func(ctx context.Context, req *GetRawIntListRequest, client GetRawIntListClient) (*Integer, error)
	GetSimpleAPIDocsList      func(ctx context.Context, req *GetSimpleAPIDocsListRequest, client GetSimpleAPIDocsListClient) (*deps.ApiDoc, error)
	GetStuffList              func(ctx context.Context, req *GetStuffListRequest, client GetStuffListClient) (*Stuff, error)
	PostStuff                 func(ctx context.Context, req *PostStuffRequest, client PostStuffClient) (*Str, error)
}

// DownstreamConfig for Simple
type DownstreamConfig struct {
	ContextTimeout time.Duration               `yaml:"contextTimeout"`
	Deps           config.CommonDownstreamData `yaml:"deps"`
	Downstream     config.CommonDownstreamData `yaml:"downstream"`
}
