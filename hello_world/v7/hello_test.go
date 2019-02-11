package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("사람들에게 인사하기", func(t *testing.T) {
		got := Hello("Sangil", "")
		want := "Hello, Sangil"
		assertCorrectMessage(t, got, want)
	})

	t.Run("이름이 빈 문자열일 때 'Hello, World'로 프린팅하기", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("스페인어", func(t *testing.T) {
		got := Hello("Sangil", "Spanish")
		want := "Hola, Sangil"
		assertCorrectMessage(t, got, want)
	})

	t.Run("프랑스어", func(t *testing.T) {
		got := Hello("Sangil", "French")
		want := "Bongjour, Sangil"
		assertCorrectMessage(t, got, want)
	})
}
