FROM golang:1.15-alpine AS build

COPY . /app
WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
RUN go build -o /bin/communed

# This results in a single layer image
FROM alpine:3.12
COPY --from=build /bin/communed /bin/communed
ENTRYPOINT [ "/bin/communed" ]