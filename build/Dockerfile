FROM golang:1.24-alpine AS build

WORKDIR /build
ENV GOOS=linux \
    GOARCH=arm64 \
    CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN apk add make
RUN make build

FROM alpine:latest

COPY --from=build /build/bin /build/config.env ./

CMD [ "./task-exporter" ]