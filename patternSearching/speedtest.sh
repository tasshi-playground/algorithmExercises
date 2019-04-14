echo "naive"
time go run naive/main.go `cat neko.txt` 彼等鈍瞎漢どんかつかん > /dev/null

echo "\nimproved"
time go run main.go `cat neko.txt` 彼等鈍瞎漢どんかつかん > /dev/null

