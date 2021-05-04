> # How to Run node.js Application in as docker container

- [x] docker build, navigate to the project root directory in terminal and build the docker image on basis of Dockerfile.

     docker build -t mynodejs .

- [x] docker run, to access the container we will map the port of the local system to the exposed port which we specified in Dockerfile to access the container from the local system.

     docker run -p 5001:3000 imageId/imageTag
     // this will map port 3001 of the local system to the container port 3000 that we exposed when we build image via the Dockerfile.

- [x] go to localhost:5001 to see your node application
