FROM golang:1.18


WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY app.go .
COPY model.go .

RUN go mod download

ENV APP_HOST=0.0.0.0 \
    APP_PORT=5432 \
    APP_DB_USERNAME=postgres \
    APP_DB_PASSWORD=postgres \
    APP_DB_NAME=postgres

RUN go build -o /go-api

EXPOSE 8080

CMD ["/go-api"]

# http://localhost:8080/api/v10/minhacesta/rs/ijui/centro/arroz
# docker build -t go-api .
# docker run --name go-api -d -p 8080:8080 go-api