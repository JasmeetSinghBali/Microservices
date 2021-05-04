# Docker Fundamentals & Commands

****By: Jasmeet****

****- Docker is based on of linux container technologies.****
- the idea is to make a small linux container on the host operating system.
- a mini linux system providing you small sandbox.
- can have multiple containers running on single host.

## VM's(Virtual Machines) Vs Docker Container

#### VM's

- A single host and rest are guest operating system that are virtual machines on this single host machine.
- need hypervisor and their are limitations on how many guest operating system can be their on the host.

- ****hypervisor : A hypervisor is a kind of emulator; it is computer software, firmware or hardware that creates and runs virtual machines****

#### Docker Container

- Multiple guest/app/microservices are basically packaged as docker container and can be spinned up on a single host machine via container engine.

## Docker Internals (IMPORTANT)

### chroot , cgroups and namespaces

                     Docker
                       |
          libvirtual LXC systemdnsspawn
                       |
                  Linux Kernel
                       |
    crgroups SElinux AppArmor netlink namespaces chroot            

### What is Linux containers (LXC) (IMPORTANT)

### - Lightweight virtualization env provided by the linux kernel to provide system level virtualization without running hypervisor.
### - Three core concepts for resource isolation Cgroups, namespaces, Chroot

        Container=Collection of namespaces(partition of the system resources and creating isolated workspaces called containers) & cgroups(restrict access to certain process/processes)

- ****Linux provides the namesspaces and cgroups that are utilized by docker engine to create ISOLATED CONTAINERS on the host machine.****

                     Host
      Container 1     Container 2    Container 3
       (namespaces)    (namespaces)   (namespaces)
           |                |                |
       {cgroups}          {cgroups}       {cgroups}
                     KERNEL
#### - > What are Namespaces(Isolated Containers/workspaces) in linux

- ****Namespaces are a feature of the Linux kernel that partitions kernel resources such that one set of processes sees one set of resources while another set of processes sees a different set of resources.****

- ****Linux kernel**** provide namespaces that refers to d/f workspaces to ****isolated workspace are namespaces and these isolated workspace is basically called Container****

- ****thier are d/f namespaces PID namespace, IPC namespaces, mount namespaces, user namespaces ,network namespaces****

#### - > What are Cgroups

- ****It is a Kernel feature to restrict access to a system resource for a set of processes or individual process.****

           Resources          cgroups                tasks

             CPU1    =>      Cgroup-1    ------ >     Task 1
             RAM1                                     Task 2

             CPU2
             RAM2    =>      Cgroup-2    ------- >    Task 3
             RAM3
- ****cgroup can be memory cgroups,network cgroups,storage cgroups,cpu cgroups****

# Summing Up
# To Sum Up Linux container(LXC) tech on which docker is built is basically VIRTUALIZED RESOURCE ISOLATION FRAMEWORK/ARCHITECTURE with clear segregation of Data and resources across various processes in the system.

***

# Docker Installation & Setup
- https://docs.docker.com/get-docker/

***
# Basic Commands and getting started with Docker

****Sources: https://docs.docker.com/get-started/****

- ****Basic Syntax****

             docker [option] [command] [arguments]

- ****to see available commands type docker****

             docker
- ****docker runs on host machine via daemon process.****

# Docker Architecture

- ****We are client interacting with docker via Daemon in host after installation on our local machone****

      CLIENT           |        HOST              |       REGISTRY
                                                             HUB
      CLI      ---->   |       DAEMON
                                 |                |    <---  pull Images
      or       ----->    Containers--> Images --->
      remote API --->  |                          |
                                                           Publish images
                       |                          |

- ****Core operations of Docker****
      docker build (build images  of our application)
      docker pull  (pulling an docker image from Docker hub)
      docker run   (running docker image on local machine that actually makes the docker image into isolated container.)

- Example
      docker pull redis
      // will pull the redis image from docker hub

## Docker Container

#### - It is an instance of the docker image
#### - Container represent execution of single process,service or application
#### - It consist content of docker image, execution env , and standard set of instructions.

## Repository in Docker Hub

#### A collection of Docker Images(mini Linux System) labelled with a tag and variants like SDK,runtime(Lightweight) & also different version.

## REGISTRY

#### A service that provides access to repos. the default registry of most public images is Docker hub(owned by Docker as an organization). Company have often private repo to store their Images.

***

# Docker Hands On

**** Docker checks wheather image is already available locally if not then it will pull imgage from docker hub****

      docker pull ubuntu

****To display all the images we have****

      docker images

****To see how many containers are their****

     docker ps

****to display all active containers****

     docker ps -a

****to remove containers we specify the container ID, we can also remove multiple containers in docker by specifiying the multiple container id seperated with spaces****

     docker rm containerID1 containerID2 containerID3

****say we want to run ubuntu image already on our local machine, we first run the image in intreactive mode and go inside the container(mini linux) via bash****

     docker run -it ubuntu /bin/bash
     exit// to exit the container bash
