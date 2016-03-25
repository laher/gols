go-ls
=====

This Go-related utility allows you to filter out folders from commandline 'wildcards' such as `./...`

go-ls was created with the go '/vendor/' folder in mind, in combination with tools like `go test`

go-ls can be used on the commandline, or alternatively you can build its functionality into your own package. 

For library usage docs, see godoc.org/github.com/laher/gols

Note that gols is simply a wrapper around `go list`


Why
---

Because lazy.

The use of /vendor/ folders has affected the use of ./... wildcards in go commands.

The Go team recommend the following approach to ignoring this folder:

	go install $(go list ./... | grep -v /vendor/)

gols is basically a more convenient way to do the same thing

See bug report https://github.com/golang/go/issues/11659 for background.


Examples
--------

For example, to run tests while excluding the 'vendor' folder:

	go-test ./...

To simply list the filtered packages:

	go-ls ./...

Other options:

	go-ls -exec="go test -v" ./...
	go-ls -exec="go install" -ignore=/vendor/,/scripts/ ./...


Installation
------------

1. To install the go-ls only:

	go get github.com/laher/gols/cmd/go-ls

2. OR, to install ALL of the tools (go-ls, go-test, go-install):

	go get github.com/laher/gols/cmd/...

3. OR you can just install one of the other tools:

	go get github.com/laher/gols/cmd/go-test

