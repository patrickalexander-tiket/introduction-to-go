# 5-golang-Linter
Just like in Javascript that using ESLint, Python that using Pylint, Golang also has some code quality tool to standarize :
- Coding Standard 
- Find Unused Code
- Error detection
- Continuous integration
- Flag programming errors, 
- Detect unused struckts
- Suspicious constructs.
- and many more..

## 1. Go-Report-Card 
### Installation
https://github.com/gojp/goreportcard

### Generate Repository Score
Example : https://goreportcard.com/report/github.com/kubernetes/kubernetes

### Feature :
- gofmt
- [go_vet](https://golang.org/cmd/vet/) 
- golint 
- gocyclo 
- ineffassign 
- license
- misspell

## 2. GolangCi-Lint (Preferable)
### Installation
```
> $ go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0
> $ golangci-lint run
```
### Enabled By Default Linters
- deadcode - Finds unused code
- errcheck - Errcheck is a program for checking for unchecked errors in go programs. These unchecked errors can be critical bugs in some cases
- gosimple - Linter for Go source code that specializes in simplifying a code
- [govet](https://golang.org/cmd/vet/) - Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string
- ineffassign - Detects when assignments to existing variables are not used
- staticcheck - Staticcheck is a go vet on steroids, applying a ton of static analysis checks
- structcheck - Finds unused struct fields
- typecheck - Like the front-end of a Go compiler, parses and type-checks Go code
- unused - Checks Go code for unused constants, variables, functions and types
- varcheck - Finds unused global variables and constants

### Disabled By Default Linters (-E/--enable)
- asciicheck - Simple linter to check that your code does not contain non-ASCII identifiers
- bodyclose - checks whether HTTP response body is closed successfully
- cyclop - checks function and package cyclomatic complexity
- and many more...

Ref : https://golangci-lint.run/usage/linters/

### Company using GolangCi-Lint
https://golangci-lint.run/product/trusted-by/
- AWS
- Facebook
- Google
- Netflix
- Arduino
- Baidu
- Eclipse Foundation
- IBM
- Istio
- Percona
- Red Hat OpenShift
- Samsung
- Serverless
- ScyllaDB
- SoundCloud
- The New York Times
- WooCart
- Xiaomi
- Yahoo