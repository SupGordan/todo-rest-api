#!/bin/bash -e


APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."


echo "[`date`] Starting server..."
./todo-rest-api