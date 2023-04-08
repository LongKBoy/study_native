FROM ubuntu
ENV MY_SERVICE_PORT=80
ENV TINI_VERSION v0.19.0
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--","/httpserver"]