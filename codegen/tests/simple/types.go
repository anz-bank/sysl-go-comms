// Code generated by sysl DO NOT EDIT.
package simple

import (
	"time"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/convert"
	"github.com/anz-bank/sysl-go/validator"
	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// Item ...
type Item struct {
	A1   string `json:"A1"`
	A2   string `json:"A2"`
	Name string `json:"-"`
}

// Response ...
type Response struct {
	Data ItemSet `json:"Data"`
}

// Status ...
type Status struct {
	StatusField string `json:"statusField"`
}

// Stuff just some stuff
type Stuff struct {
	EmptyStuff     Empty                  `json:"emptyStuff"`
	InnerStuff     string                 `json:"innerStuff"`
	RawTimeStuff   time.Time              `json:"rawTimeStuff"`
	ResponseStuff  Response               `json:"responseStuff"`
	SensitiveStuff common.SensitiveString `json:"sensitiveStuff"`
	SequenceStuff  []Str                  `json:"sequenceStuff,omitempty"`
	TimeStuff      convert.JSONTime       `json:"timeStuff"`
}

// Generate wrapper set type
type ItemSet struct {
	M map[string]Item
}

// GetApiDocsListRequest ...
type GetApiDocsListRequest struct {
}

// GetGetSomeBytesListRequest ...
type GetGetSomeBytesListRequest struct {
}

// GetJustOkAndJustErrorListRequest ...
type GetJustOkAndJustErrorListRequest struct {
}

// GetJustReturnErrorListRequest ...
type GetJustReturnErrorListRequest struct {
}

// GetJustReturnOkListRequest ...
type GetJustReturnOkListRequest struct {
}

// GetOkTypeAndJustErrorListRequest ...
type GetOkTypeAndJustErrorListRequest struct {
}

// GetOopsListRequest ...
type GetOopsListRequest struct {
}

// GetRawListRequest ...
type GetRawListRequest struct {
}

// GetRawIntListRequest ...
type GetRawIntListRequest struct {
}

// GetSimpleAPIDocsListRequest ...
type GetSimpleAPIDocsListRequest struct {
}

// GetStuffListRequest ...
type GetStuffListRequest struct {
	Dt *convert.JSONTime
	St *string
	Bt *bool
	It *int64
}

// PostStuffRequest ...
type PostStuffRequest struct {
	Request Str
}

// *Item validator
func (s *Item) Validate() error {
	return validator.Validate(s)
}

// *Response validator
func (s *Response) Validate() error {
	return validator.Validate(s)
}

// *Status validator
func (s *Status) Validate() error {
	return validator.Validate(s)
}

// *Stuff validator
func (s *Stuff) Validate() error {
	return validator.Validate(s)
}

// *Item add
func (s *ItemSet) Add(item Item) {
	s.M[item.Name] = item
}

// *Item lookup
func (s *ItemSet) Lookup(Name string) Item {
	return s.M[Name]
}

// Pdf ...
type Pdf []byte

// Integer ...
type Integer int64

// Str ...
type Str string

// Empty ...
type Empty struct {
}