****to get details about a particular container in a different terminal****

     docker inspect containerID

****To remove containers that are already dead i.e not running****

    docker container prune

***

# Understanding Dockerfile.yaml & Build Image from Dockerfile

#### STEP 1 : Specify BASE image FROM

****FROM is used to specify the BASE IMAGE on which the container will be built****

****NOTE- FROM must be the first instruction in dockerfile****


    FROM baseimage:version

#### STEP 2 : LABEL adds metadata to an image

    LABEL maintainer='myname@onecompany.com'

#### STEP 3 : RUN to execute commands on the top of the current image as a new layer and commit results    

    RUN apt-get install nginx -y
    RUN apt-get install ngingx -y

#### STEP 4 : EXPOSE port i.e when container is live and running it can connect to outside world via port 80.

    EXPOSE 80

#### STEP 5 : CMD The actual command to start up NGINX within our container

****NOTE- Their can only be one CMD instruction in Dockerfile****

    CMD ["nginx","-g","daemon off;"]

#### STEP 6 : Finally build image from dockerfile in terminal

****to build docker via Dockerfile where . specify to build container in current working directory****

    docker build .

#### Step 7 : PORT MAPPING run the docker image u created with exposing the port 80 of the container and mapping to the port 80 of host

    docker run -p 80:80 imageID

#### Step 8 : Test that you are able to access the container via localhost.

    curl http://localhost

  ****will display the nginx page****

***

# Basic Docker Commands

****Command: docker search****

****Displays the docker image available that can be pulled****

     docker search redis
     docker search node

****Command : docker run****

     docker run -it ubuntu /bin/bash

****Command : docker container stop****

     docker container stop containerID

****Command : docker rmi imageID to remove a particular image note that the image must not be used by any running container as if it is actually used by any container then this command wont work****

     docker rmi Imageid

****Command : docker build****

     docker build . -t myImage //builds an image from a docker file in the current directory and tags the image.

****Command : docker run****

     docker run myImage

****Command : docker ps or docker ps -a to see all running containers****

****Command: docker logs to see logs for a particular container****

      docker logs containerID/myImage

****Command : list all networks****

      docker network ls

****Command : Remove one or more networks****

      docker network rm networkID

****Command :Show information on one or more networks****

      docker network inspect networkID

****Command : Related to Docker volume i.e host/local machine space we will be allocating to the docker that is required to say store data and presist data or database.****


      docker volume create
      // Create a volume
      docker volume inspect
      docker volume ls
      // list volumes
      docker volume prune
      // remove all unused local volumes
      docker volume rm

****Command : prune to delete everything****

     docker container prune // all containers will be deleted
     docker image prune // all images will be deleted
     docker volume prune // will remove all local volumes

#### NOTE-> ****Volume is used by the containers to store some data, cached keys ,config keys etc****

***

# Docker volume mapping and Port Mapping

## IMPORTANT  NOTE- Volumes are stored in part of host file system but managed by docker(/var/lib/docker/volumes) on linux.

****Stateless containers are where we dont use volume when we are not storing any data i.e NO STORAGE refers to stateless containers.****

****Port mapping is used to expose our containers to outside world / host machine so that the container can interact with the outside world or host machine.****

### Two ways to store files in the host machine so that the data persisted even after the container has stopeed.

#### - volumes
#### - bind mounts
#### if on linux can use tmpfs mount also


        docker run --name=myDocker -v volume:/var/jenkins_home -p 8181:8080 -p 50001:50000 jenkins

- ****here the -v refers to the container directory(jenkins_home) to which the volume provided by the local machine is mapped to.****

- ****volume part represents host and /var/jenkins_home represeent scontainer****

- ****the port mapping (host)8181:8080(container) is for jenkins(image name) and the port 50001:50000 is for API config****



       docker run -p 8080:80 nginx
       curl http//localhost:8080 // displays nginx page

***

# Docker Networking (How containers talk to each other)

#### - Host mode networking
#### - Bridge mode networking


****Example a Host mode networking container ping to google.com****

      docker run -it -d --name my_container busybox

      // if image busybox not their it will pull it and then run the image as container mt_container.

      docker exec -it my_container ping -w3 google.com

      // now the my_container is actually able to access the public sources like google.com

      docker inspect my_container
      // in the Network settings you can see the IP used by the container this IP can be used to access the container itself.

**** In case multiple dockers are their then they will create a docker bridge with which the containers can talk to one another achieved via  docker compose commands****

****Example****

      // host mode network
      docker run -d --net=host --name container webservice:latest

      // bridge mode networking via docker bridge
      // here each docker container will have their IP assigned with which they can interact with each other with eh exposed ports and with outside world.

      # docker run -d --name -p 8801:80 container1 webservice:latest
      # docker run -d --name -p 8802:80 container2 webservice:latest
