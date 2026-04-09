FROM golang:1.21 AS build

WORKDIR /src
COPY main.go /src/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/hello /src/main.go

FROM scratch
COPY --from=build /out/hello /hello
ENTRYPOINT ["/hello"]
