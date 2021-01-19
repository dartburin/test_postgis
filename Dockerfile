FROM golang:latest

ARG A_DB_USER=postgres
ARG A_DB_PASS=postgres
ARG A_DB_BASE=coords
ARG A_DB_HOST=servgisdb
ARG A_DB_PORT=5432
ARG A_HTTP_PORT=8080
ARG A_HTTP_HOST=0.0.0.0

ENV APP_NAME serverGRPC

ENV DB_USER=${A_DB_USER} 
ENV DB_PASS=${A_DB_PASS} 
ENV DB_BASE=${A_DB_BASE} 
ENV DB_HOST=${A_DB_HOST} 
ENV DB_PORT=${A_DB_PORT} 
ENV HTTP_HOST=${A_HTTP_HOST} 
ENV HTTP_PORT=${A_HTTP_PORT} 

#RUN go version

RUN mkdir -p ${GOPATH}/src/${APP_NAME}
WORKDIR ${GOPATH}/src/${APP_NAME}

COPY go.mod ./
COPY go.sum ./

COPY ./cmd ./
COPY ./cmd/server ./cmd/
COPY ./cmd/server/* ./cmd/server/

COPY ./internal/ ./
COPY ./internal/api ./internal/
COPY ./internal/api/server ./internal/api/
COPY ./internal/api/server/* ./internal/api/server/
COPY ./internal/api/middleware ./internal/api/
COPY ./internal/api/middleware/* ./internal/api/middleware/
COPY ./internal/gis ./internal/
COPY ./internal/gis/* ./internal/gis/
COPY ./internal/logger ./internal/
COPY ./internal/logger/* ./internal/logger/

#RUN ls -l
#RUN pwd

RUN go mod download 
RUN go build cmd/server/main.go
RUN ls -l

EXPOSE ${HTTP_PORT}
EXPOSE ${DB_PORT}

CMD ./main -dbhost=${DB_HOST} -dbbase=${DB_BASE} -dbuser=${DB_USER} -dbpass=${DB_PASS} -dbport=${DB_PORT} -httpport=${HTTP_PORT} -httphost=${HTTP_HOST}

