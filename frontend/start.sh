#!/bin/sh

/bin/sed -i "s|__APP__|${APP}|" /etc/nginx/conf.d/default.conf

nginx -g "daemon off;"
