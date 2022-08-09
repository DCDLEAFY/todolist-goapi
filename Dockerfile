FROM golang:alpine3.16

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY controllers/todo_controller.go controllers/
COPY main.go .

RUN go build -o todo-api

CMD [ "./todo-api" ]

EXPOSE 8080