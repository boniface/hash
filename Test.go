package main
import "github.com/advancedlogic/GoOse"


func main(){
	println("hello World ")
	g := goose.New()
	article := g.ExtractFromUrl("http://www.nyasatimes.com/2014/08/21/pp-on-deadbed-claims-udf-malawi-opposition-fall-out/")
	println("The Domain", article.Domain)
	println("The Date", article.PublishDate)
	println("The Title", article.Title)
	println("The Link", article.FinalUrl)
	println("The TOP IMAGE", article.CanonicalLink)
}
