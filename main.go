package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/thearavind/go-tinyURL/models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var db *sql.DB
var defaultRedirectURL string
var router *gin.Engine

func init() {
	defaultRedirectURL = os.Getenv("DEFAULT_URL")
	db = models.ConnectToDb()
	router = gin.Default()
	router.POST("/short", makeTinyURL)
	router.GET("/:tinyURL", findTinyURL)
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))
}

func main() {
	router.Run(":3000")
}

func makeTinyURL(c *gin.Context) {
	var tiny models.URL
	var urlSearchResult models.URL
	c.BindJSON(&tiny)
	_, err := url.ParseRequestURI(tiny.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest,
			"message": "URL to shorten is not in correct format"})
	} else {
		h := sha1.New()
		io.WriteString(h, tiny.Address)
		tiny.Hash = base64.URLEncoding.EncodeToString(h.Sum(nil))[:6]
		row, err := db.Query(`SELECT * FROM short_url where hash=$1`, tiny.Hash)
		defer row.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError,
				"message": err})
		} else {
			var next = row.Next()
			if !next {
				inserHashToDB(tiny, c)
			} else {
				err := row.Scan(&urlSearchResult.ID, &urlSearchResult.Hash, &urlSearchResult.Address, &urlSearchResult.Clicks)
				switch err {
				case sql.ErrNoRows:
					inserHashToDB(tiny, c)
				case nil:
					c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "tiny_url": urlSearchResult})
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err})
				}
			}
		}
	}
}

func findTinyURL(c *gin.Context) {
	var foundURL models.URL
	hash := c.Param("tinyURL")
	if regexp.MustCompile(`^[a-zA-Z0-9+]{6,6}$`).MatchString(hash) {
		row, err := db.Query(`SELECT * FROM short_url where hash=$1`, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError,
				"message": err})
		} else {
			switch err := row.Scan(&foundURL.ID, &foundURL.Hash, &foundURL.Address, &foundURL.Clicks); err {
			case sql.ErrNoRows:
				c.Redirect(302, defaultRedirectURL)
			case nil:
				_, _ = db.Query(`UPDATE short_url SET clicks = $1 WHERE hash = $2;`, foundURL.Clicks+1, foundURL.Hash)
				c.Redirect(302, foundURL.Address)
			default:
				c.Redirect(302, defaultRedirectURL)
			}
		}
	} else {
		c.Redirect(302, defaultRedirectURL)
	}
}

func redirectHandler(c *gin.Context) {
	c.Redirect(302, defaultRedirectURL)
}

func inserHashToDB(tinyObj models.URL, c *gin.Context) {
	_, err := db.Query(`INSERT INTO short_url(hash, url, clicks)VALUES($1, $2, 0)`,
		tinyObj.Hash, tinyObj.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError,
			"message": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "tiny_url": tinyObj})
	}
}
