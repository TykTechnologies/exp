FROM ubuntu:16.04

USER root

RUN apt-get update && apt -yyy install curl
RUN curl -sL "https://packagecloud.io/tyk/tyk-gateway-unstable/packages/ubuntu/focal/tyk-gateway_5.3.1~alpha2_amd64.deb/download.deb?distro_version_id=210" -o ./tyk-gateway.deb
RUN dpkg -i ./tyk-gateway.deb

RUN /opt/tyk-gateway/tyk version
