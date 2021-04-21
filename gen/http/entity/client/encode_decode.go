// Code generated by goa v3.1.1, DO NOT EDIT.
//
// entity HTTP client encoders and decoders
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	entity "jpmenezes.com/idebo/gen/entity"
	entityviews "jpmenezes.com/idebo/gen/entity/views"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "entity" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListEntityPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the entity list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entity.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("entity", "list", "*entity.ListPayload", v)
		}
		if p.Authentication != nil {
			head := *p.Authentication
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the entity
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "list", err)
			}
			p := NewListEntityResultCollectionOK(body)
			view := "default"
			vres := entityviews.EntityResultCollection{Projected: p, View: view}
			if err = entityviews.ValidateEntityResultCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("entity", "list", err)
			}
			res := entity.NewEntityResultCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "entity" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*entity.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("entity", "show", "*entity.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowEntityPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the entity show
// server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entity.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("entity", "show", "*entity.ShowPayload", v)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the entity
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "not_found" (type *entity.NotFound): http.StatusNotFound
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "show", err)
			}
			p := NewShowEntityResultOK(&body)
			view := "default"
			vres := &entityviews.EntityResult{Projected: p, View: view}
			if err = entityviews.ValidateEntityResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("entity", "show", err)
			}
			res := entity.NewEntityResult(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "show", err)
			}
			err = ValidateShowNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("entity", "show", err)
			}
			return nil, NewShowNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildShowbyfieldRequest instantiates a HTTP request object with method and
// path set to call the "entity" service "showbyfield" endpoint
func (c *Client) BuildShowbyfieldRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		fieldname  string
		fieldvalue string
	)
	{
		p, ok := v.(*entity.ShowbyfieldPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("entity", "showbyfield", "*entity.ShowbyfieldPayload", v)
		}
		fieldname = p.Fieldname
		fieldvalue = p.Fieldvalue
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowbyfieldEntityPath(fieldname, fieldvalue)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "showbyfield", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowbyfieldRequest returns an encoder for requests sent to the entity
// showbyfield server.
func EncodeShowbyfieldRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entity.ShowbyfieldPayload)
		if !ok {
			return goahttp.ErrInvalidType("entity", "showbyfield", "*entity.ShowbyfieldPayload", v)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeShowbyfieldResponse returns a decoder for responses returned by the
// entity showbyfield endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeShowbyfieldResponse may return the following errors:
//	- "not_found" (type *entity.NotFound): http.StatusNotFound
//	- error: internal error
func DecodeShowbyfieldResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowbyfieldResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "showbyfield", err)
			}
			p := NewShowbyfieldEntityResultOK(&body)
			view := "default"
			vres := &entityviews.EntityResult{Projected: p, View: view}
			if err = entityviews.ValidateEntityResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("entity", "showbyfield", err)
			}
			res := entity.NewEntityResult(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowbyfieldNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "showbyfield", err)
			}
			err = ValidateShowbyfieldNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("entity", "showbyfield", err)
			}
			return nil, NewShowbyfieldNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "showbyfield", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "entity" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddEntityPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the entity add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entity.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("entity", "add", "*entity.AddPayload", v)
		}
		if p.Authentication != nil {
			head := *p.Authentication
			req.Header.Set("Authorization", head)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("entity", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the entity add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "entity" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*entity.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("entity", "update", "*entity.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateEntityPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the entity
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entity.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("entity", "update", "*entity.UpdatePayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("entity", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the entity
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "entity" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*entity.RemovePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("entity", "remove", "*entity.RemovePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveEntityPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveResponse returns a decoder for responses returned by the entity
// remove endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity", "remove", resp.StatusCode, string(body))
		}
	}
}

// unmarshalEntityResultResponseToEntityviewsEntityResultView builds a value of
// type *entityviews.EntityResultView from a value of type
// *EntityResultResponse.
func unmarshalEntityResultResponseToEntityviewsEntityResultView(v *EntityResultResponse) *entityviews.EntityResultView {
	res := &entityviews.EntityResultView{
		ID:       v.ID,
		Name:     v.Name,
		Folder:   v.Folder,
		Inactive: v.Inactive,
	}

	return res
}

// marshalEntityEntityToEntityRequestBody builds a value of type
// *EntityRequestBody from a value of type *entity.Entity.
func marshalEntityEntityToEntityRequestBody(v *entity.Entity) *EntityRequestBody {
	if v == nil {
		return nil
	}
	res := &EntityRequestBody{
		Name:     v.Name,
		Folder:   v.Folder,
		Inactive: v.Inactive,
	}

	return res
}

// marshalEntityRequestBodyToEntityEntity builds a value of type *entity.Entity
// from a value of type *EntityRequestBody.
func marshalEntityRequestBodyToEntityEntity(v *EntityRequestBody) *entity.Entity {
	if v == nil {
		return nil
	}
	res := &entity.Entity{
		Name:     v.Name,
		Folder:   v.Folder,
		Inactive: v.Inactive,
	}

	return res
}
