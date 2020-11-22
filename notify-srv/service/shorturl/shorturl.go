package shorturl

import (
	"notify/service"
)

type ShortUrlService struct {
	Base *service.Base
}

func (this *ShortUrlService) Handle(longUrl string) string {
	return "HELLO WORLD" + longUrl
}
