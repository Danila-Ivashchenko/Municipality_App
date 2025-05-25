FROM golang:1.23.2 as builder
WORKDIR /build
COPY go.mod . 
COPY go.sum . 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine
COPY --from=builder main /bin/main

COPY ./migrations ./migrations
COPY ./font ./font
ENTRYPOINT ["/bin/main"]