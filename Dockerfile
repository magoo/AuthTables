FROM golang:1.7

# Create a workspace
RUN mkdir -p /opt/authtables
WORKDIR /opt/authtables

# install deps
RUN go get github.com/willf/bloom \
           gopkg.in/redis.v4 \
           github.com/Sirupsen/logrus

# Add our files
ADD authtables.go authtables.go
ADD .env .env
ADD configuration.go configuration.go
ADD datastore.go datastore.go

# Build app
RUN go build authtables.go configuration.go datastore.go

# Default runs on 8080
EXPOSE 8080

# Run our binary
CMD /opt/authtables/authtables
