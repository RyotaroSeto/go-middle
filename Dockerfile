FROM golang:1.20.3-alpine

ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV PACKAGES="ca-certificates git curl bash zsh make"
ENV ROOT /app

RUN apk update \
  && apk add --no-cache ${PACKAGES} \
  && update-ca-certificates

WORKDIR ${ROOT}

RUN go install github.com/cosmtrek/air@latest

COPY ./ ./

WORKDIR ${ROOT}

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]
