#!/bin/bash

# Define the project structure
PROJECT_NAME="project"

DIRS=(
    "$PROJECT_NAME/cmd/service"
    "$PROJECT_NAME/internal/app"
    "$PROJECT_NAME/internal/domain/entity"
    "$PROJECT_NAME/internal/domain/service"
    "$PROJECT_NAME/internal/ports/http"
    "$PROJECT_NAME/internal/ports/messaging"
    "$PROJECT_NAME/internal/ports/repository"
    "$PROJECT_NAME/pkg/logger"
    "$PROJECT_NAME/pkg/config"
    "$PROJECT_NAME/pkg/errors"
    "$PROJECT_NAME/pkg/middleware"
    "$PROJECT_NAME/pkg/utils"
    "$PROJECT_NAME/pkg/database"
    "$PROJECT_NAME/api/proto"
    "$PROJECT_NAME/config"
    "$PROJECT_NAME/docs"
    "$PROJECT_NAME/scripts"
    "$PROJECT_NAME/migrations"
    "$PROJECT_NAME/test"
)

FILES=(
    "$PROJECT_NAME/cmd/service/main.go"
    "$PROJECT_NAME/internal/app/service.go"
    "$PROJECT_NAME/internal/domain/entity/user.go"
    "$PROJECT_NAME/internal/domain/service/user_service.go"
    "$PROJECT_NAME/internal/ports/http/handlers.go"
    "$PROJECT_NAME/internal/ports/messaging/kafka.go"
    "$PROJECT_NAME/internal/ports/messaging/rabbitmq.go"
    "$PROJECT_NAME/internal/ports/repository/user_repository.go"
    "$PROJECT_NAME/internal/ports/repository/redis.go"
    "$PROJECT_NAME/internal/ports/repository/mongo.go"
    "$PROJECT_NAME/internal/ports/repository/postgres.go"
    "$PROJECT_NAME/api/proto/user.proto"
    "$PROJECT_NAME/config/config.go"
    "$PROJECT_NAME/config/config.yaml"
    "$PROJECT_NAME/pkg/logger/logger.go"
    "$PROJECT_NAME/pkg/config/config.go"
    "$PROJECT_NAME/pkg/errors/errors.go"
    "$PROJECT_NAME/pkg/middleware/middleware.go"
    "$PROJECT_NAME/pkg/utils/utils.go"
    "$PROJECT_NAME/pkg/database/database.go"
    "$PROJECT_NAME/Dockerfile"
    "$PROJECT_NAME/docker-compose.yml"
    "$PROJECT_NAME/migrations/README.md" # Placeholder for migration documentation
    "$PROJECT_NAME/test/unit_test.go" # Placeholder for unit tests
    "$PROJECT_NAME/test/integration_test.go" # Placeholder for integration tests
)

# Create directories
for dir in "${DIRS[@]}"; do
    mkdir -p "$dir"
done

# Create files with package names
for file in "${FILES[@]}"; do
    touch "$file"
    case "$file" in
        *cmd/service/main.go)
            echo "package main" > "$file"
            ;;
        *internal/app/service.go)
            echo "package app" > "$file"
            ;;
        *internal/domain/entity/user.go)
            echo "package entity" > "$file"
            ;;
        *internal/domain/service/user_service.go)
            echo "package service" > "$file"
            ;;
        *internal/ports/http/handlers.go)
            echo "package http" > "$file"
            ;;
        *internal/ports/messaging/kafka.go)
            echo "package messaging" > "$file"
            ;;
        *internal/ports/messaging/rabbitmq.go)
            echo "package messaging" > "$file"
            ;;
        *internal/ports/repository/user_repository.go)
            echo "package repository" > "$file"
            ;;
        *internal/ports/repository/redis.go)
            echo "package repository" > "$file"
            ;;
        *internal/ports/repository/mongo.go)
            echo "package repository" > "$file"
            ;;
        *internal/ports/repository/postgres.go)
            echo "package repository" > "$file"
            ;;
        *api/proto/user.proto)
            # Protocol buffers file does not need a package declaration in Go
            ;;
        *config/config.go)
            echo "package config" > "$file"
            ;;
        *pkg/logger/logger.go)
            echo "package logger" > "$file"
            ;;
        *pkg/config/config.go)
            echo "package config" > "$file"
            ;;
        *pkg/errors/errors.go)
            echo "package errors" > "$file"
            ;;
        *pkg/middleware/middleware.go)
            echo "package middleware" > "$file"
            ;;
        *pkg/utils/utils.go)
            echo "package utils" > "$file"
            ;;
        *pkg/database/database.go)
            echo "package database" > "$file"
            ;;
        *test/unit_test.go)
            echo "package test" > "$file"
            ;;
        *test/integration_test.go)
            echo "package test" > "$file"
            ;;
        *Dockerfile|*docker-compose.yml|*migrations/README.md)
            # Dockerfile and docker-compose.yml do not require package declarations
            ;;
    esac
done

# Initialize the Go module
cd "$PROJECT_NAME" || exit
go mod init your_project_name

echo "Project structure created successfully."
