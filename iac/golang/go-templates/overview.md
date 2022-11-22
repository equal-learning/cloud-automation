## Overview

[[_TOC_]]

## Introduction

The Go package template implements **data-driven** templates for generating textual output.

Templates are executed by applying them to a data structure (usually confirming to yaml like syntax).
**Annotations** in the template refer to elements of the data structure (typically a field of a structure or a key in a map) to control execution and derive values to be displayed. Execution of the template walks the structure and sets the cursor, represented bu a period ‘.’ called ‘dot’, to the value at the current location in the structure as execution proceeds.

The INPUT text for a template is UTF-8-encoded text in any format.
Reminder, the **input** text can have dynamic content, meant to be processed by the template and the static content,
meant to be coppied as it is.
**Actions** are data evaluation or control structure, delimited by {{ }}.
All text outside actions is copied to the output unchanged.

## A simple go template

````gotemplate
type Inventory struct {
	Material string
	Count    uint
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
````

# Blank Spaces (space, horizontal tab, carriage return, and newline)
By default, all text between action (by default delimiter by ‘{{’ and ‘}}’) is copied verbatim when the template is executed. However this behavior can be altered as bellow,
{{-  : A minus sign and white space, will trim the trailing white space from the immediately preceding text.
 -}}  : A white space and a minus sign, all leading white space is trimmed from the immediately following text.
 
````gotemplate
"{{23 -}} < {{- 45}}"
````
will be formatted to,

````gotemplate
"23<45"
````

# Actions

**Actions** are data evaluation or control structure, delimited by {{ }}.
All text outside actions is copied to the output unchanged.  
Action can contain "arguments" and "pipelines", which are **evalutions** of data.

List of actions ::

{{/* a comment */}}

{{- /* a comment with white space trimmed from preceding and following text */ -}}

{{pipeline}}

The default textual representation (the same as would be printed by fmt.Print) of the value of the pipeline is copied
to the output.

{{if pipeline}} T1 {{end}}

If the value of the pipeline is empty, no output is generated; otherwise, T1 is executed. The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero. Dot is unaffected.










Reference : [Go Template Doc ](https://pkg.go.dev/text/template)