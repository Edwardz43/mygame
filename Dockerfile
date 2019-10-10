# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && ls -l
RUN cd /src/gameserver && go build -o gameserver

# final stage
FROM alpine
WORKDIR /gameserver
COPY --from=build-env /src/gameserver /gameserver/
RUN touch ./gameserver
ENTRYPOINT ./gameserver