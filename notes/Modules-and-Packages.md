# Modules and Packages

## General Rules
Every `.go` file has to be apart of some "package", this is done by declaring `package <package-name` at the begining of the file.

Go will not allow importing unsued packages to bloat your code.

modules are third party libraries, which live in `$GOROOT/pkg`

The case of a function defined in a file dictates wether it is "exported" (publicly accessible) or ont to the rest of the package.

## installing ?
```bash
go get <package-name>
```

## Initializing a `go.mod`
```bash
go mod init <url-of-repo-where-yuo are-hosting-the-code>
```

this will create a `go.mod` file in your project
```mod
module <url-of-repo-where-yuo are-hosting-the-code>

go 1.20
```

## Add additional packages a Go program
to add additional 3rd party libraries to your project you will need to add the url of the package to your `.go` file "import"
```go
import (
    <url-of-some-package>
    packagealias <url-of-some-other-package>
)
```
the above example also shows how to alias package names (equivalent to `import numpy as np` in python)

and then run 
```bash
go mod tidy
```
 to automatically update your `go.mod` file to "require" the added package url

 ## Package Prerequisites
 These can be found in the `go.sum` file