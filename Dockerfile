FROM golang:1.19.0-alpine as build

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN apk update
RUN apk add git

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN

RUN git config --global url."https://ciricc:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
# ENV GOPRIVATE github.com/ciricc

RUN go build -o printer cmd/printer/main.go

FROM alpine:latest as production

RUN apk add postgresql-client curl
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

COPY --from=build /app /

# Порт запускаемого сервиса
EXPOSE 8080
CMD ["./printer"]