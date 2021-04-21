// Code generated by goa v3.1.1, DO NOT EDIT.
//
// entity HTTP server types
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package server

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	entity "jpmenezes.com/idebo/gen/entity"
	entityviews "jpmenezes.com/idebo/gen/entity/views"
)

// AddRequestBody is the type of the "entity" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Entity to add
	Entity *EntityRequestBody `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
}

// UpdateRequestBody is the type of the "entity" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Entity to update
	Entity *EntityRequestBody `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
}

// EntityResultResponseCollection is the type of the "entity" service "list"
// endpoint HTTP response body.
type EntityResultResponseCollection []*EntityResultResponse

// ShowResponseBody is the type of the "entity" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the entity.
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of entity
	Name string `form:"name" json:"name" xml:"name"`
	// Folder name of entity. Serves as entity identifier.
	Folder string `form:"folder" json:"folder" xml:"folder"`
	// Inactive entity (in maintenance)
	Inactive *bool `form:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
}

// ShowbyfieldResponseBody is the type of the "entity" service "showbyfield"
// endpoint HTTP response body.
type ShowbyfieldResponseBody struct {
	// ID is the unique id of the entity.
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of entity
	Name string `form:"name" json:"name" xml:"name"`
	// Folder name of entity. Serves as entity identifier.
	Folder string `form:"folder" json:"folder" xml:"folder"`
	// Inactive entity (in maintenance)
	Inactive *bool `form:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "entity" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing viewer
	ID string `form:"id" json:"id" xml:"id"`
}

// ShowbyfieldNotFoundResponseBody is the type of the "entity" service
// "showbyfield" endpoint HTTP response body for the "not_found" error.
type ShowbyfieldNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing viewer
	ID string `form:"id" json:"id" xml:"id"`
}

// EntityResultResponse is used to define fields on response body types.
type EntityResultResponse struct {
	// ID is the unique id of the entity.
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of entity
	Name string `form:"name" json:"name" xml:"name"`
	// Folder name of entity. Serves as entity identifier.
	Folder string `form:"folder" json:"folder" xml:"folder"`
	// Inactive entity (in maintenance)
	Inactive *bool `form:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
}

// EntityRequestBody is used to define fields on request body types.
type EntityRequestBody struct {
	// Name of entity
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Folder name of entity. Serves as entity identifier.
	Folder *string `form:"folder,omitempty" json:"folder,omitempty" xml:"folder,omitempty"`
	// Inactive entity (in maintenance)
	Inactive *bool `form:"inactive,omitempty" json:"inactive,omitempty" xml:"inactive,omitempty"`
}

// NewEntityResultResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "entity" service.
func NewEntityResultResponseCollection(res entityviews.EntityResultCollectionView) EntityResultResponseCollection {
	body := make([]*EntityResultResponse, len(res))
	for i, val := range res {
		body[i] = marshalEntityviewsEntityResultViewToEntityResultResponse(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "entity" service.
func NewShowResponseBody(res *entityviews.EntityResultView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		Folder:   *res.Folder,
		Inactive: res.Inactive,
	}
	return body
}

// NewShowbyfieldResponseBody builds the HTTP response body from the result of
// the "showbyfield" endpoint of the "entity" service.
func NewShowbyfieldResponseBody(res *entityviews.EntityResultView) *ShowbyfieldResponseBody {
	body := &ShowbyfieldResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		Folder:   *res.Folder,
		Inactive: res.Inactive,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "entity" service.
func NewShowNotFoundResponseBody(res *entity.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowbyfieldNotFoundResponseBody builds the HTTP response body from the
// result of the "showbyfield" endpoint of the "entity" service.
func NewShowbyfieldNotFoundResponseBody(res *entity.NotFound) *ShowbyfieldNotFoundResponseBody {
	body := &ShowbyfieldNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewListPayload builds a entity service list endpoint payload.
func NewListPayload(authentication *string) *entity.ListPayload {
	v := &entity.ListPayload{}
	v.Authentication = authentication

	return v
}

// NewShowPayload builds a entity service show endpoint payload.
func NewShowPayload(id string, view *string) *entity.ShowPayload {
	v := &entity.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewShowbyfieldPayload builds a entity service showbyfield endpoint payload.
func NewShowbyfieldPayload(fieldname string, fieldvalue string, view *string) *entity.ShowbyfieldPayload {
	v := &entity.ShowbyfieldPayload{}
	v.Fieldname = fieldname
	v.Fieldvalue = fieldvalue
	v.View = view

	return v
}

// NewAddPayload builds a entity service add endpoint payload.
func NewAddPayload(body *AddRequestBody, authentication *string) *entity.AddPayload {
	v := &entity.AddPayload{}
	if body.Entity != nil {
		v.Entity = unmarshalEntityRequestBodyToEntityEntity(body.Entity)
	}
	v.Authentication = authentication

	return v
}

// NewUpdatePayload builds a entity service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *entity.UpdatePayload {
	v := &entity.UpdatePayload{}
	if body.Entity != nil {
		v.Entity = unmarshalEntityRequestBodyToEntityEntity(body.Entity)
	}
	v.ID = id

	return v
}

// NewRemovePayload builds a entity service remove endpoint payload.
func NewRemovePayload(id string) *entity.RemovePayload {
	v := &entity.RemovePayload{}
	v.ID = id

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Entity != nil {
		if err2 := ValidateEntityRequestBody(body.Entity); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Entity != nil {
		if err2 := ValidateEntityRequestBody(body.Entity); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEntityRequestBody runs the validations defined on EntityRequestBody
func ValidateEntityRequestBody(body *EntityRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Folder == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("folder", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 255, false))
		}
	}
	if body.Folder != nil {
		if utf8.RuneCountInString(*body.Folder) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.folder", *body.Folder, utf8.RuneCountInString(*body.Folder), 255, false))
		}
	}
	return
}
