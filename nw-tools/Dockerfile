FROM ubuntu:22.04

RUN apt-get -y update && \
  apt-get -y upgrade && \
  apt-get -y install \
    iputils-ping net-tools curl traceroute bind9-dnsutils iptables iptables-persistent \
    vim

CMD ["tail", "-f", "/dev/null"]
