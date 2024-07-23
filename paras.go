package logger

var (
	// MaxAgeDay ...
	MaxAgeDay = 30
	// MaxSizeMB ...
	MaxSizeMB = 50
	// MaxBackupCnt ...
	MaxBackupCnt = 100
)

// GetMaxAgeDay ...
// @Description:
// @return int
func GetMaxAgeDay() int {
	return MaxAgeDay
}

// SetMaxAgeDay ...
// @Description:
// @param v
func SetMaxAgeDay(v int) {
	MaxAgeDay = v
}

// GetMaxSizeMB ...
// @Description:
// @return int
func GetMaxSizeMB() int {
	return MaxSizeMB
}

// SetMaxSizeMB ...
// @Description:
// @param v
func SetMaxSizeMB(v int) {
	MaxSizeMB = v
}

// GetMaxBackupCnt ...
// @Description:
// @return int
func GetMaxBackupCnt() int {
	return MaxBackupCnt
}

// SetMaxBackupCnt ...
// @Description:
// @param v
func SetMaxBackupCnt(v int) {
	MaxBackupCnt = v
}

// Params ...
// @Description:
type Params struct {
	RemovePathPrefix string
}
