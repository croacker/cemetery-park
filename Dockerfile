FROM golang:1.16-alpine AS development
EXPOSE 8080
RUN mkdir /app
RUN mkdir /app/db
COPY ./cemetery-park-back /app/
WORKDIR /app
ENTRYPOINT ["/app/cemetery-park-back"]
