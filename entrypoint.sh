#!/bin/bash
set -e

# Check if blog.db exists and if its size is greater than 0
if [ ! -f /data/blog.db ] || [ ! -s /data/blog.db ]; then
  echo "Copying blog.db to /data directory..."
  cp /usr/local/bin/blog.db /data/
else
  echo "blog.db already exists and is not empty."
fi

# Run the main application
exec "$@"
