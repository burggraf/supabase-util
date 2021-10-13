GOOS=linux GOARCH=386 go build -o bin/supabase-util-linux-386 supabase-util.go
GOOS=linux GOARCH=amd64 go build -o bin/supabase-util-linux-amd64 supabase-util.go
GOOS=linux GOARCH=arm go build -o bin/supabase-util-linux-arm supabase-util.go
GOOS=linux GOARCH=arm64 go build -o bin/supabase-util-linux-arm64 supabase-util.go
