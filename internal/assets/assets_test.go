package assets_test

import (
	"testing"

	"github.com/wtask-go/gqlfront/internal/assets"
)

func TestAssets_Playground(t *testing.T) {
	if _, err := assets.Playground(); err != nil {
		t.Error(err)
	}
}
