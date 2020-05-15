package tinker

import (
	"testing"
)

const (
	// Item recipes
	fileWeapons string = "test_data/weapons.json"
	fileArt     string = "test_data/art.json"

	// Attribute recipes
	fileAdverbs     string = "test_data/adverbs.json"
	fileCreatures   string = "test_data/creatures.json"
	fileDecorations string = "test_data/decorations.json"
	fileMaterials   string = "test_data/materials.json"
	fileQualities   string = "test_data/qualities.json"

	// Verbs
	fileVerbs string = "test_data/verbs.json"

	// Edge cases
	testFileBlank     string = "test_data/blank.json"
	testFileEmptyJSON string = "test_data/empty.json"
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
			name:     "Single file",
			g:        &Generator{},
			files:    []string{fileWeapons},
			wantName: "sword",
			wantLen:  2,
			wantErr:  false,
		},
		{
			name:     "Multiple files",
			g:        &Generator{},
			files:    []string{fileWeapons, fileArt},
			wantName: "sword",
			wantLen:  3,
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
		{
			name:     "Empty JSON",
			g:        &Generator{},
			files:    []string{testFileEmptyJSON},
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
		{
			name:     "Blank file",
			g:        &Generator{},
			files:    []string{testFileBlank},
			wantName: "",
			wantLen:  0,
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.g.LoadItemRecipes(test.files...)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if len(test.g.ItemBank) != test.wantLen {
				t.Errorf("got: <%v>, want: <%v>", len(test.g.ItemBank), test.wantLen)
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
		{
			name:     "Empty JSON",
			g:        &Generator{},
			files:    []string{testFileEmptyJSON},
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
		{
			name:     "Blank file",
			g:        &Generator{},
			files:    []string{testFileBlank},
			wantName: "",
			wantLen:  0,
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.g.LoadAttributeRecipes(test.files...)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if len(test.g.AtbBank) != test.wantLen {
				t.Errorf("got: <%v>, want: <%v>", len(test.g.AtbBank), test.wantLen)
			}

			if test.wantLen <= 0 {
				return
			}

			if test.g.AtbBank[test.wantName].Name != test.wantName {
				t.Errorf("got: <%s>, want: <%s>", test.g.AtbBank[test.wantName].Name, test.wantName)
			}

		})
	}
}

func TestGenerator_LoadVerbs(t *testing.T) {
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
			files:    []string{fileVerbs},
			wantName: "made of",
			wantLen:  3,
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
		{
			name:     "Empty JSON",
			g:        &Generator{},
			files:    []string{testFileEmptyJSON},
			wantName: "",
			wantLen:  0,
			wantErr:  false,
		},
		{
			name:     "Blank file",
			g:        &Generator{},
			files:    []string{testFileBlank},
			wantName: "",
			wantLen:  0,
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.g.LoadVerbs(test.files...)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if len(test.g.VerbBank) != test.wantLen {
				t.Errorf("got: <%v>, want: <%v>", len(test.g.VerbBank), test.wantLen)
			}

			if test.wantLen <= 0 {
				return
			}

			if test.g.VerbBank[test.wantName].Name != test.wantName {
				t.Errorf("got: <%s>, want: <%s>", test.g.VerbBank[test.wantName].Name, test.wantName)
			}

		})
	}
}

func TestGenerator_Item(t *testing.T) {
	tests := []struct {
		name    string
		g       *Generator
		items   []string
		atbs    []string
		verbs   []string
		wantErr bool
	}{
		{
			name:    "No missing data",
			g:       &Generator{},
			items:   []string{fileWeapons, fileArt},
			atbs:    []string{fileAdverbs, fileCreatures, fileDecorations, fileMaterials, fileQualities},
			verbs:   []string{fileVerbs},
			wantErr: false,
		},
		{
			name:    "Missing item recipes",
			g:       &Generator{},
			items:   []string{},
			atbs:    []string{fileAdverbs, fileCreatures, fileDecorations, fileMaterials, fileQualities},
			verbs:   []string{fileVerbs},
			wantErr: true,
		},
		{
			name:    "Missing attributes",
			g:       &Generator{},
			items:   []string{fileWeapons, fileArt},
			atbs:    []string{},
			verbs:   []string{fileVerbs},
			wantErr: false,
		},
		{
			name:    "Missing verbs",
			g:       &Generator{},
			items:   []string{fileWeapons, fileArt},
			atbs:    []string{fileAdverbs, fileCreatures, fileDecorations, fileMaterials, fileQualities},
			verbs:   []string{},
			wantErr: false,
		},
		{
			name:    "Missing attributes and verbs",
			g:       &Generator{},
			items:   []string{fileWeapons, fileArt},
			atbs:    []string{},
			verbs:   []string{},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.g.LoadItemRecipes(test.items...); err != nil {
				t.Fatal(err)
			}
			if err := test.g.LoadAttributeRecipes(test.atbs...); err != nil {
				t.Fatal(err)
			}
			if err := test.g.LoadVerbs(test.verbs...); err != nil {
				t.Fatal(err)
			}

			got, err := test.g.Item()
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}

			if test.wantErr {
				return
			}

			if got.Name == "" {
				t.Errorf("got: empty name, want: non-empty name")
			}

			if got.Description == "" {
				t.Errorf("got: empty description, want: non-empty description")
			}
		})
	}
}
