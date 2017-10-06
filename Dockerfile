FROM rawmind/alpine-base:3.6-7
MAINTAINER Raul Sanchez <rawmind@gmail.com>


ENV WEB_HOME=/opt/web-monolith \
    VERSION=0.1 \
    GOMAXPROCS=2 \
    GOROOT=/usr/lib/go \
    GOPATH=/opt/src \
    GOBIN=/gopath/bin
ENV PATH=$PATH:$WEB_HOME

# Add service files
ADD root /

RUN apk add --update go git \ 
  && mkdir -p /opt/src $WEB_HOME; cd /opt/src \
  && cd /opt/src \
  && CGO_ENABLED=0 go build -v -installsuffix cgo -ldflags '-extld ld -extldflags -static' -a -x web-monolith.go \
  && mv ./web-test ${WEB_HOME}; cd ${WEB_HOME} \
  && chmod +x ${WEB_HOME}/web-monolith \
  && apk del go git \
  && rm -rf /var/cache/apk/* /opt/src 

EXPOSE 8080

ENTRYPOINT ["/opt/web-test/web-monolith"]
