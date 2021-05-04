> # Multiple Containers Docker Setup

- ****say we have an application built in python,node.js,sql,nosql and then building seperate images for each of them and then running containers becomes clumsy.****

> ## Container Communication

- ****the core idea is to create a network inside the  container such that each container can communicate with other containers.****


****Example pulling multiple images from docker hub****

       docker pull mysql
       docker pull redis
       docker pull nodejs

       docker run mysql redis nodejs
       // three sepearte containers will be running

> ## Managing Multiple Containers with docker compose

- ****docker compose****
- ****example node and mongo via docker-compose.yaml****

        // docker-compose.yml
        version: "2"
        services:
          app:
            container_name: app
            restart: always
            build: .
            ports:
              - "8080:8080"
            links:
              -mongo
          mongo:
            container_name: mongo
            image: mongo
            volumes:
              - ./data:/data/db
            ports:
              - "27017:27017"


          // Dockerfile
          FROM node:carbon


          # set working directory
          WORKDIR /src

          # Copy all the node.js project files from local working dir to container working dir /src
          COPY . .

          # Install App dependencies
          RUN npm install

          EXPOSE 8080

          CMD ["npm","start"]




-****docker-compose build will build the image from Dockerfile & docker-compose up will run mongo at 27017 and node at 8080****

         docker-compose build
         docker-compose up

> #### GUI for Docker managment kitematic
https://kitematic.com/
