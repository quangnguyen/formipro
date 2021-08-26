package job

import (
	"com.nguyenonline/formipro/internal"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const timeToLiveMinutes = 1

func DeleteProcessedTmpFiles() {
	cronJob := cron.New()
	_, err := cronJob.AddFunc("@every 1m", delete)
	if err != nil {
		log.Println(err.Error())
	}
	cronJob.Run()
	defer cronJob.Stop()
}

func delete() {
	log.Println("Checking and deleting processed tmp files...")
	files, err := ioutil.ReadDir(internal.TmpDir)
	if err != nil {
		log.Println(err.Error())
	}

	now := time.Now()
	deleteCount := 0
	for _, file := range files {
		if file.IsDir() && strings.HasSuffix(file.Name(), "processed") && file.ModTime().Before(now.Add(-1*timeToLiveMinutes*time.Minute)) {
			err = os.RemoveAll(filepath.Join(internal.TmpDir, file.Name()))
			if err != nil {
				log.Printf("Could not delete folder '%s', error is '%s'\n", file.Name(), err.Error())
			} else {
				deleteCount++
			}
		}
	}
	log.Printf("Deleted processed files of '%d' tmp folders\n", deleteCount)
}
