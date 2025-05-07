FROM node:22-alpine AS build

WORKDIR /app

COPY . .

RUN npm run build

FROM node:22-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install --omit=dev

COPY --from=build /app/build ./build

EXPOSE 3000

ENV NODE_ENV=production

CMD ["node", "build"]