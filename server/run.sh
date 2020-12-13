#! /bin/bash
echo -e "\e[1;31m 1. Initializing databases \e[0m"
while :; do
    if go run script/migrate.go; then
        break
    fi
    sleep 1
done

echo -e "\e[1;31m 2. Server start \e[0m"
go build -o server . && ./server
echo -e "\e[1;31m 3. Finished! \e[0m"
