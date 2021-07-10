> ## DevOps With Nodejs + Express + MONGO(CRUD) + REDIS(Auth)
timestamp 350:35 Load balancing with nginx proxy container

> ## AIM: Setting up workflow for developing node&express app within a docker container and understanding how production works with containerize applications i.e to deploy as different microservices.



> # Blueprint
****NODE POST API + MONGO(CRUD) + REDIS(AUTH)****
- [x] Docker setup for Prod & Dev environment
- [x] CRUD with mongoDB & Login
- [x] Auth with express-session & connect-redis https://www.npmjs.com/package/connect-redis

****Note-1 - Each time you add new dependency or change environment variable do docker-compose down then up --build,also when changing in the docker-compose file then rerun the docker-compose up it will automatically detect changes.****
****Note-2 - Remember we can also do docker up directly when we add new dependency but it will have to make sure to pass --build and -V flag to create a new anonymous volume so that the old information from the container volume is not retrieved. example docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build -V****

> ## Topics Covered

- [x] Introduction to Docker & Node
- [x] Working with multiple containers
- [x] Local Docker setup for developement
- [x] Docker Setup for Moving to Production


> ### Prerequisite
- [x] javascript
- [x] Docker
- [x] Node.js & Express
- [x] Knowledge about web apps/REST
- [x] Databases nosql/sql



=======================================
> ### Introduction to Docker & Node
> #### Example -1 (first_example)

=========================================
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

====================================================
> #### IMPORTANT Use of the Docker layers and caching!

====================================================

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

=========================================
> #### IMPORTANT
> ## How the Docker Interaction with the outside World(Internet/localmachine) Happens ?

==========================================

- ****Docker can by default talk to the outside world,external machines and the Internet however the external machine or Internet cannot interact with docker****

- ****When we mention EXPOSE 5000 in the Dockerfile it does not mean that we are exposing the docker container at 5000 it is simply for documentation purposes and will not effect the docker interaction to the outer world in any way****

- ****In order to allow our Localhost machine/outer world machine to interact with docker we need to poke a hole in our local machine****

=======================

> ## PORT mapping

========================

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

****to run the existing container again****

         docker run -v %cd%:/app -p 3000:5000 -d node-app-container:latest


***

> ## Important CASE Preventing Bind Mount Sync For node_modules changes in local and docker volume.

 ****say that we delete node_modules in localsystem and if a bindmount is set up for the docker then the node_modules gets also deleted in docker.****

****To tackle this problem and prevent the delete sync of node_modules in docker when we delete node_modules in local system we use workaround****

         docker run -v %cd%:/app -v /app/node_modules -p 3000:5000 -d --name node-app-container node-app-example

         // the above extra -v will prevent bind mount to syn the node_modules changes.

****NOTE-> the COPY command is still nedded though we are using bind mount only for development in production we dont have bind mount so for that COPY is necessary in Dockerfile.****

         docker exec -it node-app-container bash
         touch testfile
         // creates a test file both in docker and our local system project folder.

***

> ## IMPORTANT CASE to restrict the docker volume from changing the src code or add/delete files in our local file system.

****READ ONLY BIND MOUNT i.e the docker can read the changes i.e sync changes from the local file system but the changes in docker volume do not gets reflected back in the local file system.****

          docker run -v %cd%:/app:ro -v /app/node_modules -p 3000:5000 -d --name node-app-container node-app-example

          // :ro specifies read only bind mount

          docker exec-it node-app-container bash
          touch newfile
          error:read only file system.

==================================

> ## Docker and Env Variables

==================================

****Inside Dockerfile****

          // creates a ENV variable with key value as PORT 3000
          ENV PORT 3000
          EXPOSE $PORT

          // rebuild the Image after change to docker file
          docker build -t node-app-example .
          docker run -v %cd%:/app:ro -v /app/node_modules --env PORT=4000 -p 3000:4000 -d --name node-app-container node-app-example
          // will set the PORT to 4000 rather than default 5000.

          docker exec-it node-app-container bash
          printenv
          // to see the env variable PORT of the docker exposed that we set inside DOCKER container.

          // FOR say 20 or so environment variables we can use .env in the root.
          PORT=4000

          // passing the envirnment file to docker // while docker container run command

          docker run -v %cd%:/app:ro -v /app/node_modules --env-file ./.env -p 3000:4000 -d --name node-app-container node-app-example

****to remove the associatd volume to the container u want to remove****
          docker volume ls

          // to delete a certain runnning container & its volume with -v flag for voleme and -f for force to remove the running container

          docker rm node-app-container -fv  

          // to remove unused volumes i.e volume not used by any container

          docker volume prune
=======================================

> # DOCKER COMPOSE

=======================================

> ## DOCKER COMPOSE handling multiple containers and microservices running in those different containers help to automate the docker build,run,stop commands.

****In your root folder make docker-compose.yml****
****refer: https://docs.docker.com/compose/****

> #### Important Facts docker compose

