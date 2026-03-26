#bin/bash

# У тебя есть файл servers.txt, где в каждой строке написан IP-адрес. Тебе нужно написать скрипт, который пройдет по всем адресам и проверит, отвечает ли там HTTP-сервер (код 200). Результат нужно записать в файл status.log.

for ip in $(cat servers.txt); do
    code=$(curl -s -o /dev/null -w "%{http_code}" "http://$ip")
    if [ $code == "200" ]; then
        echo "$ip is OK" >> status.log
    else
        echo "$ip is DOWN (code: $code)" >> status.log
    fi
done