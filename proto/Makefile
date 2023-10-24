buf:
	rm -Rf ../backend/spec
	buf generate
	mkdir -p ../backend/spec
	cd ../backend/spec && go mod init github.com/ueckoken/plarail2023/backend/spec
	cd ../backend/spec && go mod tidy