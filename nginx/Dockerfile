# nginx/Dockerfile
# Stage 1: Build the React application
FROM node:latest as frontend-stage
WORKDIR /app
COPY ../frontend/package.json /app/
COPY ../frontend/package-lock.json /app/
RUN npm install
COPY ../frontend /app
RUN npm run build

# Stage 2: Serve the React application with nginx
FROM nginx:latest as production-stage
COPY --from=frontend-stage /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
