# Create a New OpenVEX Document

Creating a new VEX document with `vexctl` is easy. Out tool will take care of
the document creation for you, trying to fill the required fields of the OpenVEX
document with sensible defaults.

When starting a new document, `vexctl` will create a document with one 
vex statement.

## TL;DR;

Let's skip to the end. To generate a new OpenVEX document simply run
`vexctl create` with a product, a vulnerability and a status:

```console
vexctl create "pkg:generic/product@1.0.0" CVE-2014-123456 fixed 
```

That invocation of vexctl will produce the following document:

```json
{
  "@context": "https://openvex.dev/ns",
  "@id": "https://openvex.dev/docs/public/vex-11aeb821bd5ca6f4fcbfce111897f81c5dcee6a781ea364e8e35be4443a1c517",
  "author": "Unknown Author",
  "role": "Document Creator",
  "timestamp": "2023-01-19T01:34:03.27761931-06:00",
  "version": "1",
  "statements": [
    {
      "vulnerability": "CVE-2014-123456",
      "products": [
        "pkg:generic/product@1.0.0"
      ],
      "status": "fixed"
    }
  ]
}
```

TO understand deeper what's going on keep reading!

## Required Data

You will need at least three pieces of information to create the new document:

### 1. A Software Identifier

VEX statements describe the impact a vulnerability has on a piece of software.
In order to specify the "product", ie the software we are referring to, you will
need a software identifier that points to it. There are several identifying
schemes but OpenVEX recomends the use [package urls](https://github.com/package-url/purl-spec) (purls).

### 2. A Vulnerability ID

The second piece of data is the vulnerability associated with the impact. The
ID is a string that uniquely identifies the vulnerability and that can be looked 
up in a tracking system. Typically this will be a [CVE](https://cve.mitre.org/)
number or similar.

### 3. Impact Knowledge

This is the most important part and the message that VEX is designed to convey.
A vex statment will inform you consumers downstream how the vulnerability 
affects your software. 

VEX defines a set of fixed impact status labes, you need to chose the one that
best describes the message you want to express:

| Label | Description |
| --- | --- |
| `not_affected` | No remediation is required regarding this vulnerability. |
| `affected` | Actions are recommended to remediate or address this vulnerability. |
| `fixed` | These product versions contain a fix for the vulnerability. |
| `under_investigation` | It is not yet known whether these product versions are affected by the vulnerability. Updates should be provided in further VEX documents as knowledge evolves. |

#### Other Required Fields

There are some required fields whe generating statements with certain statuses.
When possible, vexctl will provide some defaults, these will often mean nothing
but are provided to compose a valid document.

- A `not_affected` status requires the addition of a `justification` to the
statement. Like status labels, justifications are fixed labels. Refer to the
spec for the complete
[status justifications list](https://github.com/openvex/spec/blob/main/OPENVEX-SPEC.md#status-justifications).

- An `affected` status must provide an action statement.

## Adding More Data

vexctl lets you add more data to the statement. The following flagas are
avilable at the time of writing:

```console
  -a, --action-statement string   action statement for affected status (default "No action statement provided")
      --author string             author to record in the new document (default "Unknown Author")
      --author-role string        author role to record in the new document (default "Document Creator")
      --file string               file to write the document (default is STDOUT)
  -h, --help                      help for create
      --id string                 ID for the new VEX document (default will be computed)
  -j, --justification string      justification for not_affected status, see 'vexctl show justifications' for list
  -p, --product strings           list of products to list in the statement, at least one is required
  -s, --status string             status of the product vs the vulnerability, see 'vexctl show statuses' for list
      --subcomponents strings     list of subcomponents to add to the statement
  -v, --vuln string               vulnerability to add to the statement (eg CVE-2023-12345)
```

You can specify more than one product. vexctl will read one from
the argument but you can control all parameters through command line
flags. Here's an example invocation with two products in the same document:

```
vexctl create --product="pkg:apk/wolfi/git@2.39.0-r1?arch=x86_64" \
              --product="pkg:apk/wolfi/git@2.39.0-r1?arch=armv7" \
              --vuln="CVE-2023-12345" \
              --status="fixed"
```
 
