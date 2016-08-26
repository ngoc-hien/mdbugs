package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
)

type ApplicationTest struct {
	testing.TestSuite
}

func (t *ApplicationTest) Before() {
	println("Set up test mdbugs")
}

var _title string = "Correct page title"
func (t *ApplicationTest) TestGetCorrectPageTitle() {	
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains(_title)
}

func (t *ApplicationTest) TestPostCorrectPageTitle() {	
	t.PostForm("/", url.Values{"message": {"test message"}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains(_title)
}

func (t *ApplicationTest) TestPostEmptyMsg() {	
	t.PostForm("/", url.Values{"message": {""}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("Please input message")	
}

func (t *ApplicationTest) TestPostMsg() {	
	t.PostForm("/", url.Values{"message": {"testing normal message"}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("Success")	
}


func (t *ApplicationTest) After() {
	println("Finish testing")
}