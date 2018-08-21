package di_test

import (
	"github.com/ts4z/di"
	"log"
	"testing"
)

type noDeps struct{}

type Fooer interface{}
type Barer interface{}

type settableDeps struct {
	Foo Fooer
	Bar Barer
}

func TestHappies(t *testing.T) {
	cases := []interface{}{
		noDeps{},
		settableDeps{
			Foo: &struct{}{},
			Bar: &struct{}{},
		},
	}
	for i, c := range cases {
		e := di.EnsureAllFieldsSet(c)
		if e != nil {
			t.Errorf("expected error, got %v in case %d", e, i)
		}

		// Also verify we are not panicking for these
		// same values.
		di.PanicUnlessAllFieldsSet(c)
	}
}

func TestSads(t *testing.T) {
	cases := []settableDeps{
		settableDeps{Bar: &struct{}{}},
		settableDeps{Foo: &struct{}{}},
		settableDeps{},
	}
	for i, c := range cases {
		e := di.EnsureAllFieldsSet(c)
		if e == nil {
			t.Errorf("expected error, got %v in case %d", e, i)
		}
		log.Printf("ok, expected error found: %v\n", e)
	}
}

func TestPanicky(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic as expected")
		}
	}()

	c := settableDeps{}
	di.PanicUnlessAllFieldsSet(c)
}
