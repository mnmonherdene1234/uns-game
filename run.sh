if [ -d "./build" ]; then
  rm -rf ./build
fi

mkdir ./build

cp ./assets ./build/assets -r
go build -o ./build .
build/uns-game.exe
