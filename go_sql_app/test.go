package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "myapp/datastore"
)

type Entry struct {
    Name    string `form:"name" binding:"required"`
    Message string `form:"message" binding:"required"`
}

func main() {
    router := gin.Default()

    router.LoadHTMLGlob("views/*")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.POST("/", func(c *gin.Context) {
        var form Entry
        if err := c.ShouldBind(&form); err == nil {
            datastore.InsertData(form.Name, form.Message)
        }
        c.Redirect(http.StatusFound, "/")
    })

    router.Run(":8080")
}
```

File: `views/index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Entry Form</title>
</head>
<body>
<form action="/" method="post">
    <label>Name:</label><br>
    <input type="text" name="name" required><br>
    <label>Message:</label><br>
    <textarea name="message" required></textarea><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>
```

File `datastore/db.go`:

```go
package datastore

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertData(name string, message string) {
    db := openDB()
    defer db.Close()

    insert, err := db.Query("INSERT INTO entries (name, message) VALUES (?, ?)", name, message)
    if err != nil {
		panic(err.Error())
    }
    
 