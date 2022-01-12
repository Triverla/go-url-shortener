package logic

import (
	"github.com/Triverla/go-url-shortener/logic"
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := logic.GenerateShortLink(initialLink1, UserId)

	initialLink2 := "https://github.com/Triverla"
	shortLink2 := logic.GenerateShortLink(initialLink2, UserId)

	initialLink3 := "https://github.com/Triverla/url-shortener"
	shortLink3 := logic.GenerateShortLink(initialLink3, UserId)

	assert.Equal(t, shortLink1, "jTa4L57P")
	assert.Equal(t, shortLink2, "QwJwA2pU")
	assert.Equal(t, shortLink3, "UDDPGqni")
}
