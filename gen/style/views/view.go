// Code generated by goa v3.1.1, DO NOT EDIT.
//
// style views
//
// Command:
// $ goa gen jpmenezes.com/idebo/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StyleResultCollection is the viewed result type that is projected based on a
// view.
type StyleResultCollection struct {
	// Type to project
	Projected StyleResultCollectionView
	// View to render
	View string
}

// StyleResult is the viewed result type that is projected based on a view.
type StyleResult struct {
	// Type to project
	Projected *StyleResultView
	// View to render
	View string
}

// StyleResultCollectionView is a type that runs validations on a projected
// type.
type StyleResultCollectionView []*StyleResultView

// StyleResultView is a type that runs validations on a projected type.
type StyleResultView struct {
	// ID is the unique id of the style.
	ID *string
	// Name of style
	Name *string
}

var (
	// StyleResultCollectionMap is a map of attribute names in result type
	// StyleResultCollection indexed by view name.
	StyleResultCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
		},
		"tiny": []string{
			"id",
			"name",
		},
	}
	// StyleResultMap is a map of attribute names in result type StyleResult
	// indexed by view name.
	StyleResultMap = map[string][]string{
		"default": []string{
			"id",
			"name",
		},
		"tiny": []string{
			"id",
			"name",
		},
	}
)

// ValidateStyleResultCollection runs the validations defined on the viewed
// result type StyleResultCollection.
func ValidateStyleResultCollection(result StyleResultCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStyleResultCollectionView(result.Projected)
	case "tiny":
		err = ValidateStyleResultCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStyleResult runs the validations defined on the viewed result type
// StyleResult.
func ValidateStyleResult(result *StyleResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStyleResultView(result.Projected)
	case "tiny":
		err = ValidateStyleResultViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStyleResultCollectionView runs the validations defined on
// StyleResultCollectionView using the "default" view.
func ValidateStyleResultCollectionView(result StyleResultCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStyleResultView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStyleResultCollectionViewTiny runs the validations defined on
// StyleResultCollectionView using the "tiny" view.
func ValidateStyleResultCollectionViewTiny(result StyleResultCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStyleResultViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStyleResultView runs the validations defined on StyleResultView
// using the "default" view.
func ValidateStyleResultView(result *StyleResultView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 255, false))
		}
	}
	return
}

// ValidateStyleResultViewTiny runs the validations defined on StyleResultView
// using the "tiny" view.
func ValidateStyleResultViewTiny(result *StyleResultView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 255, false))
		}
	}
	return
}