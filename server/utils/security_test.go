package utils_test

import (
	"testing"

	"github.com/armoredboar/account-api/server/utils"
)

func TestCalculateHmacSha256(t *testing.T) {

	expected := "d3c56d918f80f85be5c70ce6e1360aabb2bed96b3578d64480d459e71266322e"
	actual := utils.CalculateHmacSha256("sample data", "sample key")

	if *actual != expected {
		t.Error("Expected " + expected)
	}
}

func TestCalculateHmacSha256_NullKey(t *testing.T) {

	actual := utils.CalculateHmacSha256("sample data", "")

	if actual != nil {
		t.Error("Expected nil")
	}
}

func TestCalculateHmacSha256_NullData(t *testing.T) {

	actual := utils.CalculateHmacSha256("", "sample key")

	if actual != nil {
		t.Error("Expected nil")
	}
}
