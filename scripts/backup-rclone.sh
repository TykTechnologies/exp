#!/bin/bash
# Variables
HOSTNAME=$(uname -n)
DATE=$(date +%Y%m)
DEST="/home/black/backup-${HOSTNAME}-${DATE}"
RCLONE_REMOTE="Lab"

# Directories to back up
INCLUDE_DIRS=("/etc" "/home" "/root")

# Exclude patterns
EXCLUDES="--exclude **/.cache/**"

set -x

rclone mkdir ${RCLONE_REMOTE}:${DEST}

# Sync each directory
for DIR in "${INCLUDE_DIRS[@]}"; do
    echo "Backing up $DIR..."
    rclone sync $DIR ${RCLONE_REMOTE}:${DEST} $EXCLUDES --progress
done

echo "Backup completed."
