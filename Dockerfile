FROM composer/satis

RUN ln -s /satis/bin/satis /usr/local/bin/satis

ADD ./satis-go /opt/satis-go/satis-go
ADD https://github.com/benschw/satis-admin/releases/download/0.1.1/admin-ui.tar.gz /opt/satis-go

ADD ./config-docker.yaml /opt/satis-go/config.yaml

EXPOSE 80

ENTRYPOINT /opt/satis-go/satis-go
