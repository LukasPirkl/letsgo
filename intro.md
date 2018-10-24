To get amazing introduction and historical context watch __The Why of Go__ - https://www.youtube.com/watch?v=bmZNaUcwBt4

Go is bad
---------

* too simple / lack of syntactic sugar
* no generics
* no macro/templates
* bad dependency management
* stuck in 80s
* error handling
* too opinionated
* too vebose
* no ternary oparator

source: https://github.com/ksimka/go-is-not-good


Go is good
----------
* designed by Google for real humans / real problems (their problems)
* object oriented but no classes and no inheritance
* small and simple - easy to learn
    * https://golang.org/ref/spec
    * https://golang.org/pkg/builtin/
* compiled, strongly typed, garbage collected language
* strict rules
    * code formatting - no space for discussions like tabs vs. spaces
    * best practices - https://golang.org/doc/effective_go.html
* build-in documentation
    * comments on public methods
    * warning when they are missing (impossible to turn in off)
* build-in checks
    * __gofmt__ - linter
    * __go run -race__ - runtime race condition check
* build-in unit test framework with profiling
* "Batteries included" - big standard library
* super easy cross-compilation - just environment variable
* self-contained binary (with run-time within) - one file to deploy


Usage of Go
-----------
* web apps
    * golang.org (obviously)
* microservices
    * Monzo (UK bank - https://www.youtube.com/watch?v=y2j_TB3NsRc)
* devops scripts
    * Azure SDK (https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/quickstarts/deploy-vm/main.go)
* cli tools (no GUI)
    * Docker
    * Terraform
    * Kubernetes
    * Hugo (fastest static site generator)
