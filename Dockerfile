FROM alpine:latest

RUN apk add --no-cache nginx

RUN getent group nginx || addgroup -S nginx && \
    id -u nginx &>/dev/null || adduser -S -G nginx nginx

RUN mkdir -p /var/cache/nginx /run/nginx && \
    chown -R nginx:nginx /var/cache/nginx /run/nginx

COPY nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

USER nginx

CMD ["nginx", "-g", "daemon off;"]