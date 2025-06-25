#!/bin/bash

export DATABASE_URL="postgres://postgres:postgres@localhost:5432/clean-db?sslmode=disable"
migrate -database "$DATABASE_URL" -path ./migrations "$@"
