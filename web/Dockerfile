# Use Bun as the base image
FROM oven/bun:1.0.2 as builder

# Set working directory
WORKDIR /app

# Copy package.json, bun.lockb, and install dependencies
COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

# Copy the entire web application code
COPY . .

# Build the application
RUN bun build

# Use Nginx to serve the built files
FROM nginx:alpine

# Copy the build files to the Nginx HTML directory
COPY --from=builder /app/build /usr/share/nginx/html

# Expose the default Nginx port
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
