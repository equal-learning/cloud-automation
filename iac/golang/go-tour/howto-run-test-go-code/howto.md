
# Install go and setup your workspace ::

How to install go locally :: https://go.dev/doc/install

List of important environment variables ::

GOROOT (like JAVA_HOME) = GOROOT is a variable that defines where your Go SDK is located. You **do not need to change** this variable, unless you plan to use different Go versions. You can modify it in .bash_profile (export GOROOT="/usr/local/go/")

GOPATH = GOPATH is a variable that defines the root of your workspace. By default, the workspace directory is a directory that is named go within your user home directory. GOPATH is the root of your ***workspace*** and contains the following folders:

src/: location of Go source code (for example, .go, .c, .g, .s).
pkg/: location of compiled package code (for example, .a).
bin/: location of compiled executable programs built by Go.


# Code organizatin in a Go program ::

repository
  + module (go.mod)
      + package
	        +program (source code)

Go "programs" are organized into packages.
A "package" is a collection of source files in the same directory that are compiled together.
Functions, types, variables, and constants defined in one source file are **visible** to all other source files within the same package.

A "module" is a collection of related Go packages that are released together.

A "repository" contains one or more modules. A Go repository typically contains only one module, located at the root of the repository.

A "module path", the import path prefix for all packages within the module.

A "import path", a string used to import a package. A package's import path is its module path joined with its subdirectory within the module.
Ex - For example, the module github.com/google/go-cmp contains a package in the directory cmp/. That package's import path is github.com/google/go-cmp/cmp.

Note : Packages in the standard library do not have a module path prefix.



# Create your first module
mkdir hello  
cd hello  
go mod init hello    # create your first module  

Note : If you are using IDE like VSCODE, then you need to add the module folder to the **workshop**. So everytime you create a module, you need to add it to the ***workshpace** !!! It is a bit of pain for the moment !!! 


hello.go

``
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}

``

# Build & install your program

go install full_path/hello   # build & install the program : Produces an **executable binary**". It then installs that binary as $HOME/go/bin/hello

Alternatively, cd example; cd hello; go install .  

Commands like go install apply within the context of the module containing the current working directory. If the working directory is not within the hello module, go install may fail.

For convenience, go commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given. So in our working directory, the following commands are all equivalent:

$ go install absolute_pathhello  // $HOME/path_to_hello
$ go install hello
$ cd exmaple; cd hello
$ go install .
$ go install


The install directory is controlled by the GOPATH and GOBIN [environment variables](https://pkg.go.dev/cmd/go#hdr-Environment_variables). If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).

Set your GO environment variable (as per convenient)

Here is an example of setting, GOBIN environment variable,
go env -w GOBIN=/somewhere/else/bin    # settting  
go env -u GOBIN   # unsetting  


# Run your first program  

$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$

# Understanding importpath

The go command locates the repository containing a given module path by requesting a corresponding HTTPS URL and reading metadata embedded in the HTML response [see go help importpath](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths) . Many hosting services already provide that metadata for repositories containing Go code, so the easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.


# Import your first package from a local module

Let us write a package first and later import it !

```
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

```

Because our ReverseRunes function begins with an ***upper-case letter***, it is exported, and can be used in other packages that import our morestrings package.


$ cd path_to/morestrings
$ go build  

This won't produce an output file. Instead it saves the compiled package in the local build cache.

Now import the above package in our hello program,

````
package main

import (
    "fmt"
    "hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}

`````

$ go install absolute_path/example/hello
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello


#  Import your first package from a remote module

An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories. For instance, to use github.com/google/go-cmp/cmp in your program:

````
package main

import (
    "fmt"

    "hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

````

Now that you have a dependency on an external module, you need to download that module and record its version in your go.mod file. The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.

$ go mod tidy

go: finding module for package github.com/google/go-cmp/cmp
go: downloading github.com/google/go-cmp v0.5.8
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.8

Module "dependencies (libraries)" are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the ***GOPATH*** environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. 

To ***remove all downloaded modules***, you can pass the -modcache flag to go clean:

go clean -modcache

# Write your first test in go

Go has a lightweight test framework composed of the go test command and the testing package.

You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.

````

package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}

````

Run the test :

$ cd $HOME/hello/morestrings
$ go test


Run go [help test](https://pkg.go.dev/cmd/go#hdr-Test_packages) and see the testing package [documentation](https://pkg.go.dev/testing) for more detail.


Reference :: https://go.dev/doc/code

