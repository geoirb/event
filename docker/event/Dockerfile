FROM golang:1.16-alpine AS build

WORKDIR /go/src/event
COPY . ./

RUN go build \
    -o /out/service \
    ./cmd/event/main.go

FROM alpine
WORKDIR /app
COPY --from=build /out/service /app/service
CMD ["/app/service"]