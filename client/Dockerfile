# Use the official Node.js 14 image as the base image
FROM node

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install app dependencies
RUN npm install --force

# Copy the rest of the application code to the working directory
COPY . .

# Expose the port on which the app will run
EXPOSE 5173

# Start the app
CMD [ "npm", "run", "dev" ]
