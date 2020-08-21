FROM node:latest as node
workdir /app
copy . .
run npm install
run npm run build --prod

from nginx:alpine 
copy --from=node /app/dist/frontend /usr/share/nginx/html 
