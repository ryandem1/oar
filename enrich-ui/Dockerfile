FROM node:20 AS build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY .. ./
RUN npm run build

FROM nginx:1.24-alpine
COPY --from=build /app/.svelte-kit /usr/share/nginx/html
