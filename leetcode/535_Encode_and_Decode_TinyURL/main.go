package main

import "strconv"

type Codec struct {
	decoder map[int]string
	counter int
}

func Constructor() Codec {
	return Codec{decoder: make(map[int]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.decoder[this.counter] = longUrl
	this.counter++
	return strconv.Itoa(this.counter - 1)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	number, _ := strconv.Atoi(shortUrl)
	return this.decoder[number]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
