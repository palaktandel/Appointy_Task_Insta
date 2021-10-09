package insta

import (
	"fmt"
	"net/url"
	"strconv"
)

type UserApi struct {
	Instagram *Instagram
}

func (o *UserApi) Get(userId string) (*User, *Content, error) {
	var path = fmt.Sprintf("/users/%s", userId)
	var item = new(User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

func (o *UserApi) Self() (*User, *Content, error) {
	var path = "/users/self"
	var item = new(User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

func (o *UserApi) SelfFeed(params url.Values) ([]Posts, *Content, error) {
	var path = "/users/self/feed"
	var item = new([]Posts)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}

func (o *UserApi) RecentPosts(userId string, params url.Values) ([]Posts, *Content, error) {
	var path = fmt.Sprintf("/users/%s/media/recent", userId)
	var item = new([]Posts)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}

func (o *UserApi) LikedPosts(params url.Values) ([]Posts, *Content, error) {
	var path = "/users/self/media/liked"
	var item = new([]Posts)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}

func (o *UserApi) Search(query string, count int) ([]User, *Content, error) {
	var params = url.Values{}
	params.Set("q", query)
	params.Set("count", strconv.Itoa(count))

	var path = "/users/search"
	var item = new([]User)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}
