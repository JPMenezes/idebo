// Code generated by goa v3.1.1, DO NOT EDIT.
//
// style client
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package style

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "style" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "style" service client given the endpoints.
func NewClient(list, show, add, remove goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		AddEndpoint:    add,
		RemoveEndpoint: remove,
	}
}

// List calls the "list" endpoint of the "style" service.
func (c *Client) List(ctx context.Context) (res StyleResultCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StyleResultCollection), nil
}

// Show calls the "show" endpoint of the "style" service.
// Show may return the following errors:
//	- "not_found" (type *NotFound): Style not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StyleResult, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StyleResult), nil
}

// Add calls the "add" endpoint of the "style" service.
func (c *Client) Add(ctx context.Context, p *Style) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Remove calls the "remove" endpoint of the "style" service.
// Remove may return the following errors:
//	- "not_found" (type *NotFound): Style not found
//	- error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}
