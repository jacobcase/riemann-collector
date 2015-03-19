#!/bin/bash

HOST=`hostname`
DATE=`date +%s`
LOAD=`cat /proc/loadavg | awk '{print $1}'`

cat << EOF
{
    "time": $DATE,
    "description": "system load",
    "state": "ok",
    "service": "system",
    "host": "$HOST",
    "metric_d": $LOAD,
    "tags": [
        "cpu"
        ],
    "ttl": 10

}
EOF
