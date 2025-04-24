#!/bin/bash
# backup/backup.sh

# กำหนดตัวแปร
BACKUP_DIR="/backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="backup_${TIMESTAMP}"

# สร้างโฟลเดอร์ backup ถ้ายังไม่มี
mkdir -p ${BACKUP_DIR}

# ทำ Database Backup
echo "Starting backup of PostgreSQL database: ${POSTGRES_DB}"
PGPASSWORD=${POSTGRES_PASSWORD} pg_dump -h ${POSTGRES_HOST} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -F c -b -v -f "${BACKUP_DIR}/${BACKUP_FILE}.backup"

# บีบอัดไฟล์ backup
echo "Compressing backup file..."
pigz "${BACKUP_DIR}/${BACKUP_FILE}.backup"

# ลบไฟล์ backup เก่า
echo "Removing old backups..."
find ${BACKUP_DIR} -type f -name "*.backup.gz" -mtime +${BACKUP_RETENTION_DAYS} -delete

# ตรวจสอบสถานะการทำงาน
if [ $? -eq 0 ]; then
    echo "Backup completed successfully: ${BACKUP_FILE}.backup.gz"
else
    echo "Backup failed!"
    exit 1
fi

# สร้าง symlink ไปยัง backup ล่าสุด
ln -sf "${BACKUP_DIR}/${BACKUP_FILE}.backup.gz" "${BACKUP_DIR}/latest.backup.gz"
