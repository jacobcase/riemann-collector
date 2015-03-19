#!/bin/bash

HOST=`hostname`
DATE=`date +%s`
LOAD=`cat /proc/loadavg | awk '{print $1}'`

cat << EOF
{
    "time": $DATE,
    "description": "woop woop woop",
    "state": "ok",
    "service": "system",
    "host": "$HOST",
    "metric_d": $LOAD,
    "tags": [
        "woop"
        ],
    "ttl": 10

}
EOF

#cat << EOF
#{"host": "wut","service": "testing","metric_d": 2.5}
#EOF
