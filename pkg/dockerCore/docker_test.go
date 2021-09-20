package dockerCore

import "testing"

func TestLoadImage(t *testing.T) {
	Init()
	if err := LoadImage("/tmp/951b7320-77db-086f-4e0e-88b3869ba7ac.tar"); err != nil {
		panic(err)
	}
}
