FROM golang:1.13.4-alpine3.10

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/bin

RUN apk update && \
    apk upgrade && \
    apk add vim && \
    apk add git && \
    apk add build-base

# Install our dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/joho/godotenv
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/google/uuid
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/ses
RUN go get github.com/bamzi/jobrunner
RUN go get github.com/k-washi/jwt-decode/jwtdecode
RUN go get firebase.google.com/go
RUN go get github.com/gin-contrib/cors

#RUN go get -u github.com/golang/protobuf/proto
#RUN go get -u google.golang.org/grpc
#RUN go get -u github.com/grpc-ecosystem/go-grpc-middleware
#RUN go get -u github.com/urfave/cli
#RUN go get -u google.golang.org/grpc/credentials
#RUN go get -u google.golang.org/grpc/reflection
#RUN go get -u github.com/xo/xo
#RUN go get -u github.com/jmoiron/sqlx
#RUN go get -u github.com/Masterminds/squirrel

RUN apk add tzdata
ENV TZ=Asia/Tokyo

# Expose default port (8080)
EXPOSE 8080
