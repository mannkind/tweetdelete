go test -v ./...

go test -coverprofile=handlers/.coverprofile github.com/mannkind/tweetdelete/handlers
gover . .coverprofile
go tool cover -html=.coverprofile 
find . -name ".coverprofile" -exec rm {} \;
