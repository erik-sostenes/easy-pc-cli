#!/bin/bash

file='urls.txt'
urls=()
while read -r url || [ -n "$new" ]; do
     urls+=("$url")
done < "$file"

id=$(uuidgen)

cd ./cmd/
./cmd.exe offers -website=MercadoLibre  -category-id="$id" -category=ekectro "https://www.mercadolibre.com.mx/ofertas?container_id=MLM779363-1&category=MLM1000&page=9c"