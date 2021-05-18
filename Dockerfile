 
FROM golang:1.16

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]




# FROM golang:1.16

# COPY . /go/src/app

# WORKDIR /go/src/app/cmd/callback

# RUN go build -o callback main.go

# EXPOSE 9090

# CMD ["./callback"]