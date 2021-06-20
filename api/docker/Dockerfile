FROM golang:alpine3.13 AS build

WORKDIR /usr/src/project
COPY . .
# build go app
RUN go build -o bin/app

FROM alpine:3.13
WORKDIR /usr/bin/project

COPY --from=build /usr/src/project/bin/app ./
COPY --from=build /usr/src/project/docker/wait-for-postgres.sh ./
COPY --from=build /usr/src/project/.env ./

# install psql
RUN apk add --no-cache postgresql-client
# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

EXPOSE 8000
CMD [ "./app" ]