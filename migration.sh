#!/bin/bash

if [ -n $1 ]
then
  if [[ "$1" == "create" ]]
  then
    migrate create -ext sql -dir ./schema -seq init
    migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
  exit 0
  fi
fi

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
