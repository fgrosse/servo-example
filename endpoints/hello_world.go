package endpoints

import (
	"fmt"
	"net/http"
)

func HomepageEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<body>
			<h1>Hello servo-example!</h1>

			<p>Try the following links</p>
			<ul>
				<li><a href="/hello/User">/hello/User</a></li>
				<li><a href="/controller">/controller</a></li>
				<li><a href="/composed">/composed</a></li>
				<li><a href="/container">/container</a></li>
				<li><a href="/error">/error</a></li>
			</ul>
		</body>
		</html>
	`)
}
