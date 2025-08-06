FROM golang:1.24-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/dashboard cmd/app/main.go

FROM alpine:3.22

COPY --from=build /bin/dashboard /bin/dashboard
COPY --from=build /app/web ./web
RUN apk add --no-cache ca-certificates

CMD [ "/bin/dashboard" ]
