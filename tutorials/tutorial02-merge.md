# Merging VEX Documents

VEX [documents](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#document-1)
contain one or more [VEX statements](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#statement).

VEX statements are designed to be standalone. No matter if they are moved to
another document, embedded, combined or split from equivalent statements – their 
message and meaning remains the same: How a vulnerability is known to affect a
software product.

## Why Merge VEX Documents?

The self sustaining nature of the statements makes merging documents possible.
There are a couple of reasons why we may want to merge documents.

### Combining many Documents

VEX data describes the knowledge a vulnerability has _over time_. A VEX statement
expresses the known exploitability at a point in time and can be overridden by 
new documents that capture the latest impact information.

There can be more than one authoritative source of VEX information. The 
consumer of the VEX documents will ultimately be the judge on who they trust
when it comes to reading VEX data. VEX documents can originate from a number
of issuers:

- The software author
- A distributor or repackager
- An independent security researcher
- An Auditor
- The internal security team in an organization

The VEX impact history can be assembled from multiple documents but sometimes
having just one document may be preferred. 

### Assembling a Product's Metadata out of Many Components

Modern software products incorporate many components. Each of these components
can be considered a "product" in the VEX sense. Authors and other people may be
issuing VEX Statements about them. To generate a VEX document for a product with
multiple components, it may be necessary to aggregate all known information
 about these subcomponents into a single document.

Instead of providing many individual VEX documents, a software author may 
combine all the known VEX data of their dependencies into a single VEX document,
cataloging all known exploitability data. 

## Merging with `vexctl`

To combine documents, `vexctl` has a `merge` subcommand. The invocation is 
simple enough: when running it, simply pass all the documents you want to
merge and it will combine all documents into one.

Let's generate inspect the contents of some example documents:

```console
vexctl create "pkg:generic/product@1.0.0" CVE-1234-5678 under_investigation | tee document1.vex.json
```

```json
{
  "@context": "https://openvex.dev/ns/v0.2.0",
  "@id": "https://openvex.dev/docs/public/vex-bf2d70a3a8f4cb887a0996e49977fa2529f2d93818d156f017cbbebb25642692",
  "author": "Unknown Author",
  "timestamp": "2023-09-20T11:46:28.949091+02:00",
  "version": 1,
  "statements": [
    {
      "vulnerability": {
        "name": "CVE-1234-5678"
      },
      "timestamp": "2023-09-20T11:46:28.949093+02:00",
      "products": [
        {
          "@id": "pkg:generic/product@1.0.0"
        }
      ],
      "status": "under_investigation"
    }
  ]
}
```

And the second:

```console
vexctl create "pkg:generic/product@1.0.0" CVE-1234-5678 affected | tee document2.vex.json 
```

```json
{
  "@context": "https://openvex.dev/ns/v0.2.0",
  "@id": "https://openvex.dev/docs/public/vex-4af5963fd3ca9747c209da769700611c089ce7249be45cbd0fe1f4ed16679530",
  "author": "Unknown Author",
  "timestamp": "2023-09-20T11:48:20.870254+02:00",
  "version": 1,
  "statements": [
    {
      "vulnerability": {
        "name": "CVE-1234-5678"
      },
      "timestamp": "2023-09-20T11:48:20.870256+02:00",
      "products": [
        {
          "@id": "pkg:generic/product@1.0.0"
        }
      ],
      "status": "affected",
      "action_statement": "No action statement provided",
      "action_statement_timestamp": "2023-09-20T11:48:20.870256+02:00"
    }
  ]
}
```

As you can see, both documents contain one statement about the same product
(pkg:generic/product@1.0.0) and vulnerability (CVE-1234-5678).

Let's merge both documents and see what happens:

```console
vexctl merge document1.vex.json \
             document2.vex.json
```

Running the command will make vexctl output the combined document to stdout:

```json
{
  "@context": "https://openvex.dev/ns/v0.2.0",
  "@id": "merged-vex-bd1900507c34eb17c532fb3fabd6904b48fe2a07cb0d6d3b734dbd8dd837dacb",
  "author": "Unknown Author",
  "timestamp": "2023-09-20T11:49:15.587679+02:00",
  "version": 1,
  "statements": [
    {
      "vulnerability": {
        "name": "CVE-1234-5678"
      },
      "timestamp": "2023-09-20T11:47:29.038232+02:00",
      "products": [
        {
          "@id": "pkg:generic/product@1.0.0"
        }
      ],
      "status": "under_investigation"
    },
    {
      "vulnerability": {
        "name": "CVE-1234-5678"
      },
      "timestamp": "2023-09-20T11:48:20.870256+02:00",
      "products": [
        {
          "@id": "pkg:generic/product@1.0.0"
        }
      ],
      "status": "affected",
      "action_statement": "No action statement provided",
      "action_statement_timestamp": "2023-09-20T11:48:20.870256+02:00"
    }
  ]
}
```
Note that statements are sorted in the new document. This lets the human eye to
understand the evolution of the impact knowledge.

## Specifying what to Merge

Merging documents can be done with broad strokes, as in the example above, or
with finer-grained control. There are two flags that let the user control what
gets considered for the new document:

`--product`: If you pass a software identifier, only statements including said
identifier will be considered to be merged in the new document.

`--vulnerability`: If you pass a vulnerability ID, vexctl will only to the new 
document statements describing impact that specific vulnerability.
