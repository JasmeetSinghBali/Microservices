> ## DevOps With Nodejs & Express

timestamp 20
> ## AIM: Setting Up workflow for developing node&express app within a docker container rather than in our local machine environment.

> ### Prerequisite
- [x] javascript
- [x] Docker
- [x] Node.js & Express
- [x] Knowledge about web apps/REST
- [x] Databases nosql/sql

> ## Level : Intermediate

> ## Topics Covered

- [x] Introduction to Docker & Node
- [x] Working with multiple containers
- [x] Moving to Production


> ### Introduction to Docker & Node
> #### Example -1 (first_example)

> #### Steps
- [x] Create simple express app.
- [x] Install docker on your machine.
- [x] pull node image from dockerhub this node image act as base image for our custom docker image.
- [x] to set up our development flow in docker we need to make sure that all the dependencies are in our custom docker image that we will built.


> #### Detailed Explanation

- ****To create a custom Image we need to make a 'Dockerfile'.****

- ****Dockerfile is a set of instructions/cmd that run to create our custom image.****


- ****Docker Image arranges the Dockerfile commands/instructions as layers see below.****

- ****Docker layers are build on one another and the docker caches the result of each layer i.e like say when we run docker build then it will run the cmd FROM and cache it result , then layer-2 WORKDIR and caches that result and so on....****

> #### IMPORTANT Use of the Docker layers and caching!

- ****say we want to rebuild our image with set of new layers and instruction now since the docker has already cached the result of the already build layers it will directly use that result instead of rebuilding from layer 1 and will only build the new layer we just added to our docker image.****

- ****even if we are changing a layer that was already build and whose result is cached then docker is going to rebuild from that layer to the end of the layer as that change in layer may effect the layer below it****

- ****now if we specify in layer 3 to copy package.json and in layer 5 to copy all the files then if anything changes in our src code only then docker will only rebuild the layer 5 .****

                   ----------------------------------
                  |     DOCKER   Image               |
                  |                                  |
                  |      Layer-1 FROM                |
                  |      Layer-2 WORKDIR             |
                  |      Layer-3 COPY package.json . |
                  |      Layer-4 RUN                 |
                  |      Layer-5 COPY . .            |
                  |                                  |
                  |                                  |
                  |                                  |
                  |                                  |
                  ---------------------------------  |
- ****to give name to the docker container as node-app and run this container from the image we just build named node-app-example.****

      'docker run -d --name node-app-container node-app-example'
      // -d to run it in detach mode so that our terminal/cli is free it does not binds our cli.
