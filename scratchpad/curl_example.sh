#!/usr/bin/env bash

ITEM_NAME="KIKKOMAN%20PLUM%20WINE%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20750ml"
ITEM_ID="812000"

curl "https://webapps2.abc.utah.gov/Production/OnlineInventoryQuery/IQ/InventoryQuery.aspx" \
  --user-agent "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36" \
  --referer "https://webapps2.abc.utah.gov/Production/OnlineInventoryQuery/IQ/InventoryQuery.aspx" \
  -H "Origin: https://webapps2.abc.utah.gov" \
  -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" \
  -H "Cache-Control: no-cache" \
  -H "X-MicrosoftAjax: Delta=true" \
  --data "__VIEWSTATE=%2FwEPDwUKMjAyODAwNTUyNQ9kFgJmD2QWAgIDD2QWAgIFD2QWAgIDD2QWAgIBD2QWAmYPZBYGAgMPZBYCAgUPDxYCHgRUZXh0BQ9WZXJzaW9uIEFYIDczNjBkZAIFD2QWBgIRDw8WAh8AZWRkAhMPDxYCHwBlZGQCFw9kFgICAQ88KwARAgEQFgAWABYADBQrAABkAgwPZBYCAgEPZBYCAhMPZBYCAgEPZBYCAgEPZBYCAgcPPCsAEQEMFCsAAGQYAgUqY3RsMDAkQ29udGVudFBsYWNlSG9sZGVyQm9keSRlcnJvckdyaWR2aWV3D2dkBS9jdGwwMCRDb250ZW50UGxhY2VIb2xkZXJCb2R5JGd2SW52ZW50b3J5RGV0YWlscw9nZLdNSQsBSLRAk%2FanooHUpl68FAUnxxSf41DGVh4dX1h2" \
  --data "__EVENTVALIDATION=%2FwEdAAty3%2BUj%2BQkt%2BReNgIPmvfXx6NJMDzww3IJozuamiH2ntIYHsMIzYYJ%2BhEPXbUswhEgM%2BQi2wto3oIoOmOtZDQewMu1wDE8XlC69lmZmCD%2BHfJNQBJwY3SeA8fU%2FnTO3Cgd0dXDvR%2FgB01%2FGD3JYF0vk2B395ZvQbXFGN04uuEIz14KTWmUjUiEBImqToeuQtKrqZqEd5e1TGujvVLTXWKyhCPvlvwBboFdZ%2BtYjLeo3N1RBqbWpo2fhkoM93Eh7GiLKeEue2dhSZVMcVvQunURs" \
  --data "ctl00%24ContentPlaceHolderBody%24tbItemName=${ITEM_NAME}" \
  --data "ctl00%24ContentPlaceHolderBody%24hiddenItemId=${ITEM_ID}" \
  --compressed

