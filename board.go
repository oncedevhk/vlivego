package vlivego

import (
	"math"
	"strconv"
	"strings"
)

const boardURL = "https://api-vfan.vlive.tv/v3/board.{seq}/posts?locale=en&app_id=" + apiID + "&limit=100"

type boardJSON struct {
	Data  []boardResult `json:"data"`
	Total int           `json:"total_count"`
}

type boardResult struct {
	Author struct {
		IsChannelPlus bool   `json:"is_channel_plus"`
		Level         int    `json:"level"`
		Nickname      string `json:"nickname"`
		ProfileImages string `json:"profile_image"`
		Role          string `json:"role"`
		UserSeq       int    `json:"user_seq"`
		VNumber       string `json:"v_number"`
	} `json:"author"`
	Category       string `json:"category"`
	CommentCount   int    `json:"comment_count"`
	Content        string `json:"content"`
	ContentVersion string `json:"content_version"`
	CreatedAt      int    `json:"created_at"`
	EmotionCount   int    `json:"emotion_count"`
	Images         []struct {
		Height int    `json:"height"`
		Thumb  string `json:"thumb"`
		Type   string `json:"type"`
		Weight int    `json:"width"`
	} `json:"image_list"`
	IsBest       bool   `json:"is_best"`
	IsEvent      bool   `json:"is_event"`
	IsRestricted bool   `json:"is_restricted"`
	IsVisible    bool   `json:"is_visible_to_authorized_users"`
	NoticeStatus bool   `json:"notice_status"`
	PostID       string `json:"post_id"`
	Title        string `json:"title"`
	WrittenIn    string `json:"written_in"`
}

func boardSync(seq int) (data []boardResult) {
	var run = func(pageNo int) boardJSON {
		data := boardJSON{}
		url := strings.Replace(boardURL, "{seq}", strconv.Itoa(seq), -1)
		url = strings.Replace(url, "{pageNo}", strconv.Itoa(pageNo), -1)
		sync(&data, url)
		return data
	}
	firstRun := run(1)
	count := firstRun.Total
	page := 2
	data = append(data, firstRun.Data...)
	for {
		result := run(page)
		data = append(data, result.Data...)
		page = page + 1
		if page > int(math.Ceil(float64(float64(count)/100))) || page > 100 {
			break
		}
	}
	return
}
