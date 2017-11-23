FROM nginx

COPY index.html /usr/share/nginx/html
COPY default.conf /etc/nginx/conf.d

COPY start.sh start.sh
RUN chmod +x /start.sh

EXPOSE 80

CMD ["./start.sh"]
