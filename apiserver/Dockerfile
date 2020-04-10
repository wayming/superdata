# pull official base image
FROM node:13.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./

RUN npm install --silent
RUN npm install express --silent
RUN npm install pg --silent
RUN npm install cors --silent
# add app
COPY . ./

# start app
CMD ["npm", "start"]

