echo "building server..."
BUILD_PATH=/go/SMTP
docker run --rm -v $PWD:$BUILD_PATH golang:1.8 go build -o $BUILD_PATH/server /SMTP

echo 'building image...'
docker build -t "smtp" .
