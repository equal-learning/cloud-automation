## Overview

[[_TOC_]]

## Introduction

The Go package template implements **data-driven** templates for generating textual output.

Templates are executed by applying them to a data structure (usually confirming to yaml like syntax).
**Annotations** in the template refer to elements of the data structure (typically a field of a structure or a key in a map) to control execution and derive values to be displayed. Execution of the template walks the structure and sets the cursor, represented bu a period ‘.’ called **dot**, to the value at the current location in the structure as execution proceeds.

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

## Actions

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

{{if pipeline}} T1 {{else}} T0 {{end}}

If the value of the pipeline is empty, T0 is executed; otherwise, T1 is executed. Dot is unaffected.

{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}

To simplify the appearance of if-else chains, the else action of an if may include another if directly; the effect is exactly the same as writing {{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}

{{range pipeline}} T1 {{end}}

The value of the pipeline must be an array, slice, map, or channel. If the value of the pipeline has length zero, nothing is output; otherwise, dot is set to the successive elements of the array, slice, or map and T1 is executed. If the value is a map and the keys are of basic type with a defined order, the elements will be visited in sorted key order.

{{range pipeline}} T1 {{else}} T0 {{end}}

The value of the pipeline must be an array, slice, map, or channel. If the value of the pipeline has length zero, dot is unaffected and T0 is executed; otherwise, dot is set to the successive element of the array, slice, or map and T1 is executed.

{{break}}

The innermost {{range pipeline}} loop is ended early, stopping the current iteration and bypassing all remaining iterations.

{{continue}}

The current iteration of the innermost {{range pipeline}} loop is stopped, and the loop starts the next iteration.

{{template "name"}}

The template with the specified name is executed with nil data.

{{template "name" pipeline}}

The template with the specified name is executed with dot set to the value of the pipeline.

{{block "name" pipeline}} T1 {{end}}
	A block is shorthand for defining a template
		{{define "name"}} T1 {{end}}
	and then executing it in place
		{{template "name" pipeline}}
	The typical use is to define a set of root templates that are
	then customized by redefining the block templates within.

{{with pipeline}} T1 {{end}}
	
If the value of the pipeline is empty, no output is generated; otherwise, dot is set to the value of the pipeline and T1 is	executed.

{{with pipeline}} T1 {{else}} T0 {{end}}

If the value of the pipeline is empty, dot is unaffected and T0 is executed; otherwise, dot is set to the value of the pipeline and T1 is executed.

### Argument ::

An argument is a simple value, denoted by one of the following.

- A boolean, string, character, integer, floating-point, imaginary or complex constant in Go syntax.
- The keyword nil, representing an untyped Go nil.
- The character '.' (period): . The result is the value of dot.
- A variable name, which is a (possibly empty) alphanumeric string preceded by a dollar sign, such as $piOver2 or $ The result is the value of the variable.
- The name of a field of the data, which must be a struct, preceded by a period, such as .Field The result is the value of the field. Field invocations may be chained: .Field1.Field2
- The name of a key of the data, which must be a map, preceded by a period, such as .Key The result is the map element value indexed by the key. Key invocations may be chained and combined with fields to any depth: .Field1.Key1.Field2.Key2 
- The name of a niladic method of the data, preceded by a period, such as .Method 
- The name of a niladic function, such as fun The result is the value of invoking the function, fun(). 
- A parenthesized instance of one the above, for grouping. The result may be accessed by a field or map key invocation. print (.F1 arg1) (.F2 arg2) (.StructValuedMethod "arg").Field

Arguments may evaluate to any type; if they are pointers the implementation automatically indirects to the base type when required.

### Pipelines

A pipeline is a possibly chained sequence of "commands". A command is a simple value (argument) or a function or method call, possibly with multiple arguments:

A pipeline may be "chained" by separating a sequence of commands with pipeline characters '|'. In a chained pipeline, the result of each command is passed as the last argument of the following command. The output of the final command in the pipeline is the value of the pipeline.

### Variables

A pipeline inside an action may initialize a variable to capture the result. The initialization has syntax

```shell
$variable := pipeline
```

Example of templates demonstrating pipelines and variables ::

```shell
{{"\"output\""}}
	A string constant.
{{`"output"`}}
	A raw string constant.
{{printf "%q" "output"}}
	A function call.
{{"output" | printf "%q"}}
	A function call whose final argument comes from the previous
	command.
{{printf "%q" (print "out" "put")}}
	A parenthesized argument.
{{"put" | printf "%s%s" "out" | printf "%q"}}
	A more elaborate call.
{{"output" | printf "%s" | printf "%q"}}
	A longer chain.
{{with "output"}}{{printf "%q" .}}{{end}}
	A with action using dot.
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
	A with action that creates and uses a variable.
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
	A with action that uses the variable in another action.
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
	The same, but pipelined.
```

### Functions

During execution functions are found in two function maps: first in the template, then in the global function map. By default, no functions are defined in the template but the Funcs method can be used to add them.

Predefined global functions are named as follows.

```shell
and
	Returns the boolean AND of its arguments by returning the
	first empty argument or the last argument. That is,
	"and x y" behaves as "if x then y else x."
	Evaluation proceeds through the arguments left to right
	and returns when the result is determined.
call
	Returns the result of calling the first argument, which
	must be a function, with the remaining arguments as parameters.
	Thus "call .X.Y 1 2" is, in Go notation, dot.X.Y(1, 2) where
	Y is a func-valued field, map entry, or the like.
	The first argument must be the result of an evaluation
	that yields a value of function type (as distinct from
	a predefined function such as print). The function must
	return either one or two result values, the second of which
	is of type error. If the arguments don't match the function
	or the returned error value is non-nil, execution stops.
html
	Returns the escaped HTML equivalent of the textual
	representation of its arguments. This function is unavailable
	in html/template, with a few exceptions.
index
	Returns the result of indexing its first argument by the
	following arguments. Thus "index x 1 2 3" is, in Go syntax,
	x[1][2][3]. Each indexed item must be a map, slice, or array.
slice
	slice returns the result of slicing its first argument by the
	remaining arguments. Thus "slice x 1 2" is, in Go syntax, x[1:2],
	while "slice x" is x[:], "slice x 1" is x[1:], and "slice x 1 2 3"
	is x[1:2:3]. The first argument must be a string, slice, or array.
js
	Returns the escaped JavaScript equivalent of the textual
	representation of its arguments.
len
	Returns the integer length of its argument.
not
	Returns the boolean negation of its single argument.
or
	Returns the boolean OR of its arguments by returning the
	first non-empty argument or the last argument, that is,
	"or x y" behaves as "if x then x else y".
	Evaluation proceeds through the arguments left to right
	and returns when the result is determined.
print
	An alias for fmt.Sprint
printf
	An alias for fmt.Sprintf
println
	An alias for fmt.Sprintln
urlquery
	Returns the escaped value of the textual representation of
	its arguments in a form suitable for embedding in a URL query.
	This function is unavailable in html/template, with a few
	exceptions.
```

The boolean functions take any zero value to be false and a non-zero value to be true.

There is also a set of binary comparison operators defined as functions:

```shell
eq
	Returns the boolean truth of arg1 == arg2
ne
	Returns the boolean truth of arg1 != arg2
lt
	Returns the boolean truth of arg1 < arg2
le
	Returns the boolean truth of arg1 <= arg2
gt
	Returns the boolean truth of arg1 > arg2
ge
	Returns the boolean truth of arg1 >= arg2
```

### Nested template definitions

When parsing a template, another template may be defined and associated with the template being parsed. Template definitions must appear at the top level of the template, much like global variables in a Go program.

The syntax of such definitions is to surround each template declaration with a "define" and "end" action.

The define action names the template being created by providing a string constant. Here is a simple example:









Reference : [Go Template Doc ](https://pkg.go.dev/text/template)