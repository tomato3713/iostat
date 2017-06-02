// +build darwin

package iostat

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// #include <stdint.h>
// #include "iostat.h"
import "C"
import (
	"time"
)

func ReadDriveStats() ([]*DriveStats, error) {
	var buf [C.NDRIVE]C.DriveStats
	n, err := C.readdrivestat(&buf[0], C.int(len(buf)))
	if err != nil {
		return nil, err
	}
	stats := make([]*DriveStats, n)
	for i := 0; i < int(n); i++ {
		stats[i] = &DriveStats{
			Name:         C.GoString(&buf[i].name[0]),
			Size:         int64(buf[i].size),
			BlockSize:    int64(buf[i].blocksize),
			BytesRead:    int64(buf[i].read),
			BytesWritten: int64(buf[i].written),
			NumReads:     int64(buf[i].nread),
			NumWrites:    int64(buf[i].nwrite),
			ReadLatency:  time.Duration(buf[i].readtime),
			WriteLatency: time.Duration(buf[i].writetime),
		}
	}
	return stats, nil
}
