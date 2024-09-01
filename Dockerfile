# Use an official Node runtime as a parent image
FROM node:20-slim

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install production dependencies.
RUN npm install --only=production

# Copy the local code to the container's workspace.
COPY . .

# Expose the port the app runs on
EXPOSE 3000

# Run the app when the container launches
CMD ["node", "server.js"]

