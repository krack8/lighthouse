#!/bin/bash

echo '[INFO] Starting Config Initialization'

echo $KLOVERCLOUD_API_ENDPOINT

find /usr/share/nginx/html/main*.js -type f -exec sed -i 's@KLOVERCLOUD_API_ENDPOINT@'"$KLOVERCLOUD_API_ENDPOINT"'@' {} 

echo '[INFO] Config Initialization Completed'

echo '[INFO] Starting Nginx'

nginx -g 'daemon off;'
