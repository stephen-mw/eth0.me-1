FROM alpine:3.4

ENV HOME /usr/local/src/eth0.me
WORKDIR ${HOME}

ADD bin/eth0.me.linux.amd64 ${HOME}/run
ADD server.key ${HOME}/server.key
ADD server.pem ${HOME}/server.pem

EXPOSE 80 443

ENTRYPOINT ["./run"]
