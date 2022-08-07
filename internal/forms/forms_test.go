package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/someting", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.Required("a", "b")

	if form.Valid() {
		t.Error("form is valid when required fields are missing")
	}

	postData.Add("a", "go")
	postData.Add("b", "lang")
	form = New(postData)

	if !form.Valid() {
		t.Error("valid form found to be invalid")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	postData.Add("a", "go")

	form := New(postData)

	if !form.Has("a") {
		t.Error("form does not know it has 'a'")
	}

	if form.Has("xxx") {
		t.Error("form thinks it has 'xxx'")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.MinLength("x", 10)

	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	if form.Errors.Get("x") == "" {
		t.Error("Get returned value for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")

	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")

	form = New(postedValues)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}

	if form.Errors.Get("another_field") != "" {
		t.Error("should not have an error, but go one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("HHH")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when it is valid")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "z")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid email when it is invalid")
	}
}
