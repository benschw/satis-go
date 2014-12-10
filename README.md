[![Build Status](https://drone.io/github.com/benschw/satis-go/status.png)](https://drone.io/github.com/benschw/satis-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/satis-go?status.png)](http://godoc.org/github.com/benschw/satis-go)


# Satis-go

Satis-go is a web server for hosting and managing your [Satis Repository](https://github.com/composer/satis) for [Composer Packages](https://getcomposer.org/)

Some Highlights:
* Satis-go provides a simple user interface for managing the repositories you want to track in your Composer package repo
* Repo generation is delegated to [Satis](https://github.com/composer/satis) so your package repository will stay up to date with composer
* No database required: the satis config file is managed directly while still managing writes and reads safely
* RESTful API so you and integrate this into your CI


[download latest](https://drone.io/github.com/benschw/satis-go/files/satis-go)


## Install
	
	# Setup install dir
	mkdir /opt/satis

	# Get/Install Satis binary
	apt-get install -y php5-cli
	curl -sS https://getcomposer.org/installer | php
	php ./composer.phar create-project composer/satis /opt/satis/satis --stability=dev --keep-vcs

	# Get/Install satis-go server
	wget -qO- -O /opt/satis/satis-go https://drone.io/github.com/benschw/satis-go/files/satis-go 
	chmod +x /opt/satis/satis-go
	wget -qO- -O /opt/satis/config.yaml https://drone.io/github.com/benschw/satis-go/files/config.yaml

	# Get/Install ui for satis-go server
	wget -qO- -O tmp.zip https://drone.io/github.com/benschw/satis-admin/files/admin-ui.zip
	unzip tmp.zip -d /opt/satis/

	# Cleanup
	rm ./composer.phar
	rm tmp.zip


## Usage

### Start the server

	/opt/satis/satis-go

### Manage your satis repo

Just navigate your browser to (http://localhost:8080/admin) and start adding repos. They will automatically populate in your custom repo: (http://localhost:8080)

### Set up a hook script

Use the REST api to refresh your repository:

	curl -X POST http://localhost:8080/api/generate-web-job


