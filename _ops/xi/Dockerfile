FROM ubuntu:22.04 as image

RUN apt-get update
RUN apt-get install -y libluajit-5.1-dev libzmq3-dev libssl-dev zlib1g-dev mariadb-server libmariadb-dev binutils-dev
RUN apt clean

FROM image as buildimage

RUN apt-get update
RUN apt-get install -y git python3 python3-pip g++-10 cmake make wget curl

RUN wget https://r.mariadb.com/downloads/mariadb_repo_setup
RUN chmod +x mariadb_repo_setup

RUN ./mariadb_repo_setup --mariadb-server-version="mariadb-10.6"

RUN apt-get -y install libmariadb3 libmariadb-dev

FROM buildimage as build

COPY cmake /opt/xi/cmake
ADD CMakeLists.txt /opt/xi/CMakeLists.txt
ADD CMakeSettings.json /opt/xi/CMakeSettings.json
ADD .clang-format /opt/xi/.clang-format
ADD .clang-tidy /opt/xi/.clang-tidy

COPY ext /opt/xi/ext
COPY modules /opt/xi/modules
COPY res /opt/xi/res
COPY src /opt/xi/src
COPY tools /opt/xi/tools

RUN mkdir /opt/xi/out
WORKDIR /opt/xi/out
RUN cmake /opt/xi
RUN make

FROM image

COPY --from=build /opt/xi/xi_connect /opt/xi/bin/xi_connect
COPY --from=build /opt/xi/xi_map /opt/xi/bin/xi_map
COPY --from=build /opt/xi/xi_search /opt/xi/bin/xi_search
COPY --from=build /opt/xi/xi_world /opt/xi/bin/xi_world
COPY --from=build /opt/xi/res /opt/xi/res

WORKDIR /opt/xi
