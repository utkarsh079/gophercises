package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	linkedNodesList := linkedNodes(doc)
	var links []Link
	for _, linkedNode := range linkedNodesList {
		links = append(links, buildLinks(linkedNode))
	}
	return links, nil
}

func getText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
		return text
	}
	if n.Type != html.ElementNode {
		return text
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

func buildLinks(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = getText(n)
	return ret
}

func linkedNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var linkedNodesList []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkedNodesList = append(linkedNodesList, linkedNodes(c)...)
	}
	return linkedNodesList
}

// func dfs(n *html.Node, padding string) {
// 	if n.Type == html.ElementNode {
// 		fmt.Println(padding, "<"+n.Data+">")
// 	} else {
// 		fmt.Println(padding, n.Data)
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
