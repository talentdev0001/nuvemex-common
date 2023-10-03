FROM golang:1.16

RUN apt-get update && apt-get install -y git

WORKDIR /goseanto

RUN GOCACHE=OFF

RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >>/root/.gitconfig
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " >/root/.ssh/config
RUN go env -w GOPRIVATE=github.com/nuvemex

COPY . .

CMD ["/bin/bash"]
