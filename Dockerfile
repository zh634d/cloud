FROM ubuntu
RUN mkdir /httpserver
WORKDIR /httpserver
ENV MY_SERVICE_PORT=80
ENV MY_SERVICE_PORT1=80
ENV MY_SERVICE_PORT2=80
ENV MY_SERVICE_PORT3=80
LABEL multi.label1="BeiJing" multi.label2="cunshang" other="value3"
ADD ./httpserver /httpserver/
EXPOSE 80
ENTRYPOINT /httpserver/httpserver
