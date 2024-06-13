FROM ubuntu:latest

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y vim net-tools nodejs npm openssh-server supervisor sudo && \
    groupadd sshgroup && useradd -ms /bin/bash -g sshgroup user && \
    mkdir -p /home/user/.ssh


ADD ayush.pub.pem /home/user/.ssh/authorized_keys
ADD supervisord.conf /etc/supervisord.conf

RUN chown user:sshgroup /home/user/.ssh/authorized_keys && \
    chmod 600 /home/user/.ssh/authorized_keys && \
    service ssh start && \
    mkdir -p /var/log/supervisor


WORKDIR /home/user
COPY . .
RUN npm i

CMD ["/usr/bin/supervisord", "--configuration=/etc/supervisord.conf"]

EXPOSE 22