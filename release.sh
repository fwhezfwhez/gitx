echo -n 'tag'
read tag

# windows 64
GOOS=windows GOARCH=amd64 go build -o gitx.exe
mkdir -p binary/${tag}/windows
mv gitx.exe binary/${tag}/windows/gitx.exe

# darwin 64
GOOS=darwin GOARCH=amd64 go build -o gitx
mkdir -p binary/${tag}/darwin
mv gitx binary/${tag}/darwin/gitx

# linux 64
GOOS=linux GOARCH=amd64 go build -o gitx
mkdir -p binary/${tag}/linux
mv gitx binary/${tag}/linux/gitx