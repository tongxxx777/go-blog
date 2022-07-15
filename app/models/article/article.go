package article

import (
	"blog/pkg/route"
	"strconv"
)

type Article struct {
	ID    uint64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link(a Article) string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
