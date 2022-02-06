FROM golang:1.11-alpine AS build
WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/snippetbox

FROM scratch
COPY --from=build /bin/snippetbox /bin/snippetbox
ENTRYPOINT ["/bin/snippetbox"]
