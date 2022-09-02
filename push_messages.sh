#!/bin/bash

set -ex

curl -v -X POST https://api.line.me/v2/bot/message/push \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer ${LINE_CHANNEL_ACCESS_TOKEN}" \
-d "{
    \"to\":\"$1\",
    \"messages\":[
        {
            \"type\":\"text\",
            \"text\":\"$2\"
        }
    ]
}"
