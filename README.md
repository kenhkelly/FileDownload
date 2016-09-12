# FileDownload 
v0.1

#### By: Ken Kelly @kenhkelly

### Overview

This is a little server, written in Go, designed to listen to for a URL you send it to download a file. It was designed to be used with bookmarklets.

The purpose for me was to have something readily available so that I could quickly download my pictures from Facebook and Instagram. 

You will find two javascript files, one called filedownload.js (which does both Facebook and Instagram) and one called instagram.js. I recommend crunching them with [Bookmarklet Crunchinator](http://ted.mielczarek.org/code/mozilla/bookmarklet.html). 

### Usage

After getting the code, just build the repo with `go build filedownload.go`.

Easiest way to get going is to then simply run the binary file built. You can pass in arguments such as `--port 8080` to change the listening port and `--dir <dir>` to change the directory. The default directory is the one you are running the binary from.

If using the bookmarklet, make sure to update the `dl_serv` variable to the correct hostname and port. Open a image in Facebook or Instagram and press the bookmarklet. If all goes well, a window should pop open and close letting you know how many bytes were stored. The application will also output to the screen.

If you are not using the bookmarklet, simply navigate to `http://<host>:<port>/get?url=<url>` and the file should download, telling you the amount of bytes stored on the page and in the application output.

To end, just send the kill signal.

### Changelog


|Version|Change|
|---|---|
|v0.1|Initial|
|v0.2|Initial|
|v0.3|Add version const, process time to console output.|