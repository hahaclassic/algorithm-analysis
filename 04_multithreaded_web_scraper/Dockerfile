FROM ubuntu:22.04

RUN apt-get -yy update && apt-get install -yy \
    make cmake \
    curl && \
    rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://golang.org/dl/go1.23.3.linux-amd64.tar.gz -o go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

COPY . /app

WORKDIR /app

RUN make -f iu7-algan-cicd-student.mk ready/app-cli-debug
