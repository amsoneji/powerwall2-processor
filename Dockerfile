FROM golang:1.11.4

ARG app_name

WORKDIR /go/src/app
COPY ./common/ .
COPY ./${app_name}/ .

RUN go get -d -t -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["app"]