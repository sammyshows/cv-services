# Use the official Node.js 20 image as a parent image
FROM node:20

# Set the working directory
WORKDIR /usr/src/app

# Copy the current directory contents into the container at /usr/src/app
COPY . .

# Install any needed packages specified in package.json
RUN npm install

# Make port 3000 available to the world outside this container
EXPOSE 3000

# Run the app when the container launches
CMD ["npx", "nodemon", "--watch", "src", "--ext", "js,ts,json", "--exec", "ts-node src/app.ts"]
