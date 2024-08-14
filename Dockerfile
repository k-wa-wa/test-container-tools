FROM ubuntu:20.04

RUN apt-get -y update && \
  apt-get -y upgrade && \
  apt-get -y install iputils-ping net-tools curl traceroute

CMD ["tail", "-f", "/dev/null"]
