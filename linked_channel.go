package vlivego

import (
	"strconv"
	"strings"
)

var linkedChannelFields = []string{"channel_seq",
	"channel_code",
	"type",
	"channel_name",
	"comment",
	"fan_count",
	"channel_cover_img",
	"channel_profile_img",
	"representative_color",
	"background_color"}

var linkedChannelURL = "https://api-vfan.vlive.tv/channel.{seq}/linked-channels?locale=en&app_id=" + apiID + "&fields=" + strings.Join(linkedChannelFields, ",")

type linkedChannelJSON struct {
	Data []linkedChannelResult `json:"data"`
}

type linkedChannelResult struct {
	ChannelSeq          int    `json:"channel_seq"`
	ChannelCode         string `json:"channel_code"`
	Type                string `json:"type"`
	ChannelName         string `json:"channel_name"`
	Comment             string `json:"comment"`
	FanCount            int    `json:"fan_count"`
	CoverImage          string `json:"channel_cover_img"`
	ProfileImage        string `json:"channel_profile_img"`
	RepresentativeColor string `json:"representative_color"`
	BackgroundColor     string `json:"background_color"`
	HasFanClub          bool   `json:"fanclub"`
}

func linkChannelSync(seq int) []linkedChannelResult {
	data := linkedChannelJSON{}
	sync(&data, strings.Replace(linkedChannelURL, "{seq}", strconv.Itoa(seq), -1))
	return data.Data
}
