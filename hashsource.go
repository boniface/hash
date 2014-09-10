package main

import (
	"fmt"
	"log"
	"github.com/SlyMarbo/rss"
	"crypto/md5"
	"encoding/hex"
	"github.com/gocql/gocql"
	"time"
	"github.com/advancedlogic/GoOse"
	"strings"
	"regexp"
)

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "hashmedia"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	fmt.Printf("Reading Feeds from the Database!")
	getGetLinks(session)
	getContent(session)
	updateZonePosts(session)
	updateSitePosts(session)
	fmt.Println(time.Now())
}



func getGetLinks(session *gocql.Session) {
	feeds := session.Query("SELECT zone,feedLink,sitecode FROM feeds").Iter()
	var feedLink string
	var zone string
	var siteCode string
	for feeds.Scan(&zone, &feedLink, &siteCode) {
		feed, err := rss.Fetch(feedLink)
		if err != nil {
			// handle error.
		}
		var links = feed.Items
		for _, link := range links {

			if err := session.Query(`INSERT INTO links  (zone,linkhash,datepublished,site,url,sitecode) VALUES (?,?,?,?,?,?)`,
				zone, GetMD5Hash(link.ID), link.Date, link.ID, link.Link, siteCode).Exec(); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := feeds.Close(); err != nil {
		log.Fatal(err)
	}
}

func getContent(session *gocql.Session) {
	links := session.Query("SELECT url,zone,datepublished,sitecode FROM links").Iter()
	var url, zone, siteCode string
	var datepublished time.Time
	for links.Scan(&url, &zone, &datepublished, &siteCode) {
		duration := time.Now().Sub(datepublished).Minutes()
		if (duration < 600000) {
			LoadContent(zone, url, datepublished, session, siteCode)
		}
	}
	if err := links.Close(); err != nil {
		log.Fatal(err)
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}


func updateZonePosts(session *gocql.Session) {
	posts := session.Query("SELECT zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title,sitecode FROM posts").Iter()
	var zone, linkhash, domain, article, caption, imagepath, imageurl, link, metadescription, metakeywords, seo, title, sitecode string
	var date time.Time

	for posts.Scan(&zone, &linkhash, &domain, &date, &article, &caption, &imagepath, &imageurl, &link, &metadescription, &metakeywords, &seo, &title, &sitecode) {

		if err := session.Query(`INSERT INTO zoneposts (zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title,sitecode)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			zone, linkhash, domain, date, article, caption, imagepath, imageurl, link, metadescription, metakeywords, seo, title, sitecode).Exec();
			err != nil {
			log.Fatal(err)
			println(err)
		}
	}
	if err := posts.Close(); err != nil {
		log.Fatal(err)
	}



}

func updateSitePosts(session *gocql.Session) {
	posts := session.Query("SELECT zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title,sitecode FROM posts").Iter()
	var zone, linkhash, domain, article, caption, imagepath, imageurl, link, metadescription, metakeywords, seo, title, sitecode string
	var date time.Time

	for posts.Scan(&zone, &linkhash, &domain, &date, &article, &caption, &imagepath, &imageurl, &link, &metadescription, &metakeywords, &seo, &title, &sitecode) {

		if err := session.Query(`INSERT INTO siteposts (zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title,sitecode)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			zone, linkhash, domain, date, article, caption, imagepath, imageurl, link, metadescription, metakeywords, seo, title, sitecode).Exec();
			err != nil {
			log.Fatal(err)
			println(err)
		}
	}
	if err := posts.Close(); err != nil {
		log.Fatal(err)
	}


}

func LoadContent(zone string, url string, datepublished time.Time, session *gocql.Session, siteCode string) {
	g := goose.New()
	article := g.ExtractFromUrl(url)

	if err := session.Query(`INSERT INTO posts (zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title,sitecode)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		zone, GetMD5Hash(article.FinalUrl), article.Domain, datepublished, filterContent(article.CleanedText),
		"caption", "imagepath ", article.TopImage, article.CanonicalLink, article.MetaDescription, article.MetaKeywords, prettyUrl(article.Title), article.Title, siteCode).Exec();
		err != nil {
		log.Fatal(err)
		println(err)
	}

	if err := session.Query(`INSERT INTO rawposts (linkhash,zone,datepublished,rawhtml) VALUES (?,?,?,?)`, GetMD5Hash(article.FinalUrl), zone, datepublished, article.RawHtml).Exec();
		err != nil {
		log.Fatal(err)
		println(err)
	}
}


func prettyUrl(title string ) string {
	//let's make pretty urls from title
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	regStoWords := regexp.MustCompile("to|old|the|is|do|at|be|was|were|am|I|says|say|of")
	if err != nil {
		log.Fatal(err)
	}
	cleanTitle:=regStoWords.ReplaceAllLiteralString(title,"")
	prettyurl := reg.ReplaceAllString(cleanTitle, "-")
	prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))
	return prettyurl
}

func filterContent(input string) string{
	wordsRegExp := regexp.MustCompile("Main News|Editor's Choice|Breaking News|More News|Contact us|Filed under|Home / Breaking News /|TweetEmail|Related Posts")
	spaceRegExp :=regexp.MustCompile(`\t|\n`)
	paraRegExp :=regexp.MustCompile(`\.`)
	results:=wordsRegExp.ReplaceAllString(input, "")
	result:=spaceRegExp.ReplaceAllString(results, "")
	res:=paraRegExp.ReplaceAllString(result, ".\n")
	return res
}

func metaDescription(s string,pos,length int) string{
	runes:=[]rune(s)
	l := pos+length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

