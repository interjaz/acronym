FROM golang:1.6-alpine

# Copy source
ADD ./Server/src/ /go/src/

# Install
RUN apk add --update alpine-sdk

RUN go install AcronymServerFetcher
RUN go install AcronymServerWebsite
RUN mkdir /go/web
RUN cp /go/src/AcronymServerWebsite/web /go/web


CMD /go/bin/AcronymServerFetcher & /go/bin/AcronymServerWebsite && fg
EXPOSE 8080

