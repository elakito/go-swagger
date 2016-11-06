package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"

	"github.com/go-openapi/swag"
)

// FindTodosURL generates an URL for the find todos operation
type FindTodosURL struct {
	Limit *int32
	Since *int64

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *FindTodosURL) Build() (*url.URL, error) {
	var result url.URL

	var path = "/"

	result.Path = path

	qs := make(url.Values)

	var limit string
	if o.Limit != nil {
		limit = swag.FormatInt32(*o.Limit)
	}
	if limit != "" {
		qs.Set("limit", limit)
	}

	var since string
	if o.Since != nil {
		since = swag.FormatInt64(*o.Since)
	}
	if since != "" {
		qs.Set("since", since)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindTodosURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindTodosURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindTodosURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindTodosURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindTodosURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindTodosURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}