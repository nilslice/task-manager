package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Employee struct {
	item.Item

	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo"`
}

func (e *Employee) IndexContent() bool { return true }

// MarshalEditor writes a buffer of html to edit a Employee within the CMS
// and implements editor.Editable
func (e *Employee) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Employee field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", e, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", e, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", e, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload the Photo here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Employee editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Employee"] = func() interface{} { return new(Employee) }
}

// String defines how a Employee is printed. Update it using more descriptive
// fields from the Employee struct type
func (e *Employee) String() string {
	return e.Email
}
