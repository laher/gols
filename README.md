go-ls
=====

This Go-related utility allows you to filter out folders from Go tooling.

go-ls was created with the go '/vendor/' folder in mind, in combination with tools like `go test`

go-ls can be used on the commandline, or alternatively you can build its functionality into your own package. 

For library usage docs, see godoc.org/github.com/laher/gols

Note that gols is simply a wrapper around `go list`


Examples
--------

For example, to run tests while excluding the 'vendor' folder:

	go-ls -exec="go test -v" ./...

Or, you can just run tests (excluding /vendor/) like this:

	go-test ./...

Other options:

	go-ls -exec="go install" -ignore=/vendor/,/scripts/ ./...


Why
---

When you want to run projects



Installation
------------

1. To install the go-ls only:

	go get github.com/laher/gols/cmd/go-ls

2. OR, to install ALL of the tools (go-ls :

	go get github.com/laher/gols/cmd/...

3. OR you can just install one of the other tools:

	go get github.com/laher/gols/cmd/go-test

