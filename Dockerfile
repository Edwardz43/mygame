# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && ls -l
RUN cd /src/app && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
RUN touch ./app
ENTRYPOINT ./app