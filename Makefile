buf:
	rm -Rf ./backend/proto
	buf generate --path=./proto/
	cd ./backend/proto && go mod init github.com/ueckoken/plarail2023/backend/proto
	cd ./backend/proto && go mod tidy