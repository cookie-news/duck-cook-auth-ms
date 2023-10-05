FROM golang:alpine3.18 AS build-stage

WORKDIR /app/src/duck-cook-auth

ENV GOPATH=/app

COPY . .

RUN chmod +x /app/src/duck-cook-auth

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /duck-cook-auth

FROM alpine:latest AS build-release-stage

ENV TZ=America/Sao_Paulo

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app/src/duck-cook-auth

RUN chmod +x /app/src/duck-cook-auth

COPY --from=build-stage /duck-cook-auth /app/src/duck-cook-auth/duck-cook-auth

EXPOSE 8080

ENTRYPOINT ["/app/src/duck-cook-auth/duck-cook-auth"]