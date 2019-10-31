package vlivego

// Client to refresh and store all vlive channel information
type Client struct {
	ID       int                   `json:"seq"`
	Base     channelJSON           `json:"base"`
	About    aboutResult           `json:"about"`
	Channels []linkedChannelResult `json:"channels"`
	Videos   []videoResult         `json:"videos"`
	Posts    []clientPost          `json:"posts"`
	FanPosts []clientPost          `json:"fan_posts"`
}

type clientPost struct {
	ID    int           `json:"board_id"`
	Posts []boardResult `json:"posts"`
}

/*
NewClient to create a client to get vlive channel information by channel seq (number)
	Example:
	client := New Client(6)
TWICE: 6,
DAY6: 164,
Stray Kids: 633,
ITZY: 1093
*/
func NewClient(seq int) (c *Client) {
	c = &Client{}
	c.ID = seq
	return
}

// RefreshAll to refresh all columns in client without RefreshFanPosts
func (c *Client) RefreshAll() {
	c.RefreshBase()
	c.RefreshAbout()
	c.RefreshChannels()
	c.RefreshVideos()
	c.RefreshPosts()
	// c.RefreshFanPosts()
}

// RefreshBase to refresh channel base information
func (c *Client) RefreshBase() {
	c.Base = channelSync(c.ID)
}

// RefreshAbout to refresh channel about information
func (c *Client) RefreshAbout() {
	c.About = aboutSync(c.ID)
}

// RefreshChannels to refresh sub channels information that linked to this channel
func (c *Client) RefreshChannels() {
	result := linkChannelSync(c.ID)
	if len(result) > 0 {
		c.Channels = result
	}
}

// RefreshVideos to refresh last 10000 channel video information
func (c *Client) RefreshVideos() {
	result := videoSync(c.ID)
	if len(result) > 0 {
		c.Videos = result
	}
}

/*
RefreshPosts to refresh last 10000 channel post information

Should run RefreshBase at first if not run before
*/
func (c *Client) RefreshPosts() {
	var posts []clientPost
	for _, board := range c.Base.CelebBoards {
		data := clientPost{}
		result := boardSync(board.BoardID)
		data.ID = board.BoardID
		if len(result) > 0 {
			data.Posts = result
		}
		posts = append(posts, data)
	}
	if len(posts) > 0 {
		c.Posts = posts
	}
}

/*
RefreshFanPosts to refresh last 10000 channel fan post information

Should run RefreshBase at first if not run before

About 39.825s for 10000 channel fan post testing with TWICE Channel Fan Board 22
*/
func (c *Client) RefreshFanPosts() {
	var posts []clientPost
	for _, board := range c.Base.FanBoards {
		data := clientPost{}
		result := boardSync(board.BoardID)
		data.ID = board.BoardID
		if len(result) > 0 {
			data.Posts = result
		}
		posts = append(posts, data)
	}
	if len(posts) > 0 {
		c.FanPosts = posts
	}
}

/*
GetChannelSeq to get channel seq for using client by channel code

Example: https://channels.vlive.tv/EDBF => EDBF is the channel code
*/
func GetChannelSeq(code string) (seq int) {
	var data = struct {
		Result struct {
			Seq  int    `json:"channelSeq"`
			Code string `json:"channelCode"`
		} `json:"result"`
	}{}
	url := "http://api.vfan.vlive.tv/vproxy/channelplus/decodeChannelCode?app_id=" + apiID + "&channelCode=" + code
	sync(&data, url)
	seq = data.Result.Seq
	return
}
