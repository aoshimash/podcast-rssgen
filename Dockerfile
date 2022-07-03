##
## Build
##
FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o podcast-rssgen ./main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/podcast-rssgen .

USER nonroot:nonroot

ENTRYPOINT [ "/podcast-rssgen" ]
