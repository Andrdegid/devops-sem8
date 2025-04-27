FROM node:20-alpine

RUN adduser -D appuser

WORKDIR /app
COPY . .

RUN npm install && \
    npm run build && \
    npm cache clean --force

USER appuser

CMD ["npx", "serve", "-s", "build", "-l", "3000", "--single"]