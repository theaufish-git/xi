FROM mariadb:10.9.6

RUN apt update
RUN apt install -y git libmariadb-dev netcat python3-pip
RUN apt clean

RUN mkdir -p /opt/xi/tools
COPY tools/requirements.txt /opt/xi/tools/requirements.txt
RUN pip3 install -r /opt/xi/tools/requirements.txt

COPY .git /opt/xi/.git
COPY sql /opt/xi/sql
COPY modules/init.txt /opt/xi/modules/init.txt
COPY tools/dbtool.py /opt/xi/tools/dbtool.py
COPY tools/migrations /opt/xi/tools/migrations
COPY _ops/sql/schema /docker-entrypoint-initdb.d

WORKDIR /opt/xi
