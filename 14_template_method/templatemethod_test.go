package templatemethod

import (
	"testing"
)

func TestExampleHTTPDownloader(t *testing.T) {
	var downloader Downloader = NewHTTPDownloader()

	downloader.Download("http://example.com/abc.zip")
	// Output:
	// prepare downloading
	// download http://example.com/abc.zip via http
	// http save
	// finish downloading
}

func TestExampleFTPDownloader(t *testing.T) {
	var downloader Downloader = NewFTPDownloader()

	downloader.Download("ftp://example.com/abc.zip")
	// Output:
	// prepare downloading
	// download ftp://example.com/abc.zip via ftp
	// default save
	// finish downloading
}

func TestExampleUserDaoService(t *testing.T) {
	userDaoService := NewUserDaoService()
	var user = &User{
		name:     "test",
		password: "123",
	}

	userDaoService.Add(user)
}

func TestExampleShapes(t *testing.T) {
	var circle = NewCircleDrawer("blue")
	circle.Draw()

	var rec = NewRectangleDrawer("green")
	rec.Draw()
}
