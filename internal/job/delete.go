package job

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var timeToLiveSeconds = 60

func Run() {
	loadParams()

	cronJob := cron.New()
	_, err := cronJob.AddFunc("@every 1m", deleteTmpDir)
	if err != nil {
		log.Println(err.Error())
	}
	cronJob.Run()
	defer cronJob.Stop()
}

func loadParams() {
	ttlSecondString := os.Getenv("TMP_FOLDER_TTL_SECOND")
	ttl, err := strconv.ParseInt(ttlSecondString, 10, 64)
	if err == nil {
		timeToLiveSeconds = int(ttl)
	}
}

func deleteTmpDir() {
	dirEntries, err := os.ReadDir("tmp")
	if err != nil {
		log.Println(err.Error())
	}

	now := time.Now()
	deleted := 0
	for _, entry := range dirEntries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), "processed") {
			info, _ := entry.Info()

			if info.ModTime().Before(now.Add(time.Duration(-1 * timeToLiveSeconds))) {
				err = os.RemoveAll(filepath.Join("tmp", info.Name()))
				if err != nil {
					log.Printf("Could not delete tmp folder '%s', error is '%s'\n", info.Name(), err.Error())
				} else {
					deleted++
				}
			}
		}
	}
	if deleted > 0 {
		log.Printf("Deleted %d processed tmp folders\n", deleted)
	}
}
