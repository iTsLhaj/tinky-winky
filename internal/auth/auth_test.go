package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Empty Authorization Header", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Authorization", "")

		_, e := GetAPIKey(r.Header)
		assertError(t, e)
	})
	t.Run("Invalid Authorization Header #1", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Authorization", "invalid")

		_, e := GetAPIKey(r.Header)
		assertError(t, e)
	})
	t.Run("Invalid Authorization Header #2", func(t *testing.T) {
		// FIXME: i intentionally broke the test
		// 		  so i can make sure that my CI fail
		//		  when the test dont pass
		t.Errorf("temporary failure")

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Authorization", "0xdeadbeef meow/invalid ApiKey O_O")

		_, e := GetAPIKey(r.Header)
		assertError(t, e)
	})
	t.Run("Invalid Authorization Header #2", func(t *testing.T) {
		mockKey := "123456"
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", mockKey))

		key, e := GetAPIKey(r.Header)
		if e != nil {
			t.Errorf("no error expected, got %v", e)
		}
		if key != mockKey {
			t.Errorf("expected %s, got %s", mockKey, key)
		}
	})
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("error expected got %v", err)
	}
}
