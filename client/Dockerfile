# Use an official Node.js runtime as the base image
FROM node:14.17.0-alpine

# Set the working directory in the container
WORKDIR /client

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the client application code to the working directory
COPY . .

# Build the client application
RUN npm run build

# Set the environment variable for production
ENV NODE_ENV=production

# Expose the port on which the client application will run (change if needed)
EXPOSE 3000

# Start the client application
CMD ["npm", "start"]
