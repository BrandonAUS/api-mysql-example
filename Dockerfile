# Dockerfile

FROM golang:1.16


WORKDIR /
COPY . .
RUN go get -u github.com/gorilla/mux

CMD ["go","run","main.go"]