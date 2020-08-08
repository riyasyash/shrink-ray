package urlshortner

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/riyasyash/shrink-ray/api"
)

type URLShortnerController struct {
	Repo *URLShortnerRepo
}

type URLShortenRequest struct {
	URL string
}

func New(Db *sql.DB) *URLShortnerController {
	return &URLShortnerController{
		Repo: &URLShortnerRepo{
			Db: Db,
		},
	}
}

func (c *URLShortnerController) Shorten(w http.ResponseWriter, r *http.Request) {
	var req URLShortenRequest
	if err := api.Decode(w, r, &req); err != nil {
		return
	}
	url, err := c.shortenURL(req.URL)
	if err != nil {
		api.Respond(200, w, err)
	}
	api.Respond(200, w, url)
}

func (c *URLShortnerController) Redirect(w http.ResponseWriter, r *http.Request) {
	url, err := c.getURL(r.RequestURI[1:])
	if err != nil {
		fmt.Println(err)
		api.Redirect(302, w, "/error")
	}
	if url != "" {
		api.Redirect(302, w, url)
	} else {
		api.Redirect(302, w, "/")
	}
}
