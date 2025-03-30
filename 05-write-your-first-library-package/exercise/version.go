package goversion

import "runtime"

func GetGoVersion() string {
	return runtime.Version()
}
