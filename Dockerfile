FROM golang:1.24-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM alpine
WORKDIR /var/
COPY --from=build /app .
EXPOSE 8080
CMD ["./app"]