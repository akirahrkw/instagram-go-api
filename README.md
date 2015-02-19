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
data, err := client.RequestAccessToken(token)
```

or, if you have already access token
```go
client.SetAccessToken(accessToken)
```

### APIs

----
#### Return Values
in most cases, API method returns three values
1: The data you want from API (the value of "data")
2: Whole Json data including meta, pagination, data
3: error

----
#### Users
http://instagram.com/developer/endpoints/users/

```go
user, content, err := client.Users.Get(userId)

user, content, err := client.Users.Self()

var params = url.Values{}
items, content, err := client.Users.SelfFeed(params)

items, content, err := client.Users.RecentMedia(userId, params)

items, content, err := client.Users.LikedMedia(params)

items, content, err := client.Users.Search(query, 5)
```

---
#### Media
http://instagram.com/developer/endpoints/media/

```go
media, content, err := client.Media.Get(mediaId)

media, content, err := client.Media.Shortcode(shortcode)

items, content, err := client.Media.Popular()

var params = url.Values{}
items, content, err := client.Media.Search(params)
```

---
#### Likes
http://instagram.com/developer/endpoints/likes/

```go
items, content, err := client.Likes.Get(mediaId)

content, err := client.Likes.Post(mediaId)

content, err := client.Likes.Del(mediaId)
```

---
#### Relationships
http://instagram.com/developer/endpoints/relationships/

```go
items, content, err := client.Relationships.Follows(userId)

items, content, err := client.Relationships.FollowedBy(userId)

items, content, err := client.Relationships.RequestedBy()

relationship, content, err := client.Relationships.GetRelationship(userId)

relationship, content, err := client.Relationships.PostRelationship(userId, "follow")

relationship, content, err = client.Relationships.PostRelationship(userId, "unfollow")
```

---
#### Comments
http://instagram.com/developer/endpoints/comments/

```go
items, content, err := client.Comments.Get(mediaId)

content, err := client.Comments.Post(mediaId, text)

content, err := client.Comments.Get(mediaId, commentId)
```

---
#### Tags
http://instagram.com/developer/endpoints/tags/

```go
tag, content, err := client.Tags.Get(tagName)

items, content, err := client.Tags.RecentMedia(tagName, nil)

var params = url.Values{}
items, content, err := client.Tags.Search(params)
```

---
#### Locations
http://instagram.com/developer/endpoints/locations/

```go
location, content, err := client.Locations.Get(locationId)

var params = url.Values{}
items, content, err := client.Locations.RecentMedia(locationId, params)

items, content, err := client.Locations.Search(params)
```

---
#### Geography
http://instagram.com/developer/endpoints/geographies/\

```go
var params = url.Values{}
location, content, err := client.Geographies.RecentMedia(geoId, params)
```
