#!/bin/bash

# Script to generate Go Gin Postgres CRUD project structure
# Usage: ./generate_structure.sh [project_name]
# Default project: go-gin-postgres-crud

PROJECT_NAME=${1:-go-gin-postgres-crud}

echo "Creating project structure for $PROJECT_NAME..."

# Create root directory
# mkdir -p "$PROJECT_NAME"
# cd "$PROJECT_NAME" || exit 1

# Create directories
mkdir -p config models handlers routes

# Create empty files
touch go.mod main.go
touch config/db.go
touch models/user.go
touch handlers/user_handler.go
touch routes/routes.go

echo "Project structure created successfully!"
echo ""
tree . || find . -type f | sort
echo ""
echo "Next steps:"
echo "1. cd $PROJECT_NAME"
echo "2. go mod init github.com/yourusername/$PROJECT_NAME"
echo "3. go get github.com/gin-gonic/gin github.com/lib/pq"  # or gorm.io/gorm etc.
echo "4. Implement the files (e.g., db connection in config/db.go, models, handlers, routes)."
