package main
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"golang.org/x/net/html"
	"reflect"
)

func main(){
	url:=gethtml()
	docs,err :=html.Parse(url)
	//返回一个html node类型
	fmt.Println(reflect.TypeOf(docs))
	if err!=nil{
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//nil,因为按照值传递，+go的垃圾收集机制

	//link:=visit(nil,docs)
	for _, link := range visit(nil,docs) {
		fmt.Println(link)
	}
	fmt.Println(111)
}
func gethtml()[]byte{
	resp,err:=http.Get("www.sysu.edu.cn")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"fectch %v\n",err)
		os.Exit(1)
	}
	body,err:=ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err!=nil{
		fmt.Fprint(os.Stderr,"fetch%v\n",err)
		os.Exit(1)
	}
	return body
}
func visit(link []string,node* html.Node)[]string{
	if  node.Type==html.ElementNode && node.Data=="a"{
		for _,temp:=range node.Attr{
			link=append(link,temp.Val)
		}
	}
	for n:=node.FirstChild;n!=nil;n=n.NextSibling{
		link=visit(link,n)
	}
	return link
}