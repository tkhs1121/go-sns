FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download
EXPOSE 1129

RUN go get -u github.com/gofiber/fiber/v2
RUN go get -u gorm.io/gorm
RUN go get github.com/go-ozzo/ozzo-validation/v4
RUN go get github.com/go-ozzo/ozzo-validation/v4/is
RUN go get github.com/joho/godotenv

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]