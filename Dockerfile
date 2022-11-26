FROM golang:alpine
ARG PORT=3000
ENV CGO_ENABLED=0
ENV PORT=$PORT

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE $PORT

CMD [ "./main" ]