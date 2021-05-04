> # Understanding Docker File with a simple Node.js snippet


****refer Dockerfile in docker_file_basics****

- Dockerfile is a set of instruction on the basis of which an image can be created which then we can build and then run as container.


> ## Elements of Dockerfile

- [x] FROM (Base Image)

****This instruction sets the base image, searching in the docker registery and try to find the image node:carbon tag.****

****Example :****

        FROM node:carbon

- [x] WORKDIR ****specifying the working directory for the docker/image container.****


        WORKDIR /src

- [x] RUN ****(command specifies what terminal command to run to install dependancies)****

****Navigate to the src directory of the container we are building and then install dependancies with npm install.****

        RUN cd /src;npm install

- [x] COPY ****the contents of our project to the image/container src directory where . means our current working directory of our project to the src directory of the docker container.****

       COPY . /src

- [x] ADD ****Same as COPY , on top of that we can use url's to ADD content /files to the image we are building unlike in the COPY where we can only copy the files from the local file system.****

      ADD ./hello.tar /root
      // this will add the hello.tar in the root folder of the container.

      ADD http://something.com /root
      // copying the files/content from url to the root directory of container


- [x] CMD ****this is executed once the container is up and running automatically these are generatlly shell/bash/terminal commands****

      CMD ["npm","start"]

- [x] ENTRYPOINT ****Looks similar to CMD but operates differently, basically used to execute shell/terminal commands these are executed by default at runtime i.e when container is running****

      ENTRYPOINT ls
      // lists the dir & files


- [x] ENV ****used t set environment variables for the container****

      ENV secretToken=SECRET

- [x] LABEL ****used to add metadata to the docker container like email,keyvalue pair , it is generally used as if we have a lot of docker images we can filter these images on the basis of these LABEL****
