package gohaml

import "testing"

var nestingTests = []IO{
	IO{"%tag1\n  %tag2", "<tag1>\n	<tag2 />\n</tag1>"},
	IO{"%tag1\n%tag2", "<tag1 />\n<tag2 />"},
	IO{"%tag1\n%tag2\n%tag3", "<tag1 />\n<tag2 />\n<tag3 />"},
	IO{"%tag1\n  %tag2\n  %tag3", "<tag1>\n\t<tag2 />\n\t<tag3 />\n</tag1>"},
	IO{"%tag1\n  %tag2\n    %tag3", "<tag1>\n\t<tag2>\n\t\t<tag3 />\n\t</tag2>\n</tag1>"},
	IO{"%tag1\n  %tag2\n    %tag3 tag content", "<tag1>\n\t<tag2>\n\t\t<tag3>tag content</tag3>\n\t</tag2>\n</tag1>"},
//	IO{"%tag1\n  %tag2\n    %tag3 tag content\n    %tag4", "<tag1>\n\t<tag2>\n\t\t<tag3>tag content</tag3>\n\t\t<tag4 />\n\t</tag2>\n</tag1>"},
	// IO{"%tag1\n  %tag2\n    %tag3\n    %tag4 tag content", "<tag1>\n\t<tag2>\n\t\t<tag3 />\n\t\t<tag4>tag content</tag4>\n\t</tag2>\n</tag1>"},
}

func TestNesting(t *testing.T) {
	for _, io := range nestingTests {
		scope := make(map[string]interface{})
		scope["key1"] = "value1"
		scope["key2"] = "value2"
	
		engine := NewEngine(io.input)
		output := engine.Render(scope)
		if output != io.expected {
			t.Errorf("Input    %q\nexpected %q\ngot      %q", io.input, io.expected, output)
		}
	}
}
