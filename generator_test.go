package tinker

import (
	"testing"
)

const (
	fileItems       string = "test_data/items.json"
	fileAdverbs     string = "test_data/adverbs.json"
	fileCreatures   string = "test_data/creatures.json"
	fileDecorations string = "test_data/decorations.json"
	fileMaterials   string = "test_data/materials.json"
	fileQualities   string = "test_data/qualities.json"
	fileVerbs       string = "test_data/verbs.json"
)

func TestGenerator_New(t *testing.T) {
	g := New()
	if g == nil {
		t.Error("wanted non-nil value, got nil")
	}
}

func TestGenerator_LoadItemRecipes(t *testing.T) {
	tests := []struct {
		name     string
		g        *Generator
		files    []string
		wantName string
		wantLen  int
		wantErr  bool
	}{
		{
			name:     "Happy Path",
			g:        &Generator{},
			files:    []string{fileItems},
			wantName: "sword",
			wantLen:  2,
			wantErr:  false,
		},
		{
			name:     "Nil files slice",
			g:        &Generator{},
			files:    nil,
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
		{
			name:     "Empty files slice",
			g:        &Generator{},
			files:    []string{},
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.g.LoadItemRecipes(test.files...)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if len(test.g.ItemBank) != test.wantLen {
				t.Errorf("got: <%v>, want: <%v>", test.g.ItemBank, test.wantLen)
			}

			if test.wantLen <= 0 {
				return
			}

			if test.g.ItemBank[0].Name != test.wantName {
				t.Errorf("got: <%s>, want: <%s>", test.g.ItemBank[0].Name, test.wantName)
			}

		})
	}
}

func TestGenerator_LoadAttributeRecipes(t *testing.T) {
	tests := []struct {
		name     string
		g        *Generator
		files    []string
		wantName string
		wantLen  int
		wantErr  bool
	}{
		{
			name:     "Single file",
			g:        &Generator{},
			files:    []string{fileMaterials},
			wantName: "wood",
			wantLen:  6,
			wantErr:  false,
		},
		{
			name:     "Multiple files",
			g:        &Generator{},
			files:    []string{fileMaterials, fileCreatures, fileDecorations},
			wantName: "wood",
			wantLen:  10,
			wantErr:  false,
		},
		{
			name:     "Nil files slice",
			g:        &Generator{},
			files:    nil,
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
		{
			name:     "Empty files slice",
			g:        &Generator{},
			files:    []string{},
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.g.LoadAttributeRecipes(test.files...)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if len(test.g.ItemBank) != test.wantLen {
				t.Errorf("got: <%v>, want: <%v>", test.g.ItemBank, test.wantLen)
			}

			if test.wantLen <= 0 {
				return
			}

			if test.g.ItemBank[0].Name != test.wantName {
				t.Errorf("got: <%s>, want: <%s>", test.g.ItemBank[0].Name, test.wantName)
			}

		})
	}
}
