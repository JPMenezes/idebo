// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geoserverStyle client
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package geoserverstyle

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "geoserverStyle" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "geoserverStyle" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "list" endpoint of the "geoserverStyle" service.
// List may return the following errors:
//	- "not_found" (type *NotFound): Geoserver not found
//	- error: internal error
func (c *Client) List(ctx context.Context, p *ListPayload) (res StyleResultCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(StyleResultCollection), nil
}
