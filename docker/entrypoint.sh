#!/bin/bash -e


APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

ls -l

echo "[`date`] Starting server..."
./server