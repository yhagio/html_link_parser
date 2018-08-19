package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// <a href="example.com">Link Text</a>
type Link struct {
	Href     string
	LinkText string
}

// Parse HTML document
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	linkNodes := FindLinkNodes(doc)
	var links []Link
	for _, node := range linkNodes {
		builtNode := BuildLink(node)
		// Check if the node is empty
		if builtNode != (Link{}) {
			links = append(links, builtNode)
		}
	}
	return links, nil
}

// Find <a> in HTML document and build links as it finds in the nodes tree,
// and returns slice of HTML nodes
func FindLinkNodes(node *html.Node) []*html.Node {
	// Found a link node <a>, no need to go deeper, just return it
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	// Recursively find more in deeper nodes tree
	var foundLinkNodes []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		foundLinkNodes = append(foundLinkNodes, FindLinkNodes(child)...)
	}
	return foundLinkNodes
}

// Build Link(Lhref link and its text) from the passed node
func BuildLink(node *html.Node) Link {
	var link Link
	for _, attr := range node.Attr {
		if attr.Key == "href" && attr.Val != "" {
			link.Href = attr.Val
			break
		}
	}
	if link.Href != "" {
		link.LinkText = GetLinkText(node)
	}
	return link
}

// Get link text
func GetLinkText(node *html.Node) string {
	// Found the text
	if node.Type == html.TextNode {
		return node.Data
	}
	// If not element, we don't care
	if node.Type != html.ElementNode {
		return ""
	}
	// Recursively find more in deeper nodes tree
	var txt string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		txt += GetLinkText(child)
	}
	// Clean up the preceding/trailing spaces, linebreaks
	return strings.Join(strings.Fields(txt), " ")
}

func main() {
	fmt.Println("\n=====👽HTML Link Parser👽=====\n")
	fmt.Println("Type URL (i.e. https://golang.org):")

	// Get user's input url
	var url string
	fmt.Scanf("%s\n", &url)

	fmt.Printf("\nReading %s ...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Read HTML document (byte[])
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", html)
	fmt.Println("Parsing ...")

	r := strings.NewReader(string(html))
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n<<< RESULT (Total: %d) >>>\n", len(links))
	for i, link := range links {
		fmt.Printf("\n(%d) \n Link:%s \n Text:%s\n", i+1, link.Href, link.LinkText)
	}
}
