### STAGE 1: Build ###

# We label our stage as 'builder'
FROM node:14.21.1-alpine as builder

# previous: 8192
RUN export NODE_OPTIONS=--max_old_space_size=6144

WORKDIR /ng-app

COPY package*.json ./

RUN npm i -f

COPY . .

RUN npm run build:prod

### STAGE 2: Setup ###
FROM nginx:1.21

ENV KLOVERCLOUD_API_ENDPOINT=''

## Copy our default nginx config
COPY nginx/default.conf /etc/nginx/conf.d/
COPY nginx/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

## From 'builder' stage copy over the artifacts in dist folder to default nginx public folder
COPY --from=builder /ng-app/dist/kc /usr/share/nginx/html

RUN mkdir /app
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh
RUN sed -i -e 's/\r$//' /app/run.sh

EXPOSE 8000

ENTRYPOINT ["bin/sh", "/app/run.sh"]
