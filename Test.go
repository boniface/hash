package main

import "github.com/SlyMarbo/rss"
import (
	"github.com/advancedlogic/GoOse"
	"regexp"
	"strings"
	"log"
)

func main() {
//	println(filterContent("Home / Breaking News / \t\n\nHome / Breaking News / \n\n\n\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\n\t\t\t\n\nTweetEmail\nTweetEmailPOLICE in Kabwe have arrested a blind couple for allegedly stealing K15,000 meant for the Blind Skills Farmers’ Club in Serenje.\nThe police identified the couple as Isaiah Mushota, aged 56, and his wife Doreen Mando, 44, both of Katondo Township in Kabwe.\nCentral Province Commissioner of Police Standwell Lungu said that the Blind Skills Farmers’ Club was run by the visually impaired and supported by the Ministry of Community Development and Social Welfare.\nThe two were leaders in the farmers’ club at the time they allegedly stole the money before fleeing from Serenje.\nMr Lungu said police investigations revealed that the money (K15,000) was deposited in June this year through Finance Bank in Serenje into the club’s account but it was later withdrawn by Mushota from Lusaka.\nThe couple is suspected to have taken advantage of their positions in the club, where Mushota was chairperson and his wife treasurer, to access the funds illegally.\nMr Lungu said the couple had been on the run for a while but was finally apprehend yesterday in Katondo area and would be charged with theft.\nSource: Times of Zambia\n\t\t\t\n\t\t\t\n\t\t\t\t\t\n\t\t            \n            \tRelated Posts\n\t\t\t\n"))
//	getArticle()
//	getGetLinks()
	title:="Former former Opposition parties for reject ECZ ‘fake’ electronic transmission of election results’"
	println("Original Title is ",title)
	println("New Pretty title is ",metaDescription(title, 0, 156))
	println("The Keywords are  ",getKeyWords(title))

}

func getGetLinks() {
	var feedLink string = "https://www.facebook.com/feeds/page.php?id=129987587052000&format=rss20"
	feed, err := rss.Fetch(feedLink)
	if err != nil {
		// handle error.
	}
	var links = feed.Items
	for _, link := range links {

		println(" The THE TITILE!!!!!!  ",link.Title)
		println(" The THE ID!  ",link.ID)
		println(" The THE LINK !!!  ",link.Link)
		println(" The THE TITILE!!!!!!  ",link.Read)

		println(" The COntent is ",link.Content)

	}

}

