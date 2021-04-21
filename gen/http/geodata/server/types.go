// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geodata HTTP server types
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package server

import (
	geodata "jpmenezes.com/idebo/gen/geodata"
	geodataviews "jpmenezes.com/idebo/gen/geodata/views"
)

// UploadRequestBody is the type of the "geodata" service "upload" endpoint
// HTTP request body.
type UploadRequestBody struct {
	// Collection of uploaded files
	Files []*FileUploadRequestBody `form:"Files,omitempty" json:"Files,omitempty" xml:"Files,omitempty"`
}

// GeodataResultResponseCollection is the type of the "geodata" service "list"
// endpoint HTTP response body.
type GeodataResultResponseCollection []*GeodataResultResponse

// GeodataResultResponseTinyCollection is the type of the "geodata" service
// "list" endpoint HTTP response body.
type GeodataResultResponseTinyCollection []*GeodataResultResponseTiny

// GeodataResultResponse is used to define fields on response body types.
type GeodataResultResponse struct {
	// ID is the unique id of the geodata file.
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of geodata file
	Name string `form:"name" json:"name" xml:"name"`
	// Entity to which the geodata file belongs
	Entity     string  `form:"entity" json:"entity" xml:"entity"`
	Entityname *string `form:"entityname,omitempty" json:"entityname,omitempty" xml:"entityname,omitempty"`
}

// GeodataResultResponseTiny is used to define fields on response body types.
type GeodataResultResponseTiny struct {
	// ID is the unique id of the geodata file.
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of geodata file
	Name string `form:"name" json:"name" xml:"name"`
	// Entity to which the geodata file belongs
	Entity     string  `form:"entity" json:"entity" xml:"entity"`
	Entityname *string `form:"entityname,omitempty" json:"entityname,omitempty" xml:"entityname,omitempty"`
}

// FileUploadRequestBody is used to define fields on request body types.
type FileUploadRequestBody struct {
	Type  *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Bytes []byte  `form:"bytes,omitempty" json:"bytes,omitempty" xml:"bytes,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewGeodataResultResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "geodata" service.
func NewGeodataResultResponseCollection(res geodataviews.GeodataResultCollectionView) GeodataResultResponseCollection {
	body := make([]*GeodataResultResponse, len(res))
	for i, val := range res {
		body[i] = marshalGeodataviewsGeodataResultViewToGeodataResultResponse(val)
	}
	return body
}

// NewGeodataResultResponseTinyCollection builds the HTTP response body from
// the result of the "list" endpoint of the "geodata" service.
func NewGeodataResultResponseTinyCollection(res geodataviews.GeodataResultCollectionView) GeodataResultResponseTinyCollection {
	body := make([]*GeodataResultResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalGeodataviewsGeodataResultViewToGeodataResultResponseTiny(val)
	}
	return body
}

// NewListPayload builds a geodata service list endpoint payload.
func NewListPayload(authentication *string) *geodata.ListPayload {
	v := &geodata.ListPayload{}
	v.Authentication = authentication

	return v
}

// NewUploadFilesUpload builds a geodata service upload endpoint payload.
func NewUploadFilesUpload(body *UploadRequestBody) *geodata.FilesUpload {
	v := &geodata.FilesUpload{}
	if body.Files != nil {
		v.Files = make([]*geodata.FileUpload, len(body.Files))
		for i, val := range body.Files {
			v.Files[i] = unmarshalFileUploadRequestBodyToGeodataFileUpload(val)
		}
	}

	return v
}

// NewRemovePayload builds a geodata service remove endpoint payload.
func NewRemovePayload(id string) *geodata.RemovePayload {
	v := &geodata.RemovePayload{}
	v.ID = id

	return v
}
