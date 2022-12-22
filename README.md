# Proto (Место хранения всех прото файлов)

команда для генерации кода

    protoc -I=../src/ --go_out=. --go-grpc_out=. ../src/*/*.proto

    protoc --proto_path=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf:. --gogofast_out=. *.proto
    protoc -I=../src/ --gogofast_out=. ../src/*/*.proto
