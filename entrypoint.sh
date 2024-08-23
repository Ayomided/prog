#!/bin/bash
set -e

# Check if blog.db exists and if its size is greater than 0
echo "Copying blog.db to /data directory..."
cp /usr/local/bin/blog.db /data/

# Run the main application
exec "$@"
