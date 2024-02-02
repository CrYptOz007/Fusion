FROM node:21.6-bookworm-slim

WORKDIR /var/www/app/frontend

COPY package*.json ./

COPY src ./

RUN npm ci

CMD ["npm", "run", "dev"]