package content

import (
	"fmt"
	"net/http"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Task struct {
	item.Item

	Name        string `json:"name"`
	Delegate    string `json:"delegate"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

// IndexContent satisfies the search.Searchable interface and returning true
// tells Ponzu that Task content should be indexed
func (t *Task) IndexContent() bool { return true }

// Update satisfies the api.Updateable interface, enabling properly formatted
// requests to update Task data
func (t *Task) Update(res http.ResponseWriter, req *http.Request) error {
	return nil
}

// MarshalEditor writes a buffer of html to edit a Task within the CMS
// and implements editor.Editable
func (t *Task) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(t,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Task field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", t, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: reference.Select("Delegate", t, map[string]string{
				"label": "Delegate",
			},
				"Employee",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Textarea("Description", t, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Checkbox("Complete", t, map[string]string{
				"label": "Complete",
			}, map[string]string{
				// "value": "Display Name",
				"true": "Completed",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Task editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Task"] = func() interface{} { return new(Task) }
}

// String defines how a Task is printed. Update it using more descriptive
// fields from the Task struct type
func (t *Task) String() string {
	return fmt.Sprintf("%s (%d)", t.Name, t.Timestamp)
}
