#!/bin/sh

curl -sS https://getcomposer.org/installer | php
php ./composer.phar create-project composer/satis --stability=dev --keep-vcs

