package insta

import (
	"net/url"
	"testing"
)

func TestPostsGet(t *testing.T) {
	var expected = postsId

	posts, content, err := client.Posts.Get(expected)
	isInvalidMetaData(content, err, t)

	if posts.Id != expected {
		t.Errorf("expected posts_id is wrong: %s", posts.Id)
	}
}

func TestPostsShortcode(t *testing.T) {
	var expected = postsId
	var shortcode = "xZaWLwvIp5"

	posts, content, err := client.Posts.Shortcode(shortcode)
	isInvalidMetaData(content, err, t)

	if posts.Id != expected {
		t.Errorf("expected posts_id is wrong: %s", posts.Id)
	}
}

func TestPostsPopular(t *testing.T) {
	items, content, err := client.Posts.Popular()
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent post is 0")
	}
}

func TestPostsSearch(t *testing.T) {
	var params = url.Values{}
	params.Set("lat", "1.284962")
	params.Set("lng", "103.858699")
	params.Set("distance", "1000")

	items, content, err := client.Posts.Search(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent post is 0")
	}
}
