// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package ctl

func newMonitor() shutdownWatcher {
	return &shutdownWatcherFreebsd{}
}

// TODO: copied Darwin watcher due to simplicity, need to verify if this is
// the best approach for freebsd or if we should move to something similar to the Linux version
type shutdownWatcherFreebsd struct {
}

func (s *shutdownWatcherFreebsd) checkShutdownStatus(shutdown chan<- shutdownCmd) {
	shutdown <- shutdownCmd{noop: true}
}

func (s *shutdownWatcherFreebsd) init() (err error) {
	return err
}

func (s *shutdownWatcherFreebsd) stop() {}
