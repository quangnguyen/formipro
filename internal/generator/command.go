package generator

import (
	"com.nguyenonline/formipro/pkg/file"
	"com.nguyenonline/formipro/pkg/model"
	"com.nguyenonline/formipro/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const texTemplate = "main.tmpl"

func PdfLatex(templateID string, m model.Model) ([]byte, error) {
	newDirName, _ := util.RandDir()

	workDir := filepath.Join("tmp", newDirName)
	defer func(oldPath, newPath string) {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Println(err)
		}
	}(workDir, workDir+"_processed")

	err := file.CopyFiles(filepath.Join("assets/"+m.Name(), templateID), workDir)
	if err != nil {
		log.Printf("Could not copy files, error is '%s'\n", err)
		return nil, err
	}

	attachments := m.GetAttachments()
	for attachmentName, attachmentBytes := range attachments {
		err = createFile(workDir, attachmentName, attachmentBytes)
		if err != nil {
			return nil, err
		}
	}

	texFileName := m.Name() + ".tex"

	err = fillData(workDir, texFileName, m)
	if err != nil {
		log.Printf("Could not fill data in to placeholder, error is '%s'\n", err)
		return nil, err
	}

	err = runPdflatex(workDir, texFileName)
	if err != nil {
		log.Printf("Could not run pdflatex, error is '%s'\n", err)
		return nil, err
	}

	bytes, err := readFile(workDir, m.Name()+".Pdf")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func createFile(directoryName string, fileName string, bytes []byte) error {
	attachmentFile, err := os.Create(filepath.Join(directoryName, fileName))
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(attachmentFile)
	if err != nil {
		return err
	}
	err = os.WriteFile(attachmentFile.Name(), bytes, 0644)
	return err
}

func runPdflatex(directory string, texFile string) error {
	cmd := exec.Command("pdflatex", texFile)
	dir, _ := os.Getwd()
	cmd.Dir = filepath.Join(dir, directory)
	err := cmd.Run()
	return err
}

func readFile(baseDir string, fileName string) ([]byte, error) {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, baseDir, fileName)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func fillData(templateDir string, texFileName string, obj interface{}) error {
	dir, _ := os.Getwd()
	templatePath := filepath.Join(dir, templateDir, texTemplate)

	data, _ := os.ReadFile(templatePath)

	tpl := template.Must(template.New("main").Delims("#(", ")#").Parse(string(data)))

	texFile, err := os.Create(filepath.Join(dir, templateDir, texFileName))
	if texFile != nil {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println(err)
			}
		}(texFile)
	}
	if err != nil {
		return err
	}

	err = tpl.Execute(texFile, obj)
	if err != nil {
		return err
	}
	return nil
}

func HTMLToLatex(input string) (string, error) {
	cmdString := "echo \"" + input + "\" | pandoc -f html -t latex"
	cmd := exec.Command("sh", "-c", cmdString)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Could not convert html to latex, error is '%s'\n", err.Error())
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}
