[![Build Status](https://drone.io/github.com/benschw/satis-go/status.png)](https://drone.io/github.com/benschw/satis-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/satis-go?status.png)](http://godoc.org/github.com/benschw/satis-go)


# Satis-go
[download latest](https://drone.io/github.com/benschw/satis-go/files/satis-go)

Satis-go is a web server for hosting and managing your [Satis Repository](https://github.com/composer/satis) for [Composer Packages](https://getcomposer.org/)

Some Highlights:
* Satis-go provides a simple user interface for managing the repositories you want to track in your Composer package repo
* Repo generation is delegated to [Satis](https://github.com/composer/satis) so your package repository will stay up to date with composer specs
* No database required: the satis config file is managed directly while still managing writes and reads safely
* RESTful API so you and integrate this into your CI




## Install
Here is how to install [Satis](https://github.com/composer/satis) to /opt/satis and `satis-go` to /opt/satis-go/
	
	# Get Composer/Satis and install in path
	apt-get install -y php5-cli
	curl -sS https://getcomposer.org/installer | php
	php ./composer.phar create-project composer/satis /opt/satis --stability=dev --keep-vcs
	ln -s /opt/satis/bin/satis /usr/local/bin/satis
	
	# Setup install dir
	mkdir /opt/satis-go

	# Get/Install satis-go server & config
	wget -qO- -O /opt/satis-go/satis-go https://drone.io/github.com/benschw/satis-go/files/satis-go 
	chmod +x /opt/satis-go/satis-go
	wget -qO- -O /opt/satis-go/config.yaml https://drone.io/github.com/benschw/satis-go/files/config.yaml

	# Get/Install ui for satis-go server
	wget -qO- -O tmp.zip https://drone.io/github.com/benschw/satis-admin/files/admin-ui.zip
	unzip tmp.zip -d /opt/satis-go/

	# Cleanup
	rm ./composer.phar
	rm tmp.zip


## Usage

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
