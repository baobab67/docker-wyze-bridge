package go2rtcmgr

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

// SNAPSHOT_ROTATE turns JPEG snapshots per camera, for cameras mounted
// sideways (the Wyze Video Doorbell v1's sensor is landscape on its
// side). Format: "cam:deg,cam:deg" with deg 90, 180 or 270, e.g.
// "front_doorbell:90". go2rtc's frame.jpeg does the turn (ffmpeg).
// Only snapshots (the MQTT/HA camera image) are turned; the live
// streams are not.
var (
	rotateOnce sync.Once
	rotateMap  map[string]int
)

func parseRotate(v string) map[string]int {
	m := map[string]int{}
	for _, part := range strings.Split(v, ",") {
		name, deg, ok := strings.Cut(strings.TrimSpace(part), ":")
		if !ok {
			continue
		}
		d, err := strconv.Atoi(strings.TrimSpace(deg))
		if err != nil || (d != 90 && d != 180 && d != 270) {
			continue
		}
		m[strings.ToLower(strings.TrimSpace(name))] = d
	}
	return m
}

// SnapshotRotation returns the configured snapshot rotation for a
// stream name, 0 when none.
func SnapshotRotation(name string) int {
	rotateOnce.Do(func() { rotateMap = parseRotate(os.Getenv("SNAPSHOT_ROTATE")) })
	return rotateMap[strings.ToLower(name)]
}
