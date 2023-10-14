package old

import "github.com/gin-gonic/gin"

var oldList = make(map[string][]string)

func Init() {
	oldList = map[string][]string{}
}

func Set(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		return
	}

	oldList = c.Request.PostForm
}

func Get() map[string][]string {
	return oldList
}
