# ============= A single container setup for Dockerfile ====================
FROM node:carbon
# Base Image

ENV secretToken=SECRET
# Env variables for docker

WORKDIR /src
# set working dir for docker container

COPY . /src
ADD ./hello.tar /root
ADD http://something.com /root
# adding copying project file/folder to the docker /root directory

# installing dependancies
RUN cd /src;npm install

# Exposing the port on which the docker container will run
EXPOSE 3000

# Default commands CMD or ENTRYPOINT that will be executed automaically in shell/terminal when the container is started up
CMD ["npm","start"]
ENTRYPOINT ls


# ================ A  multiple container setup for Dockerfile======================
FROM node:8.16 as build-deps
WORKDIR /usr/src/app

# as we have already specified working directory we dont need to specify it every time by using ./ it refers to the /usr/src/app automatically
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROm nginx:1.12-alpine
COPY --from=build-deps /usr/src/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx","-g","daemon off;"]
