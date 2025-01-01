ARG STAGE
ARG STAGE=devo

# base image
FROM node:19-alpine as builder

# set working directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY . /usr/src/app
RUN npm install --silent
RUN npm install react-scripts@1.1.1 -g --silent

# build app
RUN npm run build

FROM nginx:alpine as prod

COPY --from=builder usr/src/app/dist /usr/share/nginx/html

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
