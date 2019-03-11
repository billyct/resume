package resume

import "testing"

func TestReadDoc(t *testing.T) {
	path := "./test_documents/test_doc.doc"
	content, _ := ReadDoc(path)

	expected := "hello world"
	if content != expected {
		t.Errorf("TestReadDoc's content should be %s", expected)
	}
}
