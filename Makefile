default: build

clean:
	rm -rf satisapi-go
	rm -rf test-config.json
	rm -rf satis
	rm -rf conposer.phar

deps:
	go get

satis:
	curl -sS https://getcomposer.org/installer | php
	php ./composer.phar create-project composer/satis --stability=dev --keep-vcs

