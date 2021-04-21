// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geodata HTTP client CLI support package
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package client

import (
	"encoding/json"
	"fmt"

	geodata "jpmenezes.com/idebo/gen/geodata"
)

// BuildListPayload builds the payload for the geodata list endpoint from CLI
// flags.
func BuildListPayload(geodataListAuthentication string) (*geodata.ListPayload, error) {
	var authentication *string
	{
		if geodataListAuthentication != "" {
			authentication = &geodataListAuthentication
		}
	}
	v := &geodata.ListPayload{}
	v.Authentication = authentication

	return v, nil
}

// BuildUploadPayload builds the payload for the geodata upload endpoint from
// CLI flags.
func BuildUploadPayload(geodataUploadBody string) (*geodata.FilesUpload, error) {
	var err error
	var body UploadRequestBody
	{
		err = json.Unmarshal([]byte(geodataUploadBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"Files\": [\n         {\n            \"bytes\": \"TGliZXJvIGFzcGVybmF0dXIgc3VzY2lwaXQgdXQgaXBzdW0gbm9uIGl1c3RvLg==\",\n            \"name\": \"Possimus quo eligendi.\",\n            \"type\": \"Velit porro.\"\n         },\n         {\n            \"bytes\": \"TGliZXJvIGFzcGVybmF0dXIgc3VzY2lwaXQgdXQgaXBzdW0gbm9uIGl1c3RvLg==\",\n            \"name\": \"Possimus quo eligendi.\",\n            \"type\": \"Velit porro.\"\n         },\n         {\n            \"bytes\": \"TGliZXJvIGFzcGVybmF0dXIgc3VzY2lwaXQgdXQgaXBzdW0gbm9uIGl1c3RvLg==\",\n            \"name\": \"Possimus quo eligendi.\",\n            \"type\": \"Velit porro.\"\n         },\n         {\n            \"bytes\": \"TGliZXJvIGFzcGVybmF0dXIgc3VzY2lwaXQgdXQgaXBzdW0gbm9uIGl1c3RvLg==\",\n            \"name\": \"Possimus quo eligendi.\",\n            \"type\": \"Velit porro.\"\n         }\n      ]\n   }'")
		}
	}
	v := &geodata.FilesUpload{}
	if body.Files != nil {
		v.Files = make([]*geodata.FileUpload, len(body.Files))
		for i, val := range body.Files {
			v.Files[i] = marshalFileUploadRequestBodyToGeodataFileUpload(val)
		}
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the geodata remove endpoint from
// CLI flags.
func BuildRemovePayload(geodataRemoveID string) (*geodata.RemovePayload, error) {
	var id string
	{
		id = geodataRemoveID
	}
	v := &geodata.RemovePayload{}
	v.ID = id

	return v, nil
}
