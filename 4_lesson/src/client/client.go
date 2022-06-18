package client

type Person struct{
	Name string
	Age int8
}

type Avatar struct{
	URL string
	Size int64
}

type Client struct{
	Id int64
	Person
	Img Avatar
}

func (c Client) HasAvatar() bool{
	if c.Img.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar(){
	c.Img.URL = "new_url"
}

// func NewClient(name string, age int, img Avatar) Client{
// 	return Client{
// 		Id: 6,
// 		Person.Name: name,
// 		Age: age,
// 		Img: img,
// 	}
// }