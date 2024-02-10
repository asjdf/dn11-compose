FROM ubuntu:latest

RUN sed -i 's@//.*archive.ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list && \
    sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

RUN apt update && \
    apt install -y curl cron iptables wireguard-tools iproute2 mtr inetutils-ping dnsutils tcpdump iperf3 bird2 && \
    mkdir -p /run/bird/

COPY ./wg-reresolve-dns.sh /usr/local/bin/wg-reresolve-dns.sh
COPY ./crontab.txt /crontab.txt

RUN crontab < /crontab.txt

COPY ./enterpoint.sh /enterpoint.sh

RUN chmod +x /enterpoint.sh

CMD ["/enterpoint.sh"]
