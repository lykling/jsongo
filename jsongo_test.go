package jsongo

import (
	"reflect"
	"testing"
)

func TestLoadStringTo(t *testing.T) {
	str := `{"name":{"first":"Leong","last":"Pride"},"list":["1", "hello"]}`
	var jg JSONGo
	jg.LoadString(str)
	str = `{"name":{"first":"Leong","last":"Pride"},"list":[1, "hello"]}`
	var jg2 JSONGo
	jg2.LoadString(str)
	ele1 := jg.Get("list").Get("1").ToString()
	ele2 := jg2.Get("list").Get("1").ToString()
	if ele1 != ele2 {
		t.Errorf("Error: value %v doesn't match %v\n", ele1, ele2)
	}
}

func TestGet(t *testing.T) {
	str := `{"name":{"first":"Leong","last":"Pride"},"list":["1", "hello"]}`
	jg := &JSONGo{}
	jg.LoadString(str)
	name_first := jg.Get("name").Get("first").ToString()
	if name_first != `"Leong"` {
		t.Errorf("Error: %v doesn't match %s\n", name_first, `"Leong"`)
	}
	empty := jg.Get("target")
	if empty != nil {
		t.Errorf("Error: %v doesn't match %s\n", empty, nil)
	}
	tmp := jg.Get("name")
	if reflect.TypeOf(tmp).String() != reflect.TypeOf(jg).String() {
		t.Errorf("Error: type %v doesn't match %v",
			reflect.TypeOf(tmp).String(), reflect.TypeOf(jg).String())
	}
}

func TestToString(t *testing.T) {
	str := `{"name":{"first":"Leong","last":"Pride"},"list":["1", "hello"]}`
	var jg JSONGo
	jg.LoadString(str)
	str1 := jg.Get("list").ToString()
	str2 := `["1","hello"]`
	if str1 != str2 {
		t.Errorf("Error: %v doesn't match %v", str1, str2)
	}
	str1 = jg.Get("name").Get("first").ToString()
	str2 = `"Leong"`
	if str1 != str2 {
		t.Errorf("Error: %v doesn't match %v", str1, str2)
	}
}

func TestPretty(t *testing.T) {
	str := `{"name":{"first":"Leong","last":"Pride"},"list":["1", "hello"]}`
	jg := &JSONGo{}
	jg.LoadString(str)
	jg2 := &JSONGo{}
	jg2.LoadString(str)
	if jg.Pretty() != jg2.Pretty() {
		t.Errorf("Error: pretty error")
	}
}

func TestGetType(t *testing.T) {
	str := `{"name":{"first":"Leong","last":"Pride"},"list":["1", "hello"]}`
	var jg JSONGo
	jg.LoadString(str)
	type1 := jg.GetType()
	type2 := jg.Get("list").GetType()
	type3 := jg.Get("list").Get("1").GetType()
	type4 := jg.Get("name").GetType()
	type5 := jg.Get("name").Get("last").GetType()
	if type1 != "Object" {
		t.Errorf("Error: type %v doesn't match %v\n", type1, "Object")
	}
	if type2 != "Array" {
		t.Errorf("Error: type %v doesn't match %v\n", type2, "Array")
	}
	if type3 != "String" {
		t.Errorf("Error: type %v doesn't match %v\n", type3, "String")
	}
	if type4 != "Object" {
		t.Errorf("Error: type %v doesn't match %v\n", type4, "Object")
	}
	if type5 != "String" {
		t.Errorf("Error: type %v doesn't match %v\n", type5, "String")
	}
}

func TestLoadFile(t *testing.T) {
	filename := "test.json"
	var jg JSONGo
	jg.LoadFile(filename)
	name := jg.Get("author").Get("name").ToString()
	if name != `"Pride Leong"` {
		t.Errorf("Error: value %v doesn't match %v\n", name, `"Pride Leong"`)
	}
}
