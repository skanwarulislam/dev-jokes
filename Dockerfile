# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /srv/makeajoke

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN make

##
## Deploy
##
FROM gcr.io/distroless/base-debian10
#
WORKDIR /app/
#
COPY --from=build /srv/makeajoke/app/  ./

EXPOSE 3000

USER nonroot:nonroot

CMD ["/app/server"]