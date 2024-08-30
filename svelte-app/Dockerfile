# Use an official Node runtime as the base image
FROM node:18 as local

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

COPY ./config/config.dev.json ./build/config.json

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "3000"]

# Stage 1: Build the SvelteKit app
FROM node:18 AS builder

WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the SvelteKit app
RUN npm run build

# Stage 2: Serve the built app with nginx
FROM nginx:alpine as prod

# Copy the build output to the nginx html folder
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Override the default nginx configuration
CMD ["nginx", "-g", "daemon off;"]
