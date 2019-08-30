FROM scratch

COPY kaau /
COPY /web /web/

EXPOSE 3333

ENTRYPOINT ["./kaau"]
