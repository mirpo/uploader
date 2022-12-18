FROM golang:1.19
RUN apt-get update -qq
RUN apt-get install -y \
  libtesseract-dev \
  tesseract-ocr-eng

ENV GOPATH=/root/go
ENV GO111MODULE=on

WORKDIR /build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server .
EXPOSE 3333
CMD ["./server"]
