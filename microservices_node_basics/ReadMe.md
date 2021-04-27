# Node.js Microservices Basics Hands On

****By- Jasmeet****

****Level : Intermediate****

### Prerequisites:

#### - Javascript
#### - Docker
#### - REST API

# Aim/Covered Topics

## Hands On in the following topics->

### - what , when & how to Build microservices
### - microservices monitoring,logging,testing,& scaling services
### - Implement d/f microservice pattern using docker-compose (EDA,CQRS Gateway)
### - Understanding different tools for setting up microservice on loacal system (docker & k85)
### - Building services & Deploying them cerating microservices via node.js with docker and k85.

## => Fundamental Characterstic every microservice should have

### Zero Configuration
- Any Microservice system will have hundreds of services so a manual configuration of addresses,ports,IP and API capabilities is infeasible.

### Highly redundant

- Service failure are common in this architecture so their must be cheap copies available at your disposal with proper failover mechanism.

### Fault-tolerant

- the system should tolerate and gracefully handle miscommunication,errors,processes,messages,timeouts etc
- ****Even if certain services are down the other unrelated services must continue to run****

### Self Healing
- Its noramal to have errors and system fails the system must automatically recover failed services and functionality.

### Auto-discovery
- the services should automatically identify new services that are introduced to the system to start communication without manual intervention or downtime.

## => Docker and Kubernetes and Containerization

- Helps different docker containers that can be of same stack tech or different tech stack to talk to one another by container orchesterations.

- ****like backend microservice in python/ django communicating with Front-end microservice built in React.js.****

- ****or like MONGODB is communicating with front-end ruby on rails****

                             Kubernetes/docker/containerization

                              front(microserv A)
      client      ------>            backend(microserv B)
                              Database(microserv C)

## => Service Registry (IMPORTANT)

- ****Consider their are different services instances REST API that are new and are going to be added to already existing microservice instance X then the new services will register themselves with the service registry****

- ****Now say if a new service or old service want to interact with the Auth microservice then it will contact the service registry directly to determine which service is auth service instead of going to each and every service and asking are you auth service.****

      - microservice A --------->  SERVICE
      - microservice B --------->  REGISTERY <----microservice X(finding microservice C will contact service Registery)
      - microservice  C --------->  SYSTEM

## => API gateway & lambda

- ****We can provide various rate limiting and IAM access policy via API gateway to ensure the availability of api endpoints to only certain individual/organization.****

- ****lambda defines the rule for the particular API for a microservice when the client want to talk to the microservice via the API****

       - Client--->API gateway---->lambda---->microservice

## => Complex Microservices

- ****Every Microservice wheather mobile or web they will have a SINGLE SOURCE  and the failure of a single microservice will not impact the system.****

      - mob app-->API gateway---->(REST API) Account microservice--->Account DB
      - Web --->front web APP---->(REST API) Inventory microservice---->Inventory DB

## => Fundamentals of Designing Microservices (IMPORTANT)

####  Each Microservice must have their own unique Data source/DB No two microservices should share any data source as we saw earlier in mob app vs web above to ensure Fault tolerance. failure of one dont effect the entire system.

================== Theory Ends Here ==========================
***
