FROM golang:alpine as buildenv
WORKDIR /go/src/github.com/rkachowski/ps1.fun
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /ps1.fun ./cmd/ps1-fun

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=buildenv /ps1.fun /
ENTRYPOINT ["/ps1.fun"]