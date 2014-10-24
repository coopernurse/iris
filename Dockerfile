FROM debian:wheezy
RUN apt-get update
RUN apt-get -y install git mercurial curl

WORKDIR /usr/local
RUN curl https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz -o go1.3.3.linux-amd64.tar.gz
RUN tar zxf go1.3.3.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /usr/local

RUN /usr/local/go/bin/go get github.com/project-iris/iris
WORKDIR /usr/local/src/github.com/coopernurse/iris
RUN /usr/local/go/bin/go build -o /usr/local/bin/iris main.go

ADD iris-test /usr/local/etc/iris-test

EXPOSE 55555
CMD /usr/local/bin/iris -net test -rsa /usr/local/etc/iris-test