func getArticle() {
	g := goose.New()
    article := g.ExtractFromUrl("http://www.times.co.zm/?p=33997")

	println("The Domain", article.Domain)
	println("The Date", article.PublishDate)
	println("The Title", article.Title)
	println("The ClearText", article.CleanedText)
	println("The TOP IMAGE", article.TopImage)

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

func stopWords() string{
	words:=
		"a|"+ "ii|"+ "about|"+ "above|"+ "according|"+ "across|"+ "actually|"+ "ad|"+ "adj|"+ "ae|"+ "af|"+ "after|"+ "afterwards|"+ "ag|"+ "again|"+ "against|"+ "ai|"+ "al|"+ "all|"+ "almost|"+ "alone|"+ "along|"+ "already|"+ "also|"+ "although|"+ "always|"+ "am|"+ "among|"+ "amongst|"+ "an|"+ "and|"+ "another|"+ "any|"+ "anyhow|"+ "anyone|"+ "anything|"+ "anywhere|"+ "ao|"+ "aq|"+ "ar|"+ "are|"+ "aren|"+ "aren't|"+ "around|"+ "arpa|"+ "as|"+ "at|"+ "au|"+ "aw|"+ "az|"+ "b|"+ "ba|"+ "bb|"+ "bd|"+ "be|"+ "became|"+ "because|"+ "become|"+ "becomes|"+ "becoming|"+ "been|"+ "before|"+ "beforehand|"+ "begin|"+ "beginning|"+ "behind|"+ "being|"+ "below|"+ "beside|"+ "besides|"+ "between|"+ "beyond|"+ "bf|"+ "bg|"+ "bh|"+ "bi|"+ "billion|"+ "bj|"+ "bm|"+ "bn|"+ "bo|"+ "both|"+ "br|"+ "bs|"+ "bt|"+ "but|"+ "buy|"+ "bv|"+ "bw|"+ "by|"+ "bz|"+ "c|"+ "ca|"+ "can|"+ "can't|"+ "cannot|"+ "caption|"+ "cc|"+ "cd|"+ "cf|"+ "cg|"+ "ch|"+ "ci|"+ "ck|"+ "cl|"+ "click|"+ "cm|"+ "cn|"+ "co|"+ "co.|"+ "com|"+ "copy|"+ "could|"+ "couldn|"+ "couldn't|"+ "cr|"+ "cs|"+ "cu|"+ "cv|"+ "cx|"+ "cy|"+ "cz|"+ "d|"+ "de|"+ "did|"+ "didn|"+ "didn't|"+ "dj|"+ "dk|"+ "dm|"+ "do|"+ "does|"+ "doesn|"+ "doesn't|"+ "don|"+ "don't|"+ "down|"+ "during|"+ "dz|"+ "e|"+ "each|"+ "ec|"+ "edu|"+ "ee|"+ "eg|"+ "eh|"+ "eight|"+ "eighty|"+ "either|"+ "else|"+ "elsewhere|"+ "end|"+ "ending|"+ "enough|"+ "er|"+ "es|"+ "et|"+ "etc|"+ "even|"+ "ever|"+ "every|"+ "everyone|"+ "everything|"+ "everywhere|"+ "except|"+ "f|"+ "few|"+ "fi|"+ "fifty|"+ "find|"+ "first|"+ "five|"+ "fj|"+ "fk|"+ "fm|"+ "fo|"+ "for|"+ "former|"+ "formerly|"+ "forty|"+ "found|"+ "four|"+ "fr|"+ "free|"+ "from|"+ "further|"+ "fx|"+ "g|"+ "ga|"+ "gb|"+ "gd|"+ "ge|"+ "get|"+ "gf|"+ "gg|"+ "gh|"+ "gi|"+ "gl|"+ "gm|"+ "gmt|"+ "gn|"+ "go|"+ "gov|"+ "gp|"+ "gq|"+ "gr|"+ "gs|"+ "gt|"+ "gu|"+ "gw|"+ "gy|"+ "h|"+ "had|"+ "has|"+ "hasn|"+ "hasn't|"+ "have|"+ "haven|"+ "haven't|"+ "he|"+ "he'd|"+ "he'll|"+ "he's|"+ "help|"+ "hence|"+ "her|"+ "here|"+ "here's|"+ "hereafter|"+ "hereby|"+ "herein|"+ "hereupon|"+ "hers|"+ "herself|"+ "him|"+ "himself|"+ "his|"+ "hk|"+ "hm|"+ "hn|"+ "home|"+ "homepage|"+ "how|"+ "however|"+ "hr|"+ "ht|"+ "htm|"+ "html|"+ "http|"+ "hu|"+ "hundred|"+ "i|"+ "i'd|"+ "i'll|"+ "i'm|"+ "i've|"+ "i.e.|"+ "id|"+ "ie|"+ "if|"+ "il|"+ "im|"+ "in|"+ "inc|"+ "inc.|"+ "indeed|"+ "information|"+ "instead|"+ "int|"+ "into|"+ "io|"+ "iq|"+ "ir|"+ "is|"+ "isn|"+ "isn't|"+ "it|"+ "it's|"+ "its|"+ "itself|"+ "j|"+ "je|"+ "jm|"+ "jo|"+ "join|"+ "jp|"+ "k|"+ "ke|"+ "kg|"+ "kh|"+ "ki|"+ "km|"+ "kn|"+ "kp|"+ "kr|"+ "kw|"+ "ky|"+ "kz|"+ "l|"+ "la|"+ "last|"+ "later|"+ "latter|"+ "lb|"+ "lc|"+ "least|"+ "less|"+ "let|"+ "let's|"+ "li|"+ "like|"+ "likely|"+ "lk|"+ "ll|"+ "lr|"+ "ls|"+ "lt|"+ "ltd|"+ "lu|"+ "lv|"+ "ly|"+ "m|"+ "ma|"+ "made|"+ "make|"+ "makes|"+ "many|"+ "maybe|"+ "mc|"+ "md|"+ "me|"+ "meantime|"+ "meanwhile|"+ "mg|"+ "mh|"+ "microsoft|"+ "might|"+ "mil|"+ "million|"+ "miss|"+ "mk|"+ "ml|"+ "mm|"+ "mn|"+ "mo|"+ "more|"+ "moreover|"+ "most|"+ "mostly|"+ "mp|"+ "mq|"+ "mr|"+ "mrs|"+ "ms|"+ "msie|"+ "mt|"+ "mu|"+ "much|"+ "must|"+ "mv|"+ "mw|"+ "mx|"+ "my|"+ "myself|"+ "mz|"+ "n|"+ "na|"+ "namely|"+ "nc|"+ "ne|"+ "neither|"+ "net|"+ "netscape|"+ "never|"+ "nevertheless|"+ "new|"+ "next|"+ "nf|"+ "ng|"+ "ni|"+ "nine|"+ "ninety|"+ "nl|"+ "no|"+ "nobody|"+ "none|"+ "nonetheless|"+ "noone|"+ "nor|"+ "not|"+ "nothing|"+ "now|"+ "nowhere|"+ "np|"+ "nr|"+ "nu|"+ "nz|"+ "o|"+ "of|"+ "off|"+ "often|"+ "om|"+ "on|"+ "once|"+ "one|"+ "one's|"+ "only|"+ "onto|"+ "or|"+ "org|"+ "other|"+ "others|"+ "otherwise|"+ "our|"+ "ours|"+ "ourselves|"+ "out|"+ "over|"+ "overall|"+ "own|"+ "p|"+ "pa|"+ "page|"+ "pe|"+ "per|"+ "perhaps|"+ "pf|"+ "pg|"+ "ph|"+ "pk|"+ "pl|"+ "pm|"+ "pn|"+ "pr|"+ "pt|"+ "pw|"+ "py|"+ "q|"+ "qa|"+ "r|"+ "rather|"+ "re|"+ "recent|"+ "recently|"+ "reserved|"+ "ring|"+ "ro|"+ "ru|"+ "rw|"+ "s|"+ "sa|"+ "same|"+ "sb|"+ "sc|"+ "sd|"+ "se|"+ "seem|"+ "seemed|"+ "seeming|"+ "seems|"+ "seven|"+ "seventy|"+ "several|"+ "sg|"+ "sh|"+ "she|"+ "she'd|"+ "she'll|"+ "she's|"+ "should|"+ "shouldn|"+ "shouldn't|"+ "si|"+ "since|"+ "site|"+ "six|"+ "sixty|"+ "sj|"+ "sk|"+ "sl|"+ "sm|"+ "sn|"+ "so|"+ "some|"+ "somehow|"+ "someone|"+ "something|"+ "sometime|"+ "sometimes|"+ "somewhere|"+ "sr|"+ "st|"+ "still|"+ "stop|"+ "su|"+ "such|"+ "sv|"+ "sy|"+ "sz|"+ "t|"+ "taking|"+ "tc|"+ "td|"+ "ten|"+ "text|"+ "tf|"+ "tg|"+ "test|"+ "th|"+ "than|"+ "that|"+ "that'll|"+ "that's|"+ "the|"+ "their|"+ "them|"+ "themselves|"+ "then|"+ "thence|"+ "there|"+ "there'll|"+ "there's|"+ "thereafter|"+ "thereby|"+ "therefore|"+ "therein|"+ "thereupon|"+ "these|"+ "they|"+ "they'd|"+ "they'll|"+ "they're|"+ "they've|"+ "thirty|"+ "this|"+ "those|"+ "though|"+ "thousand|"+ "three|"+ "through|"+ "throughout|"+ "thru|"+ "thus|"+ "tj|"+ "tk|"+ "tm|"+ "tn|"+ "to|"+ "together|"+ "too|"+ "toward|"+ "towards|"+ "tp|"+ "tr|"+ "trillion|"+ "tt|"+ "tv|"+ "tw|"+ "twenty|"+ "two|"+ "tz|"+ "u|"+ "ua|"+ "ug|"+ "uk|"+ "um|"+ "under|"+ "unless|"+ "unlike|"+ "unlikely|"+ "until|"+ "up|"+ "upon|"+ "us|"+ "use|"+ "used|"+ "using|"+ "uy|"+ "uz|"+ "v|"+ "va|"+ "vc|"+ "ve|"+ "very|"+ "vg|"+ "vi|"+ "via|"+ "vn|"+ "vu|"+ "w|"+ "was|"+ "wasn|"+ "wasn't|"+ "we|"+ "we'd|"+ "we'll|"+ "we're|"+ "we've|"+ "web|"+ "webpage|"+ "website|"+ "welcome|"+ "well|"+ "were|"+ "weren|"+ "weren't|"+ "wf|"+ "what|"+ "what'll|"+ "what's|"+ "whatever|"+ "when|"+ "whence|"+ "whenever|"+ "where|"+ "whereafter|"+ "whereas|"+ "whereby|"+ "wherein|"+ "whereupon|"+ "wherever|"+ "whether|"+ "which|"+ "while|"+ "whither|"+ "who|"+ "who'd|"+ "who'll|"+ "who's|"+ "whoever|"+ "NULL|"+ "whole|"+ "whom|"+ "whomever|"+ "whose|"+ "why|"+ "will|"+ "with|"+ "within|"+ "without|"+ "won|"+ "won't|"+ "would|"+ "wouldn|"+ "wouldn't|"+ "ws|"+ "www|"+ "x|"+ "y|"+ "ye|"+ "yes|"+ "yet|"+ "you|"+ "you'd|"+ "you'll|"+ "you're|"+ "you've|"+ "your|"+ "yours|"+ "yourself|"+ "yourselves|"+ "yt|"+ "yu|"+ "z|"+ "za|"+ "zm|"+ "zr|"+ "z|"+ "org|"+ "inc|"+ "width|"+ "length|"
	return words
}

func prettyUrl(title string ) string {
	//let's make pretty urls from title
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	regStoWords := regexp.MustCompile("for|to|old|the|is|do|at|be|was|were|am|I|says|say|of|on")
	if err != nil {
		log.Fatal(err)
	}
	cleanTitle:=regStoWords.ReplaceAllLiteralString(title,"")
	prettyurl := reg.ReplaceAllString(cleanTitle, "-")
	prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))
	return prettyurl
}

func metaDescription(s string,pos,length int) string{
	runes:=[]rune(s)
	l := pos+length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getKeyWords(title string) string{
	regSpace := regexp.MustCompile(`\s`)
	regStoWords := regexp.MustCompile("to|old|the|is|do|at|be|was|were|am|I|says|say|of")
	cleanTitle:=regStoWords.ReplaceAllLiteralString(title,"")
	shortenTitle:=metaDescription(cleanTitle,0,70)
	keyWords := regSpace.ReplaceAllString(shortenTitle, ",")
	return keyWords
}
