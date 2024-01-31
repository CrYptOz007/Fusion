FROM node:21.6-bookworm-slim

WORKDIR /var/www/app/frontend

COPY . .

RUN npm ci

RUN npm run build

USER node:node

CMD ["npm", "run", "dev"]