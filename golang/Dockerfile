#Main working space
FROM golang

#install dependency
USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim

#install goctl
USER root
RUN go install github.com/zeromicro/go-zero/tools/goctl@latest

#install protoc
USER root
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
        unzip -d /tmp/protoc /tmp/protoc.zip && \
        mv /tmp/protoc/bin/protoc $GOPATH/bin

#install protoc-gen-go
USER root
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Add GOATH to environment variable
ENV PATH $GOPATH/bin:$PATH

#clear
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

#Set working directory
WORKDIR /usr/src/code

EXPOSE 8000
EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003