# backup/Dockerfile
FROM postgres:17-alpine

# ติดตั้ง dependencies ที่จำเป็น
RUN apk add --no-cache \
    bash \
    curl \
    pigz \
    tar \
    dcron \
    tzdata

# Copy backup script
COPY backup.sh /backup.sh
RUN chmod +x /backup.sh

# สร้าง directory สำหรับ log
RUN mkdir -p /var/log/cron && \
    touch /var/log/cron/cron.log && \
    chmod 0644 /var/log/cron/cron.log

# Create a script to setup and run cron
RUN echo '#!/bin/sh' > /entrypoint.sh && \
    echo 'printenv | grep -v "no_proxy" >> /etc/environment' >> /entrypoint.sh && \
    echo 'crond -f -d 8 >> /var/log/cron/cron.log 2>&1' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]