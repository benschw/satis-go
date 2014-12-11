SHELL=/bin/bash

default: build

clean:
	rm -rf repo-ui
	rm -rf admin-ui
	rm -rf satis-go
	rm -rf data
	rm -rf golang-crosscompile
	rm -rf release

deps:
	go get
	go get gopkg.in/check.v1

satis-install:
	curl -sS https://getcomposer.org/installer | php
	php ./composer.phar create-project composer/satis /opt/satis --stability=dev --keep-vcs
	ln -s /opt/satis/bin/satis /usr/local/bin/satis
	rm ./composer.phar

admin-ui:
	curl -sS https://drone.io/github.com/benschw/satis-admin/files/admin-ui.tar.gz | tar xzv


dist: deps golang-crosscompile golang-buildsetup
	source golang-crosscompile/crosscompile.bash; \
	mkdir -p release; \
	go-darwin-386 build -o satis-go; \
	gzip -c satis-go > release/satis-go-darwin-386.gz; \
	go-darwin-amd64 build -o satis-go; \
	gzip -c satis-go > release/satis-go-darwin-amd64.gz; \
	go-linux-386 build -o satis-go; \
	gzip -c satis-go > release/satis-go-linux-386.gz; \
	go-linux-amd64 build -o satis-go; \
	gzip -c satis-go > release/satis-go-linux-amd64.gz

golang-buildsetup: golang-crosscompile
	source golang-crosscompile/crosscompile.bash; \
	go-crosscompile-build darwin/386; \
	go-crosscompile-build darwin/amd64; \
	go-crosscompile-build linux/386; \
	go-crosscompile-build linux/amd64

golang-crosscompile:
	git clone https://github.com/davecheney/golang-crosscompile.git


.PHONY: admin-ui
