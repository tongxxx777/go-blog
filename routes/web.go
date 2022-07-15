package routes

import (
	"blog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

}
