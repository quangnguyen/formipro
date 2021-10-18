package generator

import (
	"com.nguyenonline/formipro/internal"
	"com.nguyenonline/formipro/pkg/file"
	"com.nguyenonline/formipro/pkg/model"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const texTemplate = "main.tmpl"

func GeneratePdf(templateID string, model model.Model) ([]byte, error) {
	newDirName, _ := RandDir()

	workDir := filepath.Join(internal.TmpDir, newDirName)
	err := file.CopyFiles(filepath.Join("assets/"+model.Name(), templateID), workDir)
	defer os.Rename(workDir, workDir+"_processed")

	attachments := model.GetAttachments()
	for attachmentName, attachmentBytes := range attachments {
		err = createFile(workDir, attachmentName, attachmentBytes)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		log.Printf("Could not copy files, error is '%s'\n", err)
		return nil, err
	}

	texFileName := model.Name() + ".tex"

	err = fillData(workDir, texFileName, model)
	if err != nil {
		log.Printf("Could not fill data in to placeholder, error is '%s'\n", err)
		return nil, err
	}

	err = runPdflatex(workDir, texFileName)
	if err != nil {
		log.Printf("Could not run pdflatex, error is '%s'\n", err)
		return nil, err
	}

	bytes, err := readFile(workDir, model.Name()+".pdf")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func createFile(directoryName string, fileName string, bytes []byte) error {
	attachmentFile, err := os.Create(filepath.Join(directoryName, fileName))
	defer attachmentFile.Close()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(attachmentFile.Name(), bytes, 0644)
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
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func fillData(templateDir string, texFileName string, obj interface{}) error {
	dir, _ := os.Getwd()
	templatePath := filepath.Join(dir, templateDir, texTemplate)

	data, _ := ioutil.ReadFile(templatePath)

	tpl := template.Must(template.New("main").Delims("#(", ")#").Parse(string(data)))

	texFile, err := os.Create(filepath.Join(dir, templateDir, texFileName))
	if texFile != nil {
		defer texFile.Close()
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
