GOOS=windows GOARCH=amd64 go build -o bin/supabase-util-windows-amd64.exe supabase-util.go
GOOS=windows GOARCH=386 go build -o bin/supabase-util-windows-386.exe supabase-util.go
GOOS=windows GOARCH=arm go build -o bin/supabase-util-windows-arm.exe supabase-util.go
GOOS=windows GOARCH=arm64 go build -o bin/supabase-util-windows-arm64.exe supabase-util.go

