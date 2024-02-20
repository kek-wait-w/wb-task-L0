FROM golang:1.21.1-alpine AS build-stage

WORKDIR /post_data

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /post_data/post_data

RUN go build -o post_data

FROM gcr.io/distroless/base-debian11 AS post_data_stage

WORKDIR /

COPY --from=build-stage /post_data/post_data /

ENTRYPOINT ["/post_data"]
