protoc -I ./proto --go_out=plugins=grpc:./vendor `find ./proto/com -name \*.proto | cut -c3-`

#protoc -I ./proto -I ./schema-proto --go_out=plugins=grpc:./schema `find ./schema-proto -name \*.proto | cut -c3-`