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
                  -----------------------------------|
- ****to give name to the docker container as node-app and run this container from the image we just build named node-app-example.****

      'docker run -d --name node-app-container node-app-example'
      // -d to run it in detach mode so that our terminal/cli is free it does not binds our cli.

> #### IMPORTANT
> ## How the Docker Interaction with the outside World(Internet/localmachine) Happens ?

- ****Docker can by default talk to the outside world,external machines and the Internet however the external machine or Internet cannot interact with docker****

- ****When we mention EXPOSE 5000 in the Dockerfile it does not mean that we are exposing the docker container at 5000 it is simply for documentation purposes and will not effect the docker interaction to the outer world in any way****

- ****In order to allow our Localhost machine/outer world machine to interact with docker we need to poke a hole in our local machine****

> #### PORT mapping

****-p maps the port so that the docker and externalmachine/localhostmachine/internet can talk to the container****

****the port number 5000 is the port at which the traffic is allowed to enter/interact with the docker container / say u have a server and u have set the PORT to 7858 and this server is running via docker container then the port number to the right of : will be 7858****

****the port number to the left i.e 3000 refers to the port from which the traffic will be coming from outside world****

          // sends the traffic from outside world coming form 3000 and sending it to port 5000 of the docker container

          docker run -p 3000:5000 -d --name node-app-container node-app-example

          go to localhost:3000 to see your app now it is running inside container.

> ## Accessing the docker container while it is running.
-****to access the docker with interactive cli bash****

          docker exec -it node-app-container(container-name) bash

          // now we will be bu default inside the app directory inside the running docker container
          root@e63797c4535c:/app#

          commands u can use
          ls
          // will show all the files inside the docker container that got copied from our localmachine

          docker rm node-app-container -f

> ## Rebuild Image

          docker build -t node-app-example .

> ## How to restrict all files from getting into the docker container like a case .env file with our credentials

- ****same as .gitignore we have .dockerignore that ignore the files to get copied over to the docker container like node_modules,Dockerfile,.env etc...****


          // .dockerignore


> ### IMPORTANT
> ## Volume mapping & Synching to Get real time update in our server/app code to avoid  the need of rebuilding image and then deploying it whenever somthing changes in our code.


- ****Bind Mount in docker helps to sync folder or file system in our local machine to the folder or file in docker****

         -v pathtofolderonLocalMachine:pathtofolderonContainer


         // below -v flag will sync the current working directory first_example mentioned as %cd% in windows cmd if in windows powershell,mac,linux then ${pwd}  in our local machine to the /app of the container

          docker run -v %cd% :/app -p 3000:5000 -d --name node-app-container node-app-example

> ### NOTE- that the node server has to be restarted to see the changes if using node server.js then it will still not show the updates in the browser as the express server has to be restarted again or just use NODEMON also use nodemon -L server.js with L flag to tackle issues of immediate exit of container in windows.
