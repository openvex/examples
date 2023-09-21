# My First OpenVEX Program

Using the OpenVEX library in go is relatively easy. You work with documents,
add a few statements and render it into json. Here is a a short example:

```golang
package main

import (
	"os"

	"github.com/openvex/go-vex/pkg/vex"
)

func main() {
	// Create a new document
	doc := vex.VEX{
		Metadata: vex.Metadata{
			ID:         "https://openvex.dev/docs/public/my-first-vex",
			Context:    vex.Context,
			Author:     "John Doe",
			AuthorRole: "VEXing Engineer",
			Version:    1.0,
			Tooling:    "hello/vex",
			Supplier:   "You!",
		},
	}

	/// Add a statement to the document

	doc.Statements = append(doc.Statements, vex.Statement{
		Vulnerability: vex.Vulnerability{
			ID:          "CVE-2014-123456",
			Description: "It's really bad",
		},
		Products: []vex.Product{
			{
				Component:     vex.Component{ID: "pkg:generic/1.0.0"},
				Subcomponents: []vex.Subcomponent{},
			},
		},
		Status:      "fixed",
		StatusNotes: "It works now",
	})

	doc.ToJSON(os.Stdout)
}
```

The example code above will produce the following document:

```json
{
  "@context": "https://openvex.dev/ns",
  "@id": "https://openvex.dev/docs/public/my-first-vex",
  "author": "John Doe",
  "role": "VEXing Engineer",
  "timestamp": null,
  "version": 1,
  "tooling": "hello/vex",
  "supplier": "You!",
  "statements": [
    {
      "vulnerability": {
        "@id": "CVE-2014-123456",
        "description": "It's really bad"
      },
      "products": [
        {
          "@id": "pkg:generic/1.0.0"
        }
      ],
      "status": "fixed",
      "status_notes": "It works now"
    }
  ]
}

```

If you want to play around with this example, see the files in this directory.
