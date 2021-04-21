// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geoserverStyle HTTP server encoders and decoders
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	geoserverstyle "jpmenezes.com/idebo/gen/geoserver_style"
	geoserverstyleviews "jpmenezes.com/idebo/gen/geoserver_style/views"
)

// EncodeListResponse returns an encoder for responses returned by the
// geoserverStyle list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(geoserverstyleviews.StyleResultCollection)
		enc := encoder(ctx, w)
		body := NewStyleResultResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the geoserverStyle
// list endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			geoserverURL string
			entityid     string
			view         *string
			err          error
		)
		geoserverURL = r.URL.Query().Get("geoserverURL")
		if geoserverURL == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("geoserverURL", "query string"))
		}
		entityid = r.URL.Query().Get("entityid")
		if entityid == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("entityid", "query string"))
		}
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(geoserverURL, entityid, view)

		return payload, nil
	}
}

// EncodeListError returns an encoder for errors returned by the list
// geoserverStyle endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*geoserverstyle.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalGeoserverstyleviewsStyleResultViewToStyleResultResponseTiny builds a
// value of type *StyleResultResponseTiny from a value of type
// *geoserverstyleviews.StyleResultView.
func marshalGeoserverstyleviewsStyleResultViewToStyleResultResponseTiny(v *geoserverstyleviews.StyleResultView) *StyleResultResponseTiny {
	res := &StyleResultResponseTiny{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}