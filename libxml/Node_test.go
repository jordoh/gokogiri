package libxml

import (
	"testing"
)

func TestElementAttributes(t *testing.T) {
	doc := HtmlParse("<div id='hi' />")
	root := doc.RootNode()
	div := root.Search("//div").First().(*Element)
	if div.AttributeValue("id") != "hi" {
		t.Error("looking for id should return 'hi'")
	}
	if div.AttributeValue("class") != "" {
		t.Error("Non-existant attributes should return nil")
	}
	div.SetAttributeValue("class", "classy")
	if div.AttributeValue("class") != "classy" {
		t.Error("Attributes aren't set")
	}
}

func TestElementName(t *testing.T) {
	doc := HtmlParse("<div id='hi' />")
	root := doc.RootNode()
	div := root.Search("//div").First()
	if div.Name() != "div" {
		t.Error("Something is wrong with XMLNode.Name()")
	}
	div.SetName("span")
	if div.Name() != "span" {
		t.Error("Something is wrong with XMLNode.SetName()")
	}
}

func TestElementDump(t *testing.T) {
	doc := HtmlParse("<div id='hi' />")
	root := doc.RootNode()
	div := root.Search("//div").First()
	result := div.Dump()
	if result != "<div id=\"hi\"/>" {
		t.Error("Node dumping is being... dumpy. Got back this pile of poo: ", result)
	}
	div.SetName("span")
	result = div.Dump()
	if result != "<span id=\"hi\"/>" {
		t.Error("Node dumping is being... dumpy. Got back this pile of poo: ", result)
	}
}

func TestElementRemove(t *testing.T) {
	doc := HtmlParse("<html><body><div><span>hi</span></div></body></html>")
	root := doc.RootNode()
	span := root.Search("//span").First()
	span.Remove()
	result := doc.DumpHTML()
	if result != "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html><body><div></div></body></html>\n" {
		t.Error("Node dumping is being... dumpy. Got back this pile of poo: ", result)
	}

}
