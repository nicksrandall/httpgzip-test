# HTTPGzip Error Example

## Error
When routing to the root path of my file server, I'm getting in error in my browser that says that I'm getting infinite redirects.

## Steps to reproduce
Start the server and direct your browser to `localhost:8080/static/`

> I have commented out a line that implements `http.FileServer`. Feel free to uncomment it to see that you do not get this error when using it.
