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
	"github.com/garyburd/redigo/redis"
	"bytes"
)

const (
	ADDRESS = "127.0.0.1:6379"
)

var (
	c, err = redis.Dial("tcp", ADDRESS)
)

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "hashmedia"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	fmt.Println("Reading Feeds from the Database Start Time!", time.Now())
	getGetLinks(session)
	getContent(session)
	updateZonePosts(session)
	updateSitePosts(session)
	fmt.Println("Process Complete at ", time.Now())
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
				zone, GetMD5Hash(link.ID), link.Date, link.Link, link.ID, siteCode).Exec(); err != nil {
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
		if (duration < 60) {
			LoadContent(zone, url, datepublished, session, siteCode)
		}
	}
	if err := links.Close(); err != nil {
		log.Fatal("Error in Loading Content",err)
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
			log.Fatal("The Key was Empty:",err)
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
		"caption", "imagepath ", article.TopImage, article.CanonicalLink, metaDescription(filterContent(article.CleanedText), 0, 200), getKeyWords(article.Title), prettyUrl(article.Title), article.Title, siteCode).Exec();
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

func prettyUrl(title string) string {
	//let's make pretty urls from title
	reg, err := regexp.Compile("[^A-Za-z0-9]+")

	if err != nil {
		log.Fatal("Error generating Pretty URL",err)
	}
	cleanTitle := RemoveStopWords(GetTerms(title))
	prettyurl := reg.ReplaceAllString(cleanTitle, "-")
	prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))
	return prettyurl
}

func filterContent(input string) string {
	wordsRegExp := regexp.MustCompile("Main News|Home|You are here: /|Posted in|ZNBC User|Featured||About usOrg StructureHistoryRate CardsProducts & ServicesOnline SubscriptionsSelect a PageAbout us- Org Structure- History- Rate CardsProducts & Services- Online SubscriptionsAbout usOrg StructureHistoryRate CardsProducts & ServicesOnline SubscriptionsSelect a PageAbout us- Org Structure- History- Rate CardsProducts & Services- Online SubscriptionsLatest NewsStoriesCourt NewsBusinessStoriesMoney/Stock ExchangeColumnsLetters to the EditorEntertainmentMusicTheatreFilmsOthersColumnsFeaturesOpinionSportsStoriesFootballRugbyBoxingVolleyballColumnsOthersSelect a PageLatest News- Stories- Court NewsBusiness- Stories- Money/Stock Exchange- ColumnsLetters to the EditorEntertainment- Music- Theatre- Films- Others- ColumnsFeaturesOpinionSports- Stories- Football- Rugby- Boxing- Volleyball- Columns- OthersLatest NewsStoriesCourt NewsBusinessStoriesMoney/Stock ExchangeColumnsLetters to the EditorEntertainmentMusicTheatreFilmsOthersColumnsFeaturesOpinionSportsStoriesFootballRugbyBoxingVolleyballColumnsOthersSelect a PageLatest News- Stories- Court NewsBusiness- Stories- Money/Stock Exchange- ColumnsLetters to the EditorEntertainment- Music- Theatre- Films- Others- ColumnsFeaturesOpinionSports- Stories- Football- Rugby- Boxing- Volleyball- Columns- Others HOME SLIDE SHOW, SHOWCASE Register to vote!|About us|News|Headlines|Statements|Advertise|Posted by|Editor's Choice|Breaking News|More News|Contact us|Filed under|Home / Breaking News /|TweetEmail|Related Posts")
	spaceRegExp := regexp.MustCompile(`\t|\n`)
	paraRegExp := regexp.MustCompile(`\.`)
	results := wordsRegExp.ReplaceAllString(input, "")
	result := spaceRegExp.ReplaceAllString(results, "")
	res := paraRegExp.ReplaceAllString(result, ".\n")
	return res
}

func metaDescription(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	results := string(runes[pos:l])
	return strings.Title(results)
}

func getKeyWords(title string) string {
	regSpace := regexp.MustCompile(`\s`)
	cleanTitle := RemoveStopWords(GetTerms(title))
	shortenTitle := metaDescription(cleanTitle, 0, 70)
	keyWords := regSpace.ReplaceAllString(shortenTitle, ",")
	return keyWords
}


func GetTerms(sentence string) [] string {
	terms := strings.Split(string(sentence), " ")
	return terms;
}

func RemoveStopWords(input []string) string {
	var buffer bytes.Buffer
	if err != nil {
		log.Fatal("Fatal Error Occured",err)
	}
	c.Send("MULTI")
	c.Send("DEL", "inputWords")
	c.Send("SADD", redis.Args{}.Add("inputWords").AddFlat(input)...)
	c.Send("SDIFF", "inputWords", "stopwords")
	reply, err := c.Do("EXEC")
	if err != nil {
		fmt.Println("Error Executing Commands", err)
	}
	values, _ := redis.Values(reply, nil)
	fliteredWords, err := redis.Strings(values[2], nil)
	if err != nil {
		fmt.Println("Wrong Type Received", err)
	}
	if (len(fliteredWords)) > 0 {
		for _, v := range fliteredWords {
			buffer.WriteString(v + " ")
		}
	} else {
		fmt.Println(">>Nothing found")
		for _, v := range input {
			buffer.WriteString(v + " ")
		}
	}
	return buffer.String()
}


