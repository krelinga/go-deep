package exam_test

import (
	"testing"

	"github.com/krelinga/go-deep"
	"github.com/krelinga/go-deep/exam"
)

func TestHarness(t *testing.T) {
	t.Run("env propagation", func(t *testing.T) {
		h := exam.Harness{
			DeepEnv: deep.NewEnv(testValOpt("foo")),
		}
		h.Run(func(innerT exam.T) {
			val := getTestVal(t, innerT.DeepEnv())
			if val != "foo" {
				t.Errorf("expected test value 'foo', got %q", val)
			}
			innerT.Run("nested", func(innerT exam.T) {
				val := getTestVal(t, innerT.DeepEnv())
				if val != "foo" {
					t.Errorf("expected test value 'foo', got %q", val)
				}
			})
		})
	})
}
