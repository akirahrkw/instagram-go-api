package instagram

// error
type instagramError struct {
	Message string
	Status	int
}
func (e *instagramError) Error() string {
	return e.Message
}

// Base
type Content struct {
	Meta  *Meta  `json:"meta,omitempty"`
	Data  interface{}	 `json:"data,omitempty"`
 	Pagination *Pagination  `json:"pagination,omitempty"`
}

type Meta struct {
	Code  int `json:"code"`
	ErrorType  string `json:"error_type,omitempty"`
	ErrorMessage  string `json:"error_message,omitempty"`
}

type Pagination struct {
	NextUrl  string  `json:"next_url,omitempty"`
	NextMaxId  string  `json:"next_max_id,omitempty"`
}

type User struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	FullName string `json:"full_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Bio string `json:"bio,omitempty"`
	Website string `json:"website,omitempty"`
	Counts *UserCount `json:"counts,omitempty"`
	Type string `json:"type,omitempty"`
}

type UserCount struct {
	Media int `json:"media,omitempty"`
	Follows int `json:"follows,omitempty"`
	FollowedBy int `json:"followed_by,omitempty"`
}

type Auth struct {
	//error
	Code int `json:"code,omitempty"`
	ErrorType string `json:"error_type,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	AccessToken	string `json:"access_token,omitempty"`
	User *User `json:"user,omitempty"`
}

type Image struct {
	Url string `json:"url,omitempty"`
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Images struct {
	LowResolution *Image `json:"low_resolution,omitempty"`
  Thumbnail *Image `json:"thumbnail,omitempty"`
  StandardResolution *Image `json:"standard_resolution,omitempty"`
}

type Video struct {
	Url string `json:"url,omitempty"`
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Videos struct {
	LowResolution *Image `json:"low_resolution,omitempty"`
  StandardResolution *Image `json:"standard_resolution,omitempty"`
}

type Likes struct {
	Count int `json:"count,omitempty"`
	Data  []User `json:"data,omitempty"`
}

type Comment struct {
	Id string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	CreatedTime string `json:"created_time,omitempty"`
	From *User `json:"from,omitempty"`
}
type Comments struct {
	Count int `json:"count,omitempty"`
	Data 	[]Comment `json:"data,omitempty"`
}

type Location struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Media struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Filter string `json:"filter,omitempty"`
	Link string `json:"link,omitempty"`
	CreatedTime string `json:"created_time,omitempty"`
	Tags []string `json:"tags,omitempty"`
	User *User `json:"user,omitempty"`
	Images *Images `json:"images,omitempty"`
	Caption *Comment `json:"caption,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type Relationship struct {
	OutgoingStatus string `json:"outgoing_status,omitempty"`
	IncomingStatus string `json:"incoming_status,omitempty"`
	TargetUserIsPrivate bool `json:"target_user_is_private,omitempty"`
}

type Tag struct {
	MediaCount int `json:"media_count,omitempty"`
	Name string `json:"name,omitempty"`
}
