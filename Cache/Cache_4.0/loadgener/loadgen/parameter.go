package loadgen

import (
	"bytes"
	"errors"
	"fmt"
	"loadgener/loadgen/lib"
	"strings"
	"time"
)

type ParamSet struct {
	Caller     lib.Caller
	TimeoutNS  time.Duration
	LPS        uint32
	DurationNS time.Duration
	ResultCh   chan *lib.CallResult
}

func (pset *ParamSet) Check() error {
	var errMsgs []string

	if pset.Caller == nil {
		errMsgs = append(errMsgs, "Invalid caller!")
	}
	if pset.TimeoutNS == 0 {
		errMsgs = append(errMsgs, "Invalid timeoutNS!")
	}
	if pset.LPS == 0 {
		errMsgs = append(errMsgs, "Invalid lps(load per second)!")
	}
	if pset.DurationNS == 0 {
		errMsgs = append(errMsgs, "Invalid durationNS!")
	}
	if pset.ResultCh == nil {
		errMsgs = append(errMsgs, "Invalid result channel!")
	}

	var buf bytes.Buffer

	buf.WriteString("Checking the ParamSet...")
	if errMsgs != nil {
		errMsg := strings.Join(errMsgs, " ")
		buf.WriteString(fmt.Sprintf("没有通过检查(%s)", errMsg))
		logger.Infoln(buf.String())
		return errors.New(errMsg)
	}

	buf.WriteString(
		fmt.Sprintf("检查通过. (timeoutNS=%s, lps=%d, durationNS=%s)",
			pset.TimeoutNS, pset.LPS, pset.DurationNS))

	logger.Infoln(buf.String())
	return nil

}
