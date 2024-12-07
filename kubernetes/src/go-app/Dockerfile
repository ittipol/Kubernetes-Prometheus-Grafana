# FROM golang:1.23.0-bookworm AS build
FROM golang:1.23.0 AS build

WORKDIR /go-app

COPY go.mod go.sum ./
#RUN go mod download
RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./app

# =================================================
# Distroless Container Image
# FROM gcr.io/distroless/base-debian12 AS production
FROM golang:1.23.0 AS production

# RUN apt-get -y update; apt-get -y install curl

WORKDIR /go-app

# RUN export GO_ENV_MODE=production

COPY --from=build /go-app/app ./
COPY ./config.yml ./config.yml

#ENV APP_PORT=8009

EXPOSE 8009 8010

ENTRYPOINT ["./app"]

# ENTRYPOINT ["tail", "-F", "/dev/null"]