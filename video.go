package vlivego

import (
	"math"
	"strconv"
	"strings"
)

const videoURL = "https://api-vfan.vlive.tv/vproxy/channelplus/getChannelVideoList?app_id=" + apiID + "&channelSeq={seq}&maxNumOfRows=100&pageNo={pageNo}"

type videoJSON struct {
	Result struct {
		Total     int           `json:"totalVideoCount"`
		VideoList []videoResult `json:"videoList"`
	} `json:"result"`
}

type videoResult struct {
	VideoSeq                   int    `json:"videoSeq"`
	VideoType                  string `json:"videoType"`
	Title                      string `json:"title"`
	PlayCount                  int    `json:"playCount"`
	LikeCount                  int    `json:"likeCount"`
	CommentCount               int    `json:"commentCount"`
	Thumbnail                  string `json:"thumbnail"`
	PickSortOrder              int    `json:"pickSortOrder"`
	Orientation                string `json:"screenOrientation"`
	WillStartAt                string `json:"willStartAt"`
	WillEndAt                  string `json:"willEndAt"`
	CreatedAt                  string `json:"createdAt"`
	IsUpcoming                 string `json:"upcomingYn"`
	IsSpecialLive              string `json:"specialLiveYn"`
	IsLiveThumb                string `json:"liveThumbYn"`
	ProductID                  string `json:"productId"`
	PackageProductD            string `json:"packageProductId"`
	ProductType                string `json:"productType"`
	PlayTime                   int    `json:"playTime"`
	IsChannelPlusPublic        string `json:"channelPlusPublicYn"`
	ExposeStatus               string `json:"exposeStatus"`
	RepresentChannelName       string `json:"representChannelName"`
	RepresentChannelProfileImg string `json:"representChannelProfileImg"`
	OnAirStartAt               string `json:"onAirStartAt"`
	VliveType                  string `json:"@type"`
}

func videoSync(seq int) (data []videoResult) {
	var run = func(pageNo int) videoJSON {
		data := videoJSON{}
		url := strings.Replace(videoURL, "{seq}", strconv.Itoa(seq), -1)
		url = strings.Replace(url, "{pageNo}", strconv.Itoa(pageNo), -1)
		sync(&data, url)
		return data
	}
	firstRun := run(1)
	count := firstRun.Result.Total
	page := 2
	data = append(data, firstRun.Result.VideoList...)
	for {
		result := run(page)
		data = append(data, result.Result.VideoList...)
		page = page + 1
		if page > int(math.Ceil(float64(float64(count)/100))) || page > 100 {
			break
		}
	}
	return
}
