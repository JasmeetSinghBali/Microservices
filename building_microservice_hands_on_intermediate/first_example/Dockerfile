FROM node:15
# Specify the Base Image on which our custom image will be built.
# the above cmd will pull the node image version 15 from docker hub.

# To Set Up the working directory within our container so inside the container we will have an app directory that will be contaning all the code related to our app.
# Setting up the working directory helps to run any cmd to run from the working directory we specify.
# example we specify cmd node app.js then this cmd will automatically run from the workdir we specify.

WORKDIR /app

# Copy cmd to copy the package.json file so that we can include our dependencies in the docker.
# the below command will copy the package.json from the local dir to the docker workdir we specified we reffered the /app workdir via .

COPY package.json .


# To finally installed the dependancies in the docker container we can use the RUN cmd

RUN npm install

# Copy rest of our app files from our local dir to the docker workdir /app
# Copy all the files & folders specified by . from our local dir to the workdir /app specified by .

COPY . .

# Since our application is running on localhost:5000 so we expose the container port 5000, however technically we are just mentioning a type of documentation as the Expose statement do not has any real effect on how docker interact to the outer world.
# Note that Docker can access outside world but the outside world cannot interact with the docker container this is a by default security mechanism

# a env variable key(PORT) and value(5000) for the DOCKER container reffered as $PORT now.
ENV PORT 5000
EXPOSE $PORT

# CMD to run when the container is deployed i.e starts/live

CMD ["npm", "run", "dev"]

# finally to build the image 'docker build -t node-app-example .(path to the docker file)
