file-bucket
===========

Post files to me and I'll save them!

## Why? ##
This is little program is virtually useless, but it helped me learn some Go, which is not useless!

## Installation ##
~~~
go get github.com/ccraig/file-bucket
~~~

Then cd to the project's source directory and run it with your upload path as it's only parameter
~~~
cd $GOPATH/src/github.com/ccraig/file-bucket
go run file-bucket /path/to/uploads
~~~

You may now make Multipart HTTP POST requests to port 1337 using the parameter name "file". It will then write that file to your uploads directory.
