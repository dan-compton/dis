FROM quay.io/opsee/build-go:latest

MAINTAINER Dan Compton <dan@opsee.co>

RUN apk add --update \
        curl \
    && rm -rf /var/cache/apk/*

RUN apk-install iptables ca-certificates lxc e2fsprogs 
RUN apk-install docker
RUN apk-install python

EXPOSE 9009

ADD ./wrapdocker /usr/local/bin/wrapdocker
COPY /bin/dis /dis

RUN chmod +x /usr/local/bin/wrapdocker

ENV PATH=$PATH:/usr/local/bin
VOLUME /var/lib/docker

CMD ["wrapdocker"]
