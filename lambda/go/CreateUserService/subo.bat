git add .
git commit -m "Actualizado"
git push

set GOOS=linux
set GOARCH=amd64

cd cmd
go build -o main main.go
cd ..

del main.zip

"C:\Program Files\Git\bin\tar.exe" -a -cf main.zip cmd/main

