// Code generated by goa v3.1.1, DO NOT EDIT.
//
// geoserverStyle HTTP client CLI support package
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package client

import (
	goa "goa.design/goa/v3/pkg"
	geoserverstyle "jpmenezes.com/idebo/gen/geoserver_style"
)

// BuildListPayload builds the payload for the geoserverStyle list endpoint
// from CLI flags.
func BuildListPayload(geoserverStyleListGeoserverURL string, geoserverStyleListEntityid string, geoserverStyleListView string) (*geoserverstyle.ListPayload, error) {
	var err error
	var geoserverURL string
	{
		geoserverURL = geoserverStyleListGeoserverURL
	}
	var entityid string
	{
		entityid = geoserverStyleListEntityid
	}
	var view *string
	{
		if geoserverStyleListView != "" {
			view = &geoserverStyleListView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &geoserverstyle.ListPayload{}
	v.GeoserverURL = geoserverURL
	v.Entityid = entityid
	v.View = view

	return v, nil
}
