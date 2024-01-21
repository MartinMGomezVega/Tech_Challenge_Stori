git add .
git commit -m "Actualizado"
git push

set GOOS=linux
set GOARCH=amd64

cd cmd
go build -o main main.go
cd ..

del main.zip

tar.exe -a -cf main.zip main

