FROM golang:1.13.8-alpine3.11 as build
WORKDIR /src/
COPY container.go . 
RUN export GO111MODULE=on
RUN apk update && apk add git && go get github.com/gin-gonic/gin
#RUN go build -o ./demo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo .


FROM alpine:3.11.3
WORKDIR /root/
COPY --from=0 /src/demo ./

EXPOSE 8080
ENTRYPOINT ["./demo"]
