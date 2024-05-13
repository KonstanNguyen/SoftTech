package models

import (
	"SoftTech/blogs/types"
	"errors"
)

var blogList = []types.Blog{
	{Id: 1, Title: "Blog 1"},
	{Id: 2, Title: "Blog 2"},
	{Id: 3, Title: "Blog 3"},
}

func GetBlogList() []types.Blog {
	blogList := []types.Blog{{Id: 1, Title: "Blog 1"}, {Id: 2, Title: "Blog 2"}, {Id: 3, Title: "Blog 3"}}
	return blogList
}

func GetBlogDetails(blogId uint64) (types.Blog, error) {
	blogList := GetBlogList()
	for _, blogDetail := range blogList {
		if uint64(blogDetail.Id) == blogId {
			return blogDetail, nil
		}
	}
	return types.Blog{}, errors.New("blog not found")
}

func CreateBlog(data *types.Blog) error {
	blogList = append(blogList, *data)
	return errors.New("blog already exists")
}

func DeleteBlog(id uint64) error {
	for i, _ := range blogList {
		if uint64(blogList[i].Id) == id {
			blogList = append(blogList[:i], blogList[i+1:]...)
			return nil
		}
	}
	return errors.New("blog is not found")
}

func UpdateBlog(data *types.Blog) error {
	for i, _ := range blogList {
		if blogList[i].Id == data.Id {
			blogList[i] = *data
			return nil
		}
	}
	return errors.New("blog is not found")
}
