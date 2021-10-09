package insta

import (
	"net/url"
	"testing"
)

func TestUsersGet(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Get(expected)
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}

func TestUsersSelf(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Self()
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}

func TestUsersSelfFeed(t *testing.T) {
	//testing the feed
	var params = url.Values{}
	params.Set("count", "10")

	items, content, err := client.Users.SelfFeed(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of feeds is 0")
	}
}

func TestUsersRecentPosts(t *testing.T) {
	//testing recent posts
	items, content, err := client.Users.RecentPosts(selfId, nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent post is 0")
	}
}

func TestUsersLikedPosts(t *testing.T) {
	//testing the liked posts by the user
	items, content, err := client.Users.LikedPosts(nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of liked post is 0")
	}
}

func TestUsersSearch(t *testing.T) {
	//testing the search results
	items, content, err := client.Users.Search("japan", 5)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of search result is 0")
	}
}
