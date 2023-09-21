# Create a New OpenVEX Document

Creating a new VEX document with `vexctl` is easy. Our tool takes care of
document creation for you, filling the required fields of the OpenVEX
document with sensible defaults.

When starting a new document, `vexctl` will create a document with one 
vex statement.

## TL;DR;

To generate a new OpenVEX document simply run `vexctl create`
with a product, a vulnerability and a status:

```console
vexctl create "pkg:generic/product@1.0.0" CVE-2014-123456 fixed 
{
  "@context": "https://openvex.dev/ns/v0.2.0",
  "@id": "https://openvex.dev/docs/public/vex-091509fa4db8e7948456630cbccf3ebd957b648a16d471d78c84b43040f49a60",
  "author": "Unknown Author",
  "timestamp": "2023-09-20T11:42:12.074772+02:00",
  "version": 1,
  "statements": [
    {
      "vulnerability": {
        "name": "CVE-2014-123456"
      },
      "timestamp": "2023-09-20T11:42:12.07479+02:00",
      "products": [
        {
          "@id": "pkg:generic/product@1.0.0"
        }
      ],
      "status": "fixed"
    }
  ]
}
```

For a deeper understanding of what's going on, keep reading!

## Required Data

A VEX statement describes the impact a vulnerability has on a piece of software.
You will need at least three pieces of information to create the new document:

### 1. A Software Identifier

In order to specify the "product" (i.e. the software we are referring to) you will
need a software identifier that points to it. Several identifying
schemes exist, but OpenVEX recommends the use [package urls](https://github.com/package-url/purl-spec) (purls).

### 2. A Vulnerability ID

The second piece of required data is the vulnerability. The
ID is a string that uniquely identifies the vulnerability and that can be looked 
up in a tracking system. Typically this will be a [CVE](https://cve.mitre.org/)
number or similar.

### 3. Impact Knowledge

_This is the most important part and the message that VEX is designed to convey._

Given the above two pieces of data (Software & Vulnerability Identifiers), a VEX statement adds information about whether
or not the vulnerability actually has impact on a software product, and if so, the status of the impact. This allows any
downstream consumers to programmatically understand whether a vulnerability affects your software.

VEX defines a set of impact status labels, you need to chose the one that
best describes the message you want to express:

| Label | Description |
| --- | --- |
| `not_affected` | No remediation is required regarding this vulnerability. |
| `affected` | Actions are recommended to remediate or address this vulnerability. |
| `fixed` | These product versions contain a fix for the vulnerability. |
| `under_investigation` | It is not yet known whether these product versions are affected by the vulnerability. Updates should be provided in further VEX documents, as knowledge evolves. |

### Other Required Fields

There are some required fields when generating statements with certain statuses.
When possible, `vexctl` will provide some sane defaults.

- A `not_affected` status requires the addition of a `justification` to the
statement. Like status labels, justifications are fixed labels. Refer to the
spec for the complete
[status justifications list](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#status-justifications).

- An `affected` status must provide an action statement.

## Adding More Data

`vexctl` lets you add more data to the statement. The following flags are
available at the time of writing:

```console
vexctl add
...
  -a, --action-statement string   action statement for "affected" status (only when status=affected) (default "No action statement provided")
      --file string               file to write the document to (default is STDOUT)
  -h, --help                      help for add
      --impact-statement string   text explaining why a vulnerability cannot be exploited (only when status=not_affected)
  -i, --in-place                  add a statement to an existing file
  -j, --justification string      justification for "not_affected" status (see vexctl list justification)
  -p, --product string            main identifier of the product, a package URL or another IRI
  -s, --status string             impact status of the product vs the vulnerability
      --status-note string        statement on how status was determined
      --subcomponents strings     list of subcomponents to add to the statement, package URLs or other IRIs
  -v, --vuln string               vulnerability to add to the statement (eg CVE-2023-12345)
```

You can specify more than one product. `vexctl` will read one from
the argument but you can control all parameters through command line
flags. Here's an example invocation with two products in the same document:

```
vexctl create --product="pkg:apk/wolfi/git@2.39.0-r1?arch=x86_64" \
              --product="pkg:apk/wolfi/git@2.39.0-r1?arch=armv7" \
              --vuln="CVE-2023-12345" \
              --status="fixed"
```
 
