// https://leetcode.com/problems/design-browser-history/
// https://www.codingninjas.com/studio/problems/browser_2427908

package main

import "fmt"

/************ NODE ************/

type Node struct {
	data string
	next *Node
	back *Node
}

/************ BROWSER HISTORY ************/

type BrowserHistory struct {
	currentPage *Node
}

/************ CONSTRUCTOR ************/

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		currentPage: &Node{data: homepage},
	}
}

/************ VISIT ************/

func (b *BrowserHistory) Visit(url string) {

	newNode := &Node{data: url}

	// clear forward history automatically
	b.currentPage.next = newNode
	newNode.back = b.currentPage

	b.currentPage = newNode
}

/************ BACK ************/

func (b *BrowserHistory) Back(steps int) string {

	for steps > 0 && b.currentPage.back != nil {
		b.currentPage = b.currentPage.back
		steps--
	}

	return b.currentPage.data
}

/************ FORWARD ************/

func (b *BrowserHistory) Forward(steps int) string {

	for steps > 0 && b.currentPage.next != nil {
		b.currentPage = b.currentPage.next
		steps--
	}

	return b.currentPage.data
}

/************ TEST ************/

func main() {

	browser := Constructor("leetcode.com")

	browser.Visit("google.com")
	browser.Visit("facebook.com")
	browser.Visit("youtube.com")

	fmt.Println(browser.Back(1))    // facebook.com
	fmt.Println(browser.Back(1))    // google.com
	fmt.Println(browser.Forward(1)) // facebook.com

	browser.Visit("linkedin.com")

	fmt.Println(browser.Forward(2)) // linkedin.com
	fmt.Println(browser.Back(2))    // google.com
	fmt.Println(browser.Back(7))    // leetcode.com
}