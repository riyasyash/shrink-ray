package urlshortner

import ("time"
	"fmt"
)

// ShortenedURL the shrinked url and the key
type ShortenedURL struct {
	Key       string
	URL       string
	Banned    bool
	CreatedAt time.Time
}

func (c *URLShortnerController) shortenURL(url string) (string, error) {
	key := generateUniqueHash(url)
	su := ShortenedURL{
		Key:    key,
		URL:    url,
		Banned: false,
	}
	eKey, err := c.Repo.Exists(url)
	if err != nil {
		fmt.Println("error",err.Error())
		return "", err
	}
	if eKey != "" {
		return eKey, nil
	}
	err = c.Repo.Insert(su)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (c *URLShortnerController) getURL(key string) (string, error) {
	return c.Repo.Find(key)
}
