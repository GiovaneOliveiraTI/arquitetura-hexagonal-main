FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 \
    \
# Instala el compilador de C y las bibliotecas necesarias
RUN apt-get update && apt-get install -y gcc libc6-dev

RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]
