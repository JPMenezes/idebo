// Code generated by goa v3.1.1, DO NOT EDIT.
//
// entity client
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package entity

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "entity" service client.
type Client struct {
	ListEndpoint        goa.Endpoint
	ShowEndpoint        goa.Endpoint
	ShowbyfieldEndpoint goa.Endpoint
	AddEndpoint         goa.Endpoint
	UpdateEndpoint      goa.Endpoint
	RemoveEndpoint      goa.Endpoint
}

// NewClient initializes a "entity" service client given the endpoints.
func NewClient(list, show, showbyfield, add, update, remove goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:        list,
		ShowEndpoint:        show,
		ShowbyfieldEndpoint: showbyfield,
		AddEndpoint:         add,
		UpdateEndpoint:      update,
		RemoveEndpoint:      remove,
	}
}

// List calls the "list" endpoint of the "entity" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res EntityResultCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(EntityResultCollection), nil
}

// Show calls the "show" endpoint of the "entity" service.
// Show may return the following errors:
//	- "not_found" (type *NotFound): Entity not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *EntityResult, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EntityResult), nil
}

// Showbyfield calls the "showbyfield" endpoint of the "entity" service.
// Showbyfield may return the following errors:
//	- "not_found" (type *NotFound): Entity not found
//	- error: internal error
func (c *Client) Showbyfield(ctx context.Context, p *ShowbyfieldPayload) (res *EntityResult, err error) {
	var ires interface{}
	ires, err = c.ShowbyfieldEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EntityResult), nil
}

// Add calls the "add" endpoint of the "entity" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Update calls the "update" endpoint of the "entity" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Remove calls the "remove" endpoint of the "entity" service.
// Remove may return the following errors:
//	- "not_found" (type *NotFound): Entity not found
//	- error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}