- [x] each docker container that will run act as services.
- [x] with yml files spacing matters.

           // to build the image and run the container
           // for first times
           docker-compose up -d

           // to stop the container , note u must
           // specify anonymous volumes it should be
           // mentioned via -v explicitely
           docker-compose down -v


           // subsequent future it directly runs the ///container
           docker-compose up -d

- [x] IMPORTANT- docker compose is pretty dumb as if we change the Dockerfile content say we made change to the port and run docker-compose up -d it will still run the old image that has the old port.

- [x] note that docker-compose just looks for the already existing image means it is kinda of lazy and do not bothers to crosscheck wheather the initial Dockerfile has undergone any changes.

==================================================
> ### IMPORTANT So to tackle the problem of docker-compose running stale images even when we had made changes in the original Dockerfile from which image was built.

==================================================

****'--build' flag forces the docker-compose to rebuild the image****

        docker-compose up -d --build

======================================

IMPORTANT
> # Seperated Docker Set up for Production and Development environments.

****A Single Dockerfile and dockercompose.dev.yml/docker-compose.prod.yml files for productiona and dev environment configurations****

======================================

****Note the orders matters first base file then dev docker-compose.dev.yml****

      // for dev environment
      docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

      // for production environement
      // note the changes now made in server.js will
      // note reflect back as we didnt include
      //volumes bind mount in prod.yml
      docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

      // to reflect back changes in the prod //environement the image has to be rebuild each time
      docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build

====================================

> ## IMPORTANT To avoid dev-dependancie to get installed in our production env
(their is some issue in the bash script in Dockerfile
debug later)

=====================================
****we modify our Dockerfile with embedded bash script that excludes dev dependancies if production environment.****

        // make sure the spaces are correct
        //[ "something" ]
        ARG NODE_ENV
        RUN if [ "$NODE_ENV" = "development" ]; \
                then npm install; \
                else npm install --only=production; \
                fi

        docker exec -it containerName bash
        printenv
        // will show the environment variables and
        // production or development environment.

========================================

> ## Adding MongoDB docker container

========================================

****each new service we define like mongoDB , Node goes under the service section in the docker-compose.yml****
****refer :https://hub.docker.com/_/mongo****

       docker-compose -f docker-compose.yml -f docker-compose.dev.yml -d

       // to enter into the mongo shell Inside
       // docker container
       docker exec -it mongo_container mongo -u "databaseusername" -p "password"



==============================================

> ## Handling The Issue of Fresh Database set every time
we start the mongo container with docker-compose up

==============================================

****In order to presist the state of the Database when we run the docker-compose up next time we use named volumes  docker-compose.yml and the mongo section inside it.****

               volumes:
                  - mongo-db:/data/db # stores the volume with name mongo-db
            volumes:
            mongo-db:

****NOTE-> now when you use docker-compose down dont include the -v flag as it will also delete the named volume including the mongo-db one.****

          docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

          // to clean out all unecessary volumes
          //run prune when u have already run
          //docker-compose up so that the in use //container,volume do net gets deleted in //the prune process.

          docker volume prune

          // make sure your rebuild image if new package installed.
          docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build

          docker inspect node-container-address
          //under network we can see the default network created by docker-compose so all the containers can talk to each other with this network.

          docker inspect mongo-container-address
          //under network we have the IP address of the mongo container we want to grab this IP address to connect our node app and place in the mongo url in server.js.

          docker logs node-container-address
          // we will see the message that we connected to mongodb.

=========================================

> # Custom networks to make the containers talk to each other.

==========================================

            docker network ls
            //shows network host and bridge are default you will also see one with the name of your 'rootdirname_default' via which the two container are talking.

****Consider the case where we have to lookup for the IP address of the mongo container via docker inspect to mention it in the server.js to avoid this we can instead make use of custom networks that helps communication between containers.****

****So when we want one container/service to tak to another service/container we can do that by DNS i.e just by providing the name of the service we want to talk to****

****Example in the case we can talk to mongodb container by reffering to mongo so just mention mongo instead of IP address in server.js where u specified mongodb url****

           docker logs node_app_name -f
           // to follow i.e see the running container logs realtime

           docker exec -it first_example_node-app_1 bash
           #ping mongo
           // and you can access mongo container from the node container via DNS i.e specifiying the name of the service mongo in ther server.js file.

> ## IMPORTANT IN short you just need the service name to talk to it via other container as DNS is built inside of docker but note that it is only applicable on the network we create not on the bridge or host default networks.

          docker network inspect first_example_default

          //to see the container and gateway,subnet and other information on your network.


****Make sure to rebuild containers if you change your environment variables****

          docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

          docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

          docker logs first_example_node-app_1 -f



=========================================

> # Order Of Spinning Up Containers
> ## To Tell Docker to spin up the mongo container first as if node container spins up first while mongo is not running it can lead to issues

==========================================

****depends_on in yaml****

            // so in docker-compose.yaml in node_app
            depends_on:
              - mongo
            // this means that since node app container depends on the mongo container so mongo container will start up first when we do a docker-compose.


****now tear down adn start up the docker-containers again via docker-compose up -d you will se  that mongo starts first then node_app starts****

Creating first_example_mongo_1 ... done
Creating first_example_node-app_1 ... done


