FROM golang:1.19-alpine as builder

ENV BASE_DIR /go/src/github.com/powerslider/flight-tracker
WORKDIR ${BASE_DIR}

RUN apk --no-cache add git ca-certificates

COPY go.mod go.sum ${BASE_DIR}/

RUN go mod download -x

COPY cmd ${BASE_DIR}/cmd
COPY docs ${BASE_DIR}/docs
COPY pkg ${BASE_DIR}/pkg

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /dist/flight-tracker ./cmd/flight-tracker/main.go

FROM alpine

ENV BASE_DIR /go/src/github.com/powerslider/flight-tracker

COPY --from=builder ${BASE_DIR}/docs /docs/
COPY --from=builder /dist .

CMD ["/flight-tracker"]
