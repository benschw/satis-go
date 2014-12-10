[![Build Status](https://drone.io/github.com/benschw/satis-go/status.png)](https://drone.io/github.com/benschw/satis-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/satis-go?status.png)](http://godoc.org/github.com/benschw/satis-go)


# satis-go


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
	wget -qO- -O /opt/satis/config.yaml https://drone.io/github.com/benschw/satis-go/files/config.yaml

	# Get/Install ui for satis-go server
	wget -qO- -O tmp.zip https://drone.io/github.com/benschw/satis-admin/files/admin-ui.zip
	unzip tmp.zip -d /opt/satis/
	rm tmp.zip


