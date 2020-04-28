package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	v := "docker"
	s := `<html>
<script>
function startTime() {
var today = new Date();
var h = today.getHours();
var m = today.getMinutes();
var s = today.getSeconds();
m = checkTime(m);
s = checkTime(s);
document.getElementById('clock').innerHTML =
h + ":" + m + ":" + s;
var t = setTimeout(startTime, 500);
}
function checkTime(i) {
if (i < 10) {i = "0" + i};  
return i;
}
</script>
	<body onload="startTime()">
	<h2>Hello from  ` + v + ` !</h2>

	<h1>The curent server time is </h1>
	<div id="clock"></div>
	
	</body>


</body>
</html>`
	fmt.Fprintf(file, s)
}
