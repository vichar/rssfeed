package rssfeed_test

import (
	"strconv"
	"testing"

	"github.com/vichar/rssfeed"
)

func TestParseRSSDataWithValidRSSURL(t *testing.T) {

	t.Run("ParseHTTPResponse", func(t *testing.T) {
		getResponse, httpError := rssfeed.HTTPGet("https://www.npr.org/rss/podcast.php?id=500005")
		if httpError == nil {
			response := rssfeed.ParseHTTPResponse(getResponse)
			rssChannel, parsingError := rssfeed.ParseRSSData(response.Body)
			if parsingError == nil {
				if rssChannel.Items.Len() <= 0 {
					t.Errorf("Parsed Feed should contain at least 1 item")
				}
				if len(rssChannel.Title) <= 0 {
					t.Errorf("Parsed Feed should contain title")

				}
			} else {
				t.Errorf("ParseRSSData should not return error for a valid URL but gets %s", parsingError.Error())
			}
		} else {
			t.Errorf("ParseHTTPResponse should not return error for a valid URL but gets %s", httpError.Error())
		}
	})
}
func TestParseHTTPResponseWithValidURL(t *testing.T) {
	t.Run("ParseHTTPResponse", func(t *testing.T) {
		response, error := rssfeed.HTTPGet("https://www.npr.org/rss/podcast.php?id=500005")
		if error == nil {
			var getResponse rssfeed.HTTPResponse = rssfeed.ParseHTTPResponse(response)
			if getResponse.HTTPStatus != "200 OK" {
				t.Errorf("HTTP Status shoud be %s but gets %s", getResponse.HTTPStatus, "200 OK")
			}
			if getResponse.HTTPStatusCode != 200 {
				t.Errorf("HTTP Status Code should be %s but gets %s", strconv.Itoa(getResponse.HTTPStatusCode), "200")
			}
			if len(getResponse.Body) <= 0 {
				t.Errorf("HTTP Response Body should not be empty but length is %s", strconv.Itoa(len(getResponse.Body)))
			}
		} else {
			t.Errorf("HTTPGet should not return error for a valid URL but gets %s", error.Error())
		}
	})
}

func TestHTTPGetWithValidURL(t *testing.T) {
	t.Run("HTTPGet should return valid String URL", func(t *testing.T) {
		response, error := rssfeed.HTTPGet("https://www.npr.org/rss/podcast.php?id=500005")
		want := "200 OK"
		if want != response.Status {
			t.Errorf("It should say %s but gets %s", want, response.Status)
		}
		if error != nil {
			t.Errorf("An Unexpected Error occurs %s", error.Error())
		}

	})

}

func TestGetContentWithInvalidURL(t *testing.T) {
	t.Run("HTTPGet should return Empty String", func(t *testing.T) {
		response, error := rssfeed.HTTPGet("")
		if error == nil {
			t.Errorf("Invalid URL is Expected %s %s", response.Status, error.Error())
		}
		if error.Error() != "Invalid URL" {
			t.Errorf("Unexpected Error has occurred %s %s", response.Status, error.Error())
		}
	})

}
