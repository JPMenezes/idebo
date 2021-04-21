// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geodata HTTP client encoders and decoders
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	geodata "jpmenezes.com/idebo/gen/geodata"
	geodataviews "jpmenezes.com/idebo/gen/geodata/views"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "geodata" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListGeodataPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("geodata", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the geodata list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*geodata.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("geodata", "list", "*geodata.ListPayload", v)
		}
		if p.Authentication != nil {
			head := *p.Authentication
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the geodata
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
				return nil, goahttp.ErrDecodingError("geodata", "list", err)
			}
			p := NewListGeodataResultCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := geodataviews.GeodataResultCollection{Projected: p, View: view}
			if err = geodataviews.ValidateGeodataResultCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("geodata", "list", err)
			}
			res := geodata.NewGeodataResultCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("geodata", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildUploadRequest instantiates a HTTP request object with method and path
// set to call the "geodata" service "upload" endpoint
func (c *Client) BuildUploadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UploadGeodataPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("geodata", "upload", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUploadRequest returns an encoder for requests sent to the geodata
// upload server.
func EncodeUploadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*geodata.FilesUpload)
		if !ok {
			return goahttp.ErrInvalidType("geodata", "upload", "*geodata.FilesUpload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("geodata", "upload", err)
		}
		return nil
	}
}

// NewGeodataUploadEncoder returns an encoder to encode the multipart request
// for the "geodata" service "upload" endpoint.
func NewGeodataUploadEncoder(encoderFn GeodataUploadEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*geodata.FilesUpload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeUploadResponse returns a decoder for responses returned by the geodata
// upload endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeUploadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("geodata", "upload", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("geodata", "upload", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "geodata" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*geodata.RemovePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("geodata", "remove", "*geodata.RemovePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveGeodataPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("geodata", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveResponse returns a decoder for responses returned by the geodata
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
			return nil, goahttp.ErrInvalidResponse("geodata", "remove", resp.StatusCode, string(body))
		}
	}
}

// unmarshalGeodataResultResponseToGeodataviewsGeodataResultView builds a value
// of type *geodataviews.GeodataResultView from a value of type
// *GeodataResultResponse.
func unmarshalGeodataResultResponseToGeodataviewsGeodataResultView(v *GeodataResultResponse) *geodataviews.GeodataResultView {
	res := &geodataviews.GeodataResultView{
		ID:         v.ID,
		Name:       v.Name,
		Entity:     v.Entity,
		Entityname: v.Entityname,
	}

	return res
}

// marshalGeodataFileUploadToFileUploadRequestBody builds a value of type
// *FileUploadRequestBody from a value of type *geodata.FileUpload.
func marshalGeodataFileUploadToFileUploadRequestBody(v *geodata.FileUpload) *FileUploadRequestBody {
	if v == nil {
		return nil
	}
	res := &FileUploadRequestBody{
		Type:  v.Type,
		Bytes: v.Bytes,
		Name:  v.Name,
	}

	return res
}

// marshalFileUploadRequestBodyToGeodataFileUpload builds a value of type
// *geodata.FileUpload from a value of type *FileUploadRequestBody.
func marshalFileUploadRequestBodyToGeodataFileUpload(v *FileUploadRequestBody) *geodata.FileUpload {
	if v == nil {
		return nil
	}
	res := &geodata.FileUpload{
		Type:  v.Type,
		Bytes: v.Bytes,
		Name:  v.Name,
	}

	return res
}
