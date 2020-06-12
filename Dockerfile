FROM composer/satis

RUN ln -s /satis/bin/satis /usr/local/bin/satis

ADD ./satis-go /opt/satis-go/satis-go
ADD https://github.com/benschw/satis-admin/releases/download/0.1.1/admin-ui.tar.gz /opt/satis-go/
RUN cd /opt/satis-go && tar xzvf ./admin-ui.tar.gz

ADD ./config-docker.yaml /opt/satis-go/config.yaml

ADD ./composer-config.json /composer/config.json
ADD ./composer-auth.json /composer/auth.json

EXPOSE 80

ENTRYPOINT sed -i "s/GITHUB_API_KEY/${GITHUB_API_KEY}/g" /composer/auth.json &&  /opt/satis-go/satis-go
