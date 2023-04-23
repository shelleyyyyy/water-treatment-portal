# FROM node:19

# WORKDIR /app

# COPY package.json ./
# COPY package-lock.json ./
# RUN npm install
# COPY . ./
# EXPOSE 5173
# CMD ["npm", "run", "dev", "--", "--host", "--port", "5173"]

FROM node:19 AS build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build

FROM nginx:1.19-alpine
COPY --from=build /app/build /usr/share/nginx/html