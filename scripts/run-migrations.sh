#!/bin/bash
export POSTGRES_USER=muhriddin
export POSTGRES_PASSWORD=1
export POSTGRES_DATABASE=book_shop
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path "./migrations"  up
