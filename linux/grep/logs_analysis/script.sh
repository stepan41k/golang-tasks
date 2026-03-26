#bin/sh

grep " 404" access.log | awk '{print $7}' | sort | uniq -c | sort -rn | head -n 3