****To only start up a specific service with docker compose****

            docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d specific_service_name/node-app

****but this will start the mongo also due to depends_on*****


****To tackle this we use --no-deps it means dont start linked services****

            docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --help

            docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --no-deps node-app

            // only starts the node container

            docker logs first_example_node-app_1 -f

            // you will see that this keeps going on after 5 seconds
            MongooseServerSelectionError: connection timed out
            at NativeConnection.Connection.openUri (/app/node_modules/mongoose/lib/connection.js:846:32)
            at /app/node_modules/mongoose/lib/index.js:351:10
            at /app/node_modules/mongoose/lib/helpers/promiseOrCallback.js:32:5
            at new Promise (<anonymous>)
            at promiseOrCallback (/app/node_modules/mongoose/lib/helpers/promiseOrCallback.js:31:10)
            at Mongoose._promiseOrCallback (/app/node_modules/mongoose/lib/index.js:1149:10)
            at Mongoose.connect (/app/node_modules/mongoose/lib/index.js:350:20)
            at Timeout.connectWithRetry [as _onTimeout] (/app/server.js:14:12)
            at listOnTimeout (node:internal/timers:557:17)
            at processTimers (node:internal/timers:500:7) {
          reason: TopologyDescription {
            type: 'Single',
            setName: null,
            maxSetVersion: null,
            maxElectionId: null,
            servers: Map(1) { 'mongo:27017' => [ServerDescription] },
            stale: false,
            compatible: true,
            compatibilityError: null,
            logicalSessionTimeoutMinutes: null,
            heartbeatFrequencyMS: 10000,
            localThresholdMS: 15,
            commonWireVersion: null
          }
        }


****Finally run the mongo to remove this error****

            docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d mongo

++++++++++++++++++++++++++++++++++++++

> ## Interacting with redis Microservice Container CLI

++++++++++++++++++++++++++++++++++++++

          docker exec -it redis_container redis-cli

- [x] 'KEYS *' to see all the entries in the database
- [x] hit the login route and then rerun KEYS * to see the cookie.
- [x] GET the details of the cookie via GET "key"

          127.0.0.1:6379> KEYS *
          1) "sess:ZKsYOvtlGJ-6KvtpojZCX2dMwtyU3TdH"
          127.0.0.1:6379> GET "sess:ZKsYOvtlGJ-6KvtpojZCX2dMwtyU3TdH"
          "{\"cookie\":{\"originalMaxAge\":30000,\"expires\":\"2021-07-09T13:09:16.137Z\",\"secure\":false,\"httpOnly\":true,\"path\":\"/\"}}"


&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

> ## Session & Cookies

&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

- [x] Remember that http is naturally a stateless protocol it is the session and the cookies that make it statefull. consider example of the starbucks that provide you a card this is the cookie that is sent each time you request for a resource from the server and then this card/cookie is checked with some specific data with in the server so as to identify and isolate and assign a particular session to that user so that next time you visit it just saves your early prefferences.

- [x] cookie are always handled client side while the session sometimes but not always server side.
- [x] advantage of having session at client side is that it optimizes the response time as if we maintained session in server the cookie is first checked against some data on the server side and then starts/enable the session but a downside is the space as if we have to store fancy data like shoopping cart preferneces and it do not fits on client side then we have to eventually switch to server side.

        // at the server we assigned a user key with users username and password once the server logic checks the password
        127.0.0.1:6379> GET "sess:Ut1pbjxwEViTEIJjV2qAE8O6MymyS2RI"
        "{\"cookie\":{\"originalMaxAge\":30000,\"expires\":\"2021-07-09T13:34:33.929Z\",\"secure\":false,\"httpOnly\":true,\"path\":\"/\"},\"user\":{\"_id\":\"60e84767fc29af0064da84d0\",\"username\":\"john\",\"password\":\"$2a$12$r.ydgZxuTeSCQUAag.NCWOB9ncF3aZ6M49NwSjw7Wt0fI7.z8Tdaa\",\"__v\":0}}"

> ## Docker nature of  security
- [x] In the first_example the outside port 3000 is mapped to the server 3000 express container and then the express is talking to our mongoDB so their is no way that someone can directly interact with the mongo database contianer.

==============================

> ## Scaling via Docker & Adding Nginx Container as middleware between the outside world and express containers.

==============================

****We spin up another node express container that can interact with mongo container with the same internal port as 3000 while the external port 3001 via which  the outer world can interact with this new spin up node container****

- [x] first step is to expose a port 3001 via which outside world can interact with the node express container.

- [x] ****Refer default.conf****

            server{
              listen 80;

              location / {
                  // to catch the origin IP who make the request to the nginx further rate limiting or IP blacklisting could be done.
                  proxy_set_header X-Real-IP $remote_addr;
                  proxy_pass http://node-app:3000;
              }
            }

- [x] To link default.conf with the nginx we use bind mount.other way could be making a custom image but it is too hectic.

          3000 or external port--> port 80--> nginx 3000---> express1
                                                          |          |-> 27017 ->Mongo
                                                          |-> express2
