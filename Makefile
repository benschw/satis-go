default: build

clean:
	rm -rf repo-ui
	rm -rf admin-ui
	rm -rf satis-go
	rm -rf data
	rm -rf lib
	rm -rf conposer.phar

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


.PHONY: satis admin-ui
