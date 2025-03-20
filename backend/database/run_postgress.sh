#!/bin/bash

# Terminate script on error
set -e

# Get the script directory (Ensures correct .env path)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ENV_FILE="$SCRIPT_DIR/../.env"  # Adjusted to point to backend/.env

# Load environment variables from .env file
if [ -f "$ENV_FILE" ]; then
    export $(grep -v '^#' "$ENV_FILE" | xargs)
else
    echo "❌ .env file not found at $ENV_FILE!"
    exit 1
fi

# Default DB_PORT if not set in .env
DB_PORT=${DB_PORT:-5432}  # Uses 5432 if DB_PORT is not set

# PostgreSQL container name
CONTAINER_NAME="postgres"

# Stop and remove existing PostgreSQL container if it exists
if [ "$(docker ps -a -q -f name=$CONTAINER_NAME)" ]; then
    echo "🚀 Stopping and removing existing PostgreSQL container..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# Start a new PostgreSQL container using your `.env` DB variables
echo "🚀 Starting PostgreSQL container on port $DB_PORT..."
docker run -d --name $CONTAINER_NAME \
  -e POSTGRES_USER=$DB_USER \
  -e POSTGRES_PASSWORD=$DB_PASSWORD \
  -e POSTGRES_DB=$DB_NAME \
  -p $DB_PORT:5432 postgres

# Wait for PostgreSQL to fully start
echo "⏳ Waiting for PostgreSQL to become ready..."
until docker exec $CONTAINER_NAME pg_isready -U $DB_USER > /dev/null 2>&1; do
  sleep 2
done

echo "✅ PostgreSQL is ready!"

# Create database and user (if not created by env vars)
echo "✅ Setting up database and user..."
docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME <<EOF
CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
ALTER USER $DB_USER WITH SUPERUSER;
GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
EOF

echo "🎉 PostgreSQL is ready! Use this in your .env file:"
echo "DATABASE_URL=postgres://$DB_USER:$DB_PASSWORD@localhost:$DB_PORT/$DB_NAME?sslmode=disable"