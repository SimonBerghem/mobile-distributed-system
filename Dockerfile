FROM golang:latest

RUN echo "hello world"

WORKDIR $GOPATH/labCode/github.com/mobile-distributed-system
RUN apt update
COPY go.mod .
RUN go mod download
COPY . .
EXPOSE 80/udp
RUN go run src/main.go
CMD  ["./main"]

# Add the commands needed to put your compiled go binary in the container and
# run it when the container starts.
#
# See https://docs.docker.com/engine/reference/builder/ for a reference of all
# the commands you can use in this file.
#
# In order to use this file together with the docker-compose.yml file in the
# same directory, you need to ensure the image you build gets the name
# "kadlab", which you do by using the following command:
#
# $ docker build . -t kadlab
