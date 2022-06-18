cat words.txt |tr -s ' ' '\n' |sort|uniq -c|sort -r|awk '{print $2,$1}'

cat file.txt | grep -P "^(\([0-9]{3}\)\s|[0-9]{3}-)[0-9]{3}-[0-9]{4}$"

COUNT=`head -1 file2.txt | wc -w`
for (( i = 1; i <= $COUNT; i++ )); do
cut -d' ' -f$i file.txt | xargs
done

COUNT=`head -1 file2.txt | wc -w`
for (( i = 1; i <= $COUNT; i++ )); do
awk -v arg=$i '{print $arg}' file.txt | xargs
done

sed -n "10p" file3.txt
