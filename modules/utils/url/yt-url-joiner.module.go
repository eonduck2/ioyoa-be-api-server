package url

import (
	"ioyoa/modules/utils/url"
	"testing"
)

func TestYtUrlJoiner(t *testing.T) {
	first := "test.com"
	second := "hola"
	third := "cool"

	// 예상 결과
	expected := "test.com/hola/cool"

	// UrlJoiner 함수 호출
	result := url.UrlJoiner(first, second, third)

	// 결과 확인
	if result != expected {
		t.Errorf("UrlJoiner() = %v, want %v", result, expected)
	}
}