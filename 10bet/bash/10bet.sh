#!/bin/bash

echo "Getting data"

time JSON=`curl -s 'https://sbapi.sbtech.com/10betcom/sportscontent/sportsbook/v1/Events/GetBySportId'   -H 'authority: sbapi.sbtech.com'   -H 'pragma: no-cache'   -H 'cache-control: no-cache'   -H 'block-id: Center_TopLeaguesResponsiveBlock_49897'   -H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTaXRlSWQiOjU2LCJTZXNzaW9uSWQiOiI4NzA5ZWNjMS01ZjM1LTQ4NDEtOThmZC1hZjBmYzNmZmI3ZDQiLCJuYmYiOjE1OTUwNTM2NjgsImV4cCI6MTU5NTY1ODQ5OCwiaWF0IjoxNTk1MDUzNjk4fQ.2yym9FNqjK7Lcua6SrVKj5J0QwZqla7DqFs-dvD7y-A'   -H 'locale: en'   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36'   -H 'content-type: application/json-patch+json'   -H 'accept: */*'   -H 'origin: https://www.10bet.com'   -H 'sec-fetch-site: cross-site'   -H 'sec-fetch-mode: cors'   -H 'sec-fetch-dest: empty'   -H 'referer: https://www.10bet.com/sports/politics/'   -H 'accept-language: '   --data-binary '{"eventState":"Mixed","eventTypes":["Outright"],"ids":["152"],"eventTags":[],"leagueState":"Regular","pagination":{"top":100,"skip":0}}'   --compressed`

echo "Parsing json"

time DATA=`echo "$JSON" | jq '.markets[] | select (.id == "55405316").selections[] | .name + " " + (.trueOdds|tostring)'`

echo "Parsing finished"

echo $DATA