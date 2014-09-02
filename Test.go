package main
import "github.com/advancedlogic/GoOse"
import ("net/http"; "io";

"os"
"log"
	"io/ioutil"

)

func main(){
	println("hello World ")
	g := goose.New()
	article := g.ExtractFromUrl("http://www.nyasatimes.com/2014/08/21/pp-on-deadbed-claims-udf-malawi-opposition-fall-out/")
	println("The Domain", article.Domain)
	println("The Date", article.PublishDate)
	println("The Title", article.Title)
	println("The Link", article.FinalUrl)
	println("The TOP IMAGE", article.TopImage)
	out, err := os.Create("downloadedImage.jpg")

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get("http://www.nyasatimes.com/wp-content/uploads/ndanga1.jpg")
	data1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("google_logo.png", data1,0666 )




	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	println(" Output DOWN",n)
}
