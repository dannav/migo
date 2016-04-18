package migo

import "testing"

func TestSharedLayout(t *testing.T) {
	r := New("tests/sharedlayout")

	// make sure each key has page and layout defined
	for _, tMap := range r.TemplateMap {
		if len(tMap) != 2 {
			t.Errorf("invalid length for layout key: %v", tMap)
		}
	}

	if len(r.TemplateMap) == 2 {
		t.Log("proper amount of keys created for layout dir")
		if r.TemplateMap["home/about"][1] == "tests/sharedlayout/shared/layout.tmpl" && r.TemplateMap["home/index"][1] == "tests/sharedlayout/shared/layout.tmpl" {
			t.Log("proper value set for account layout since no sublayout defined")
		} else {
			t.Error("imporoper layout set for home")
		}
	}
}

func TestSubAndSharedLayout(t *testing.T) {
	r := New("tests/subLayoutAndSharedLayout")

	// make sure each key has page and layout defined
	for _, tMap := range r.TemplateMap {
		if len(tMap) != 2 {
			t.Errorf("invalid length for layout key: %v", tMap)
		}
	}

	if len(r.TemplateMap) == 4 {
		t.Log("proper amount of keys created for layout dir")

		if r.TemplateMap["account/login"][1] == "tests/subLayoutAndSharedLayout/shared/layout.tmpl" {
			t.Log("proper value set for account layout since no sublayout defined")
		} else {
			t.Error("account layout should be shared layout : tests/subLayoutAndSharedLayout/shared/layout.tmpl")
		}

		// check home for proper sub layout definition
		if r.TemplateMap["home/index"][1] == "tests/subLayoutAndSharedLayout/home/shared/layout.tmpl" {
			t.Log("proper sub layout defined for home/index")
		} else {
			t.Error("sub layout incorrect for home")
		}
	} else {
		t.Error("did no create keys for all pages")
	}
}
