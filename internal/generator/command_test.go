package generator

import (
	"com.nguyenonline/formipro/pkg/model"
	"reflect"
	"testing"
)

func TestGeneratePdf(t *testing.T) {
	type args struct {
		templateID string
		model      model.Model
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePdf(tt.args.templateID, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePdf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePdf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillData(t *testing.T) {
	type args struct {
		templateDir string
		texFile     string
		obj         interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fillData(tt.args.templateDir, tt.args.texFile, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("fillData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		baseDir  string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.baseDir, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HtmlToLatex(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Strong", args{"<strong>Hello</strong>"}, "\\textbf{Hello}", false},
		{"Paragraph", args{"<p>Hello</p>"}, "Hello", false},
		{"Italic", args{"<i>Hello</i>"}, "\\emph{Hello}", false},
		{"Italic and bold", args{"<i>Hello <b>World</b></i>"}, "\\emph{Hello \\textbf{World}}", false},
		{"Italic and bold", args{"<u>All</u>"}, "\\underline{All}", false},
		{"Special characters", args{"<p>~!@#$%^&*()_+-=[]{};:'<>,.</p>"}, "\\textasciitilde!@\\#\\$\\%\\^{}\\&*()\\_+-={[}{]}\\{\\};:'\\textless\\textgreater,.", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HTMLToLatex(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("runPandocHtmlToLatex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runPandocHtmlToLatex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runPdflatex(t *testing.T) {
	type args struct {
		directory string
		texFile   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runPdflatex(tt.args.directory, tt.args.texFile); (err != nil) != tt.wantErr {
				t.Errorf("runPdflatex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
