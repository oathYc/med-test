package test

import (
	"hello/boot"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := boot.Bootstrap("../config/config.toml"); err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
