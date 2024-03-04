# RUN go install github.com/cosmtrek/air@latest


# run TailwindCSS stylesheet compilation
FROM node:18-alpine as build-css

WORKDIR /app
COPY . .
RUN npm ci 
RUN npx tailwindcss -i ./static/tailwind.css -o ./static/output.css

#final step in golang 
FROM golang:1.22-alpine as production

# install Templ templating library
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ./
COPY --from=build-css /app/static/output.css ./static/output.css

RUN templ generate
RUN go build -o /main
CMD [ "/main" ]

# set user to non-root

EXPOSE 8080
