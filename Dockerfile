FROM golang:latest
EXPOSE 2525
RUN mkdir /app
COPY ./SMTP ./app/
WORKDIR ./app

ENTRYPOINT ["./SMTP"]
