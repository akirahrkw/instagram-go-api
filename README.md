# instagram-go-api

This is Instagram API library for Go lang

## Environment
	Developed on Go 1.4.1

## How to use

----
#### Configuration

	```go
	var client = instagram.NewClient(func(config *instagram.Config){
		config.ClientId = ""
		config.ClientSecret = ""
		config.RedirectUri = "http://localhost"
	})
	```

#### AuthorizeURL

	```go
	var url = client.AuthorizeURL()
	```
	or
	```go
	var url = client.AuthorizeURLWithScope([]string{"likes","comments","relationships"})
	```

#### RequestToken

	```go
	data, err := client.RequestAccessToken("token")
	```
	or
	```go
	# if you have already access_token
	client.SetAccessToken(ACCESS_TOKEN)
	```

### APIs

----
#### Users

	```go
 	user, content, err := client.Users.Get(userId)

 	user, content, err := client.Users.Self()

	var params = url.Values{}
	params.Set("count", "10")
	items, content, err := client.Users.SelfFeed(params)

	items, content, err := client.Users.RecentMedia(userId, nil)

	items, content, err := client.Users.LikedMedia(nil)

	items, content, err := client.Users.Search("japan", 5)
	```
