FROM golang:1.23 AS build

WORKDIR /src
COPY main.go go.mod /src/
RUN go mod download github.com/lib/pq && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app /src/main.go

FROM scratch
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
