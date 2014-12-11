[![Build Status](https://drone.io/github.com/benschw/satis-go/status.png)](https://drone.io/github.com/benschw/satis-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/satis-go?status.png)](http://godoc.org/github.com/benschw/satis-go)


# Satis-go
[download latest](https://drone.io/github.com/benschw/satis-go/files)

Satis-go is a web server for hosting and managing your [Satis Repository](https://github.com/composer/satis) for [Composer Packages](https://getcomposer.org/)

Some Highlights:
* Satis-go provides a simple user interface for managing the repositories you want to track in your Composer package repo
* Repo generation is delegated to [Satis](https://github.com/composer/satis) so your package repository will stay up to date with composer specs
* No database required: the satis config file is managed directly while still managing writes and reads safely
* RESTful API so you and integrate this into your CI




## Getting Started

### Install 
[More here](http://txt.fliglio.com/satis-go/)

### Start the server

	/opt/satis-go/satis-go

### Manage your satis repo

Just navigate your browser to (http://localhost:8080/admin) and start adding repos. They will automatically populate in your custom repo: (http://localhost:8080)

### Set up a web-hook script to refresh your repo on commit/push

Use the REST api to refresh your repository:

	curl -X POST http://localhost:8080/api/generate-web-job


## Hacking

install satis to your path like above (or use the supplied `make` target)

	make satis-install

get a copy of the admin ui; you can store this in your checkout of satis-go

	make admin-ui

get your go deps

	make deps

start building

	go test ./...
	go build
	./satis-go -config config-local.yaml
