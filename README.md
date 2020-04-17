# url-shortener
Url shortener code test for a job interview

Part of the hiring process was to build a url shortener in Rails and React. However I decided to built it also in Go for practice.

The code is divided just in case somebody may want to use some part of it. The 'algorithm' to generate the code for the short url is separated.

As for the storage I decided to go for a very simple array of structs. Because in the rare case that somebody may want to use this. I didn't want
to get specific for the storage since they may want to use something very different. The functions that interact with the storage are very
short, easy to understand and on the top of the main.go file. So, modifying the functions to use any storage should be very simple.

