# Stage 1: Build the application
FROM node:20-alpine AS builder
WORKDIR /app
COPY /web /app 
RUN npm ci
RUN npm run build
RUN npm prune --production

# Stage 2: Create the final image
EXPOSE 3000
ENV NODE_ENV=production
CMD ["node", "build"]
