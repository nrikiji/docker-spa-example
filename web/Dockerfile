FROM node:16.14.0 as build

WORKDIR /srv/app
ADD ./package.json .
ADD ./package-lock.json .
RUN npm install

COPY . .
RUN npm run build

FROM nginx:1.20.2
COPY --from=build /srv/app/public /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
