#!/bin/bash

set -e

mkdir -p artifacts
sudo docker run -i --rm -v `pwd`:/mnt/src/github.com/coopernurse/iris golang /bin/bash <<EOF
set -e

export GOPATH=/mnt
export CGO_ENABLED=0
cd /mnt/src/github.com/coopernurse/iris

echo "Fetching dependencies"
go get google.golang.org/api/compute/v1
go get code.google.com/p/goauth2/compute/serviceaccount
go get code.google.com/p/go.crypto/hkdf
go get gopkg.in/inconshreveable/log15.v2

echo "Building iris"
go build -a -ldflags '-s' -o artifacts/iris main.go

echo Changing owner from \$(id -u):\$(id -g) to $(id -u):$(id -u)
chown -R $(id -u):$(id -u) /mnt/src/github.com/coopernurse/iris/artifacts
EOF

sudo docker build -t coopernurse/iris .
