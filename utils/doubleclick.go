//go:build !windows

package utils

func RunningByDoubleClick() bool {
	return false
}
