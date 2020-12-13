FROM golang:1.15.1-alpine3.12

ADD . /server
WORKDIR /server

ENV GOPROXY https://goproxy.cn
RUN go mod tidy
RUN ln -sfv /server/conf/conf.ini.pro /server/conf/conf.ini
EXPOSE 9000

CMD sh run.sh
