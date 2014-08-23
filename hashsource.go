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
)

func main() {

	fmt.Printf("Reading Feeds from the Database!")
	getGetLinks()
	getContent()
	fmt.Println(time.Now())
}



func getGetLinks() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "hashmedia"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()


	feeds := session.Query("SELECT zone,feedLink FROM feeds").Iter()
	var feedLink string
	var zone string
	for feeds.Scan(&zone,&feedLink) {
		feed, err := rss.Fetch(feedLink)
		if err != nil {
			// handle error.
		}
		var links = feed.Items
		for _, item := range links {

			if err := session.Query(`INSERT INTO links  (zone,linkhash,datepublished,site,url) VALUES (?,?,?,?,?)`,
				zone,GetMD5Hash(item.ID), item.Date, item.ID, item.Link).Exec(); err != nil {
				log.Fatal(err)

			}

		}
	}
	if err := feeds.Close(); err != nil {
		log.Fatal(err)
	}
}

func getContent() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "hashmedia"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()
	links := session.Query("SELECT url,zone,datepublished FROM links").Iter()
	var url string
	var zone string
	var datepublished time.Time
	for links.Scan(&url, &zone, &datepublished) {
		duration:=time.Now().Sub(datepublished).Minutes()
		if(duration < 60){
			LoadContent(zone,url,datepublished, session)
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

func LoadContent(zone string,url string, datepublished time.Time,session *gocql.Session) {
	g := goose.New()
	article := g.ExtractFromUrl(url)

	if err := session.Query(`INSERT INTO posts (zone,linkhash,domain,date,article,caption,imagepath,imageurl,link,metadescription,metakeywords,seo,title)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		zone,GetMD5Hash(article.FinalUrl), article.Domain, datepublished, article.CleanedText,
		"caption","imagepath ",article.TopImage,article.CanonicalLink,article.MetaDescription, article.MetaKeywords,"SEO", article.Title).Exec();
		err != nil {
		log.Fatal(err)
		println(err)
	}

	if err := session.Query(`INSERT INTO rawposts (linkhash,zone,datepublished,rawhtml) VALUES (?,?,?,?)`, GetMD5Hash(article.FinalUrl),zone,datepublished, article.RawHtml).Exec();
		err != nil {
		log.Fatal(err)
		println(err)
	}

}
