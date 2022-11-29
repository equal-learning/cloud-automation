## Overview

[[_TOC_]]


## Introduction ::
Go's template designe to be **extended** and provides access to **data objects** and **additional functions** that are **passed** into the template engine **programmatically**.

## A tour of basic Go template ::

```script
This is a simple template
It can {{ "output" }} something.
It also
{{- " demonstrates" }} trim markers.
{{/* it has a comment */}}
It can {{ "output" }} something.
It can demonstrate {{ "output" | print }} using pipelines.
It also {{ $A := "assigns variables" }}{{ $A }}.
And conditionals:
{{ $B := 2 }}{{ if eq $B 1 }}B is 1{{ else }}B is 2{{ end }}
```

### Basic terminology ::

#### Template
It is the **entire text** that is submitted to the **template engine** to be rendered. The template consists of **text** and **actions**.
All text **outside** of actions is copied to the output **unchanged**.







