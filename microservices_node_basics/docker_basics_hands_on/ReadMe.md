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

TIMESTAMP->https://www.youtube.com/watch?v=1rAx4-sHRKU&list=PLIGDNOJWiL1-svqMFkNEiNdDyhs41Vnib&index=5
