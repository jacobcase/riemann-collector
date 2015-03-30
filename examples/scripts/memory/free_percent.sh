#!/bin/bash

HOST=`hostname`
DATE=`date +%s`
TOTALMEM=`cat /proc/meminfo | grep 'MemTotal' | awk '{ print $2 }'`
FREEMEM=`cat /proc/meminfo | grep 'MemFree' | awk '{ print $2 }'`
PERCENT=$(bc -l <<< "scale = 4;  100 * ($FREEMEM / $TOTALMEM)")

cat << EOF
{
    "time": $DATE,
    "description": "memory used percent",
    "state": "ok",
    "service": "system",
    "host": "$HOST",
    "metric_d": $PERCENT,
    "tags": [
        "memory"
        ],
    "ttl": 10

}
EOF
