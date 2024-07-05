FROM node:latest as frontend-stage
COPY frontend /app
WORKDIR /app
RUN npm install
RUN npm run build
FROM nginx:latest as production-stage
COPY --from=frontend-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]