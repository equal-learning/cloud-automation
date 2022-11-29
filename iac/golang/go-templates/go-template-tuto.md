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

### Actions

Actions provide the dynamic components of the template. They are delimited using {{ }}, as illustrated bellow :

```script
{{ "output" }}
{{- " demonstrates" }}
{{/* it has a comment */}}
{{ "output" | print }}
{{ $A := "assigns variables" }}{{ $A }}.
{{ $B := 2 }}{{ if eq $B 1 }}{{ else }}{{ end }}
```

Actions are **composed** of **control structure** or **data evaluation** through **pipelines**.

### Delimiters

Actions are set between a left delimiter **{{ by default** and  a right delimiter **}} by default**. The implementation can use **[[** and **]]** alternatively.

### Trim markers
Go template outputs everything between actions, including whitespace and line feeds. This can make it challenging to read **nested** template code.
If an actions’s left delimiter is followed immediately by a minus sign and ASCII space character **{{-**, all trailing white space is trimmed from the immediately preceding text. Likewise, if the right delimiter is preceded by a space and minus sign **-}}**, all leading white space id trimmed from the immediately following text.
Note, the ASCII space must be present in these trim markers.

### Comments
Comments are a special action created using a **/*** immediately following the left delimiter and a ***/** immediately preceding the right delimiter.

```script
{{/* it has a comment */}}
```

Comments can span multiple lines.

Note, comment out sections of a template while you are debugging it to simplify troubleshooting.







