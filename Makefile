gen:
	rm -dRf ./pkg/restapi & rm -dfR ./pkg/models & swagger generate server -f ./tmp/swagger.yaml --target ./pkg --exclude-main

flt:
	swagger flatten -o tmp/swagger.yaml --format=yaml --with-flatten=full pkg/doc/swagger.yml


