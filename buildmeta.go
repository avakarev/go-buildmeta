package buildmeta

import (
	"runtime"
	"time"

	"github.com/avakarev/go-timeutil"
)

// Build meta information, populated at build-time
var (
	// Commit is git commit sha
	Commit string
	// Compiler is Go compiler version
	Compiler string
	// Date is a build datetime in UTC
	Date string
	// LocalDate is a build datetime in `TZ`
	LocalDate string
	// Ref is git branch or tag ref
	Ref string
	// Uptime is the application uptime
	Uptime string
)

// Fields returns build meta as map
func Fields() map[string]interface{} {
	return map[string]interface{}{
		"commit":   Commit,
		"ref":      Ref,
		"compiler": Compiler,
		"date":     LocalDate,
		"uptime":   Uptime,
	}
}

// Init initializes buildmeta
func Init() error {
	if Commit == "" {
		Commit = "n/a"
	}
	Compiler = runtime.Version()
	if Date == "" {
		Date = "n/a"
		LocalDate = "n/a"
	}
	if Date != "n/a" {
		t, err := time.Parse(time.RFC3339, Date)
		if err != nil {
			return err
		}
		if !t.IsZero() {
			LocalDate = timeutil.Local(t).Format(time.RFC3339)
		}
	}
	if Ref == "" {
		Ref = "n/a"
	}
	Uptime = timeutil.Local(time.Now()).Format(time.RFC3339)
	return nil
}

// MustInit is like Init but panics in case of error
func MustInit() {
	if err := Init(); err != nil {
		panic("buildmeta.MustInit(): " + err.Error())
	}
}

func init() {
	MustInit()
}
