package vlivego

import (
	"strconv"
	"strings"
)

var channelFields = []string{"channel_seq",
	"channel_code",
	"type",
	"channel_name",
	"comment",
	"fan_count",
	"channel_cover_img",
	"channel_profile_img",
	"representative_color",
	"background_color",
	"celeb_boards",
	"fan_boards",
	"is_show_banner",
	"vstore",
	"is_show_upcoming",
	"media_channel",
	"gfp_ad_enabled",
	"banner_ad_enabled",
	"ad_channel_id",
	"ad_cp_id",
	"fanclub",
	"agency_seq"}

var channelURL = "https://api-vfan.vlive.tv/v2/channel.{seq}?locale=en&app_id=" + apiID + "&fields=" + strings.Join(channelFields, ",")

type channelJSON struct {
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
	CelebBoards         []struct {
		BoardID        int  `json:"board_id"`
		SpecialOngoing bool `json:"special_ongoing"`
	} `json:"celeb_boards"`
	FanBoards []struct {
		BoardID        int  `json:"board_id"`
		SpecialOngoing bool `json:"special_ongoing"`
	} `json:"fan_boards"`
	IsShowBanner bool `json:"is_show_banner"`
	VStore       struct {
		VStoreLink string `json:"vstore_home_link"`
		VStoreSeq  int    `json:"vstore_seq"`
	} `json:"vstore"`
	IsShowUpComing  bool   `json:"is_show_upcoming"`
	HasMediaChannel bool   `json:"media_channel"`
	GFPADEnabled    bool   `json:"gfp_ad_enabled"`
	BannerADEnabled bool   `json:"banner_ad_enabled"`
	ADChannelID     string `json:"ad_channel_id"`
	ADCPID          string `json:"ad_cp_id"`
	HasFanClub      bool   `json:"fanclub"`
	AgencySeq       int    `json:"agency_seq"`
}

func channelSync(seq int) channelJSON {
	data := channelJSON{}
	sync(&data, strings.Replace(channelURL, "{seq}", strconv.Itoa(seq), -1))
	return data
}
