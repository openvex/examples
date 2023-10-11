package main

import (
	"log"
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
			Version:    1,
			Tooling:    "hello/vex",
			Supplier:   "You!",
		},
	}

	/// Add a stetement to the document

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
		Status:      vex.StatusFixed,
		StatusNotes: "It works now",
	})

	err := doc.ToJSON(os.Stdout)
	if err != nil {
		log.Fatalf("Error writing JSON: %v", err)
	}
}
