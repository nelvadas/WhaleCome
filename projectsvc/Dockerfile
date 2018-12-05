############################
# STEP 1 build executable binary
############################

FROM golang:onbuild

ADD  . $GOPATH/src/app/

RUN  go build -o /go/bin/app

EXPOSE 5000

ENTRYPOINT ["/go/bin/app"]
