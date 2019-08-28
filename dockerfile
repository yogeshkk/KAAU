FROM scratch


COPY kaau /opt/
COPY web /opt/
EXPOSE 3333


ENTRYPOINT ["/opt/kaau"]