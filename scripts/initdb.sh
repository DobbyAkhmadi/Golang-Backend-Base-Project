#!/bin/bash
# This script initializes the database with migrations and seeds

# Run database migrations
echo "Running database migrations..."
go run ./platform/migrations

# Seed the database with initial data (if needed)
echo "Seeding the database..."
go run ./platform/seeds

echo "Database initialization complete."
