package insta

import (
	"fmt"
	"net/url"
)

type PostsApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/media/#get_posts
func (o *PostsApi) Get(postsId string) (*Posts, *Content, error) {
	var path = fmt.Sprintf("/posts/%s", postsId)
	var item = new(Posts)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_posts_by_shortcode
func (o *PostsApi) Shortcode(shortcode string) (*Posts, *Content, error) {
	var path = fmt.Sprintf("/posts/shortcode/%s", shortcode)
	var item = new(Posts)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_posts_popular
func (o *PostsApi) Search(params url.Values) ([]Posts, *Content, error) {
	var path = "/posts/search/"
	var item = new([]Posts)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_posts_popular
func (o *PostsApi) Popular() ([]Posts, *Content, error) {
	var path = "/posts/popular/"
	var item = new([]Posts)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}
