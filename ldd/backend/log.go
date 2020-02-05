package main

import (
	"fmt"
	"time"
)

/* Log */

type debugLevel int

const (
    Debug  debugLevel = iota
    Info
    Warning
    Critical
)

var DEBUG_LEVEL debugLevel = Info

func logEvent(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s - %s\n", time.Now().Format(time.RFC850), msg)
    }
}

func logDetail(msg string, level debugLevel){
    if (level >= DEBUG_LEVEL){
        fmt.Printf("%s\n%s\n", time.Now().Format(time.RFC850), msg)
    }
}