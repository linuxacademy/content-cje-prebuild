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
	s := `<%@ page import = "java.io.*,java.util.*, javax.servlet.*" %>
<html>
<body>
<h2>Hello `+ v +` the Tomcat server is running!</h2>
<h1>The curent server time is 
<%
         Date date = new Date();
         out.print( "<h2 align = \"center\">" +date.toString()+"</h2>");
 %>
</h1>
</body>
</html>`
	fmt.Fprintf(file, s)
}
