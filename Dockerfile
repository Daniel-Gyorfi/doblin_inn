FROM node:18-alpine as dev

WORKDIR /app

# directly download Go
RUN wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz &&\
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz &&\
    rm go1.21.3.linux-amd64.tar.gz

# add Go to path
ENV PATH="${PATH}:/usr/local/go/bin:/root/go/bin"

# install Templ templating library
RUN go install github.com/a-h/templ/cmd/templ@latest &&\
    go install github.com/cosmtrek/air@latest

# install TailwindCSS
# FROM node:18-alpine as build-css

# WORKDIR /app
# COPY . .
# RUN npm ci && npx tailwindcss -i ./build/input.css -o /output.css


#final step in golang 
FROM golang:1.22-alpine as production

RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ./
# COPY --from=build-css /output.css ./static/output.css

RUN templ generate
RUN go build -o /main
CMD [ "/main" ]

EXPOSE 8080
