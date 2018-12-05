############################
# STEP 1 build executable binary
############################
FROM golang:onbuild
WORKDIR .
RUN  go build -o /go/bin/app
EXPOSE 5000
ENTRYPOINT ["/go/bin/app"]
