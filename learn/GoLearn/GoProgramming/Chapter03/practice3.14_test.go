package Chapter03

import (
	"testing"
)

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB

	Kb int64 = 1000
	// Mb is byte size of megabyte
	Mb int64 = 1000 * Kb
	// Gb is byte size of gigabyte
	Gb int64 = 1000 * Mb
	// Tb is byte size of terabyte
	Tb int64 = 1000 * Gb
	// KiB is byte size of kibibyte
	KiB int64 = 1024
	// MiB is byte size of mebibyte
	MiB int64 = 1024 * KiB
	// GiB is byte size of gibibyte
	GiB int64 = 1024 * MiB
	// TiB is byte size of tebibyte
	TiB int64 = 1024 * GiB
)


func Test0314(t *testing.T) {

}
