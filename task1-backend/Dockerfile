FROM alpine as builder

ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:${PATH}

RUN mkdir -p /app
WORKDIR /app

RUN apk add --no-cache go git curl gcc g++ make

COPY go.sum .
COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o task1 github.com/mainnika/a-quest/task1-backend

FROM alpine

ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:${PATH}

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/task1 .
COPY --from=builder /app/task1-backend/config config

EXPOSE 8081

CMD ["/app/task1"]