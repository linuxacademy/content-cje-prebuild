package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("index.jsp")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	v := os.Getenv("name")
	s := `<html>
<body>
<h2>Hello `+ v +` the Tomcat server is running!</h2>
<h1>The curent server time is <%= new java.util.Date() %></h1>
</body>
</html>`
	fmt.Fprintf(file, s)
}
