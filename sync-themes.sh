#!/bin/sh
curl https://raw.githubusercontent.com/atomcorp/themes/master/app/src/backupthemes.json | jq -c > themes.json
curl https://raw.githubusercontent.com/atomcorp/themes/master/app/src/custom-colour-schemes.json | jq -c > themes_custom.json
