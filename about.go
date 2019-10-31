package vlivego

import (
	"strconv"
	"strings"
)

const aboutURL = "https://api-vfan.vlive.tv/vproxy/channel/{seq}/about?app_id=" + apiID

type aboutJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  aboutResult `json:"result"`
}

type aboutResult struct {
	Comment        string `json:"comment"`
	OpenAt         string `json:"openAt"`
	TagContentList []struct {
		TagSeq  int    `json:"tagSeq"`
		TagName string `json:"tagName"`
	} `json:"tagContentList"`
	PopularCountryList []string `json:"popularCountryList"`
	FanCount           int      `json:"fanCount"`
	VideoCount         int      `json:"videoCount"`
	PostCount          int      `json:"postCount"`
	VideoPlayCount     int      `json:"videoPlayCount"`
	VideoLikeCount     int      `json:"videoLikeCount"`
	VideoCommentCount  int      `json:"videoCommentCount"`
}

func aboutSync(seq int) aboutResult {
	data := aboutJSON{}
	sync(&data, strings.Replace(aboutURL, "{seq}", strconv.Itoa(seq), -1))
	return data.Result
}
