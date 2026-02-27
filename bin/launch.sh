#!/bin/bash

APP_NAME="site-api"
APP_PATH="./"
CONFIG_PATH="./"
APP_EXEC="$APP_PATH/$APP_NAME $CONFIG_PATH"
PID_FILE="$APP_PATH/$APP_NAME.pid"

start_app() {
    if [ -f "$PID_FILE" ]; then
        echo "$APP_NAME is already running."
    else
        echo "Starting $APP_NAME..."
        # nohup $APP_EXEC > $APP_PATH/$APP_NAME.log 2>&1 &
        nohup $APP_EXEC > /dev/null 2>&1 &
        echo $! > "$PID_FILE"
        echo "$APP_NAME started with PID $(cat $PID_FILE)."
    fi
}

stop_app() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        echo "Stopping $APP_NAME with PID $PID..."
        kill $PID
        rm "$PID_FILE"
        echo "$APP_NAME stopped."
    else
        echo "$APP_NAME is not running."
    fi
}

restart_app() {
    stop_app
    start_app
}

status_app() {
    if [ -f "$PID_FILE" ]; then
        echo "$APP_NAME is running with PID $(cat $PID_FILE)."
    else
        echo "$APP_NAME is not running."
    fi
}

case "$1" in
    start)
        start_app
        ;;
    stop)
        stop_app
        ;;
    restart)
        restart_app
        ;;
    status)
        status_app
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status}"
        exit 1
        ;;
esac