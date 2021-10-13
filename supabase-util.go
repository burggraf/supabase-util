package main

import (
	"fmt"
	"os"

	// "strings"
	"github.com/supabase/postgrest-go"
)

func main() {
	var (
		SUPABASE_SERVICE_KEY = os.Getenv("SUPABASE_SERVICE_KEY")
		SUPABASE_API_URL     = os.Getenv("SUPABASE_API_URL")
		headers              = map[string]string{"apikey": SUPABASE_SERVICE_KEY}
		schema               = "public"
	)

	client := postgrest.NewClient(SUPABASE_API_URL+"/rest/v1/", schema, headers)
	if client.ClientError != nil {
		panic(client.ClientError)
	}

	res, err := client.From("mytable").Select("*", "", false).Execute()

	// if err := json.Unmarshal(res, &result); err != nil {
	if err != nil {
		fmt.Println("err was")
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println("result was")
	// s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(string(res))

}
