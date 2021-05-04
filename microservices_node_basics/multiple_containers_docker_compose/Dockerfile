FROM node:carbon


# set working directory
WORKDIR /src

# Copy all the node.js project files from local working dir to container working dir /src
COPY . .

# Install App dependencies
RUN npm install

EXPOSE 8080

CMD ["npm","start"]
