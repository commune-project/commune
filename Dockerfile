FROM golang:1.15-alpine AS build

COPY . /app
WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
RUN go build -o /bin/commune

# This results in a single layer image
FROM alpine:3.12
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add postgresql-client
COPY --from=build /bin/commune /bin/commune
COPY migrations /migrations
CMD [ "/bin/commune" ]