# Merging VEX Documents

VEX [documents](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#document-1)
contain one or more [VEX statements](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#statement).

VEX statetements are designed to be standalone. No matter if they are moved to
another document, embedded or combined or split from equivalent statements their 
message and meaning remains the same: How a vulnerability is known to affect a
software product.

## Why Merge VEX Documents?

The self sustaining nature of the statments makes merging documents possible.
There are a couple of reasons why we may want to merge documents.

### Combining many Documents

VEX data describes the knowledge a vulnerability has over time. A VEX statement
expresses the known exploitabilty at a point in time and can be overriden by 
new documents that capture the latest impact information.

There can be more than one authoritative source of VEX information. The 
consumer of the VEX documents will ultimately be the judge on who they trust
whe it comes to reading VEX data. VEX documents can originate from a number
of issuers:

- The software author
- A distributor or repackager
- An independent scurity researcher
- An Auditor
- The internal security team in an organization

The VEX impact history can be assembled from multiple documents but sometimes
having just one document may be preferred. 

### Assembling a Product's Metadata out of Many Components'

Modern software products incorporate many components. Each of these components
can be considered a "product" in the VEX sense. Authors and other people may be
issuing VEX Statements about them. In order to generate a VEX document of 
a product with many components, we may need to pull into the document all known
information its components. 

Instead of providing many individual VEX documents, a software author may 
combine all the known VEX data of their dependencies into a VEX document cataloguing
known exploitability data. 

## Merging with `vexctl`

To combine documents, `vexctl` has a `merge` subcommand. The invocation is 
simple enough: when running it, simply pass all the documents you want to
merge and it will combine lla documents into one. This example can be run 
within the vexctl repo. 

Lets inspect the contents of the documents:

```console
cat pkg/ctl/testdata/document2.vex.json
```

```json
{
    "id": "my-vexdoc",
	"format": "text/vex+json",
	"author": "John Doe",
	"role": "vex issuer",	
	"statements": [
        {
            "timestamp": "2022-12-22T16:36:43-05:00",
            "products": ["pkg:apk/wolfi/bash@1.0.0"],
            "vulnerability": "CVE-1234-5678",
            "status": "under_investigation",
            "status_notes": ""
        }
    ]
}
```

And the second:

```console
cat pkg/ctl/testdata/document2.vex.json
```
{
    "id": "my-vexdoc",
	"format": "text/vex+json",
	"author": "John Doe",
	"role": "vex issuer",	
    "statements": [
        {
            "timestamp": "2022-12-22T20:56:05-05:00",
            "products": ["pkg:apk/wolfi/bash@1.0.0"],
            "vulnerability": "CVE-1234-5678",
            "status": "affected"
        }
    ]
}

```

As you can see, both documents containe one statement about the same product
(pkg:apk/wolfi/bash@1.0.0) and vulnerability (CVE-1234-5678).

Let merge both documents and see what happens:

```console
vexctl merge pkg/ctl/testdata/document1.vex.json \
             pkg/ctl/testdata/document2.vex.json 

```

Running the command will make vexctl output the combined document to stdout:

```json
{
  "@context": "",
  "@id": "merged-vex-67124ea942ef30e1f42f3f2bf405fbbc4f5a56e6e87684fc5cd957212fa3e025",
  "author": "Unknown Author",
  "role": "Document Creator",
  "timestamp": "2023-01-19T02:36:03.290252574-06:00",
  "version": "",
  "statements": [
    {
      "vulnerability": "CVE-1234-5678",
      "timestamp": "2022-12-22T16:36:43-05:00",
      "products": [
        "pkg:apk/wolfi/bash@1.0.0"
      ],
      "status": "under_investigation"
    },
    {
      "vulnerability": "CVE-1234-5678",
      "timestamp": "2022-12-22T20:56:05-05:00",
      "products": [
        "pkg:apk/wolfi/bash@1.0.0"
      ],
      "status": "affected"
    }
  ]
}
```
Note that statements are sorted in the new document. This lets the human eye
understand the evolution of the impact knwoledge.

## Specifying what to Merge

Merging documents can be done with a broad stroke, as in the example above, or
with finer-grained control. There are two flags that let the user control what
gets considered for the new document:

`--product`: If you pass a software identifier, only statements including said
identifier will be considered to be merged in the new document.

`--vulnerability`: If you pass a vulnerability ID, vexctl will only to the new 
document statements describing impact that specific vulnerability.
