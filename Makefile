default: build

clean:
	rm -rf web
	rm -rf satisapi-go
	rm -rf test-config.json
	rm -rf satis
	rm -rf conposer.phar

deps:
	go get
	go get github.com/gorilla/http

satis:
	curl -sS https://getcomposer.org/installer | php
	php ./composer.phar create-project composer/satis lib/satis --stability=dev --keep-vcs

