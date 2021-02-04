#!/bin/bash
if [ "$(docker ps -q -f name=task-management-mongo-db)" ]; then
  docker rm -f task-management-mongo-db
fi
