#!/bin/bash

HOST=`hostname`
DATE=`date +%s`
LOAD=`cat /proc/loadavg | awk '{print $1}'`

cat << EOF
{
    "time": $DATE,
    "state": "ok",
    "service": "system",
    "host": "$HOST",
    "metric_d": $LOAD
    "tags": [
        "woop"
        ]
    "ttl": 10.1
}
EOF
