package utilities

import "github.com/gin-gonic/gin"

// import "github.com/gin-contrib/cors"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Writer.Header()
		header.Set("Content-Type", "text/html; charset=utf-8")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		// header.Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, Origin, Authorization")
		if c.Request.Method == "OPTIONS" {
			header.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

// 		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
// 		if c.Request.Method == "OPTIONS" {
// 			c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	}
// }
