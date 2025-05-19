#!/bin/bash
# This script generates random tasks and sends them to $ADDRESS each second.
ADDRESS="http://localhost:8080/api/tasks"
TOOLS=("upgrader" "downgrader" "watcher" "listener" "system-monitor")
TASKS=("healthchecks" "upgrade" "rollback" "backup" "restore")
STATUSES=("completed" "failed" "succeeded")

echo "Sending random events to $ADDRESS. press Ctrl+C to abort."
while [ 1 ]; do
    tool=${TOOLS[ $RANDOM % ${#TOOLS[@]} ]}
    task=${TASKS[ $RANDOM % ${#TASKS[@]} ]}
    status=${STATUSES[ $RANDOM % ${#STATUSES[@]} ]}
    PAYLOAD="{\"tool\": \"$tool\", \"task\": \"$task\", \"status\": \"$status\", \"duration\": $RANDOM}"
    echo $PAYLOAD
    curl --header "Content-Type: application/json" --request "POST" --data "$PAYLOAD" \
        $ADDRESS
    echo ""
    sleep 1
done
