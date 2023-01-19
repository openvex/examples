# My First OpenVEX Program

Using the OpenVEX library in go is relatively easy. You work with documents,
add a few statements and render it into json. Here is a a shoret example:

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
			Version:    "1.0",
			Tooling:    "hello/vex",
			Supplier:   "You!",
		},
	}

	/// Add a stetement to the document

	doc.Statements = append(doc.Statements, vex.Statement{
		Vulnerability:   "CVE-2014-123456",
		VulnDescription: "Its really bad",
		Products:        []string{"pkg:generic/1.0.0"},
		Subcomponents:   []string{},
		Status:          "fixed",
		StatusNotes:     "It works now",
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
  "timestamp": "2023-01-09T09:08:42-06:00",
  "version": "1.0",
  "tooling": "hello/vex",
  "supplier": "You!",
  "statements": [
    {
      "vulnerability": "CVE-2014-123456",
      "vuln_description": "Its really bad",
      "products": [
        "pkg:generic/1.0.0"
      ],
      "status": "fixed",
      "status_notes": "It works now"
    }
  ]
}

```

If you want to play around with this example, download the hello-vex.zip
file from this directory.
