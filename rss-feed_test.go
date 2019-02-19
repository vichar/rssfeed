package rssfeed_test

import (
	"testing"

	"github.com/vichar/rssfeed"
)

func TestGetContentWithValidURL(t *testing.T) {
	t.Run("GetContent should return valid String URL", func(t *testing.T) {
		response, error := rssfeed.GetContent("https://www.npr.org")
		want := "200 OK"
		if want != response.Status {
			t.Errorf("It should say %s but get %s", want, response.Status)
		}
		if error != nil {
			t.Errorf("An Unexpected Error occurs %s", error.Error())
		}

	})

}

func TestGetContentWithInvalidURL(t *testing.T) {
	t.Run("GetContent should return Empty String", func(t *testing.T) {
		response, error := rssfeed.GetContent("")
		if error == nil {
			t.Errorf("An Error is Expected %s %s", response.Status, error.Error())
		}
	})

}
