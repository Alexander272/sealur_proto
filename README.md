# Proto (Место хранения всех прото файлов)

команда для генерации кода

    protoc -I=../src/ --go_out=. --go-grpc_out=. ../src/*/*.proto