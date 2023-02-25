package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/go-units"
)

type SizeValue struct {
	Valid bool
	Range bool
	Min   int64
	Max   int64
}

func ParseSizeValue(dest *SizeValue, strSize string, maxSize int64) error {
	dest.Valid = false
	dest.Range = false
	dest.Min = 0
	dest.Max = 0

	if strSize == "" {
		return nil
	}

	var err error

	parts := strings.Split(strSize, ":")

	dest.Max, err = units.RAMInBytes(parts[0])
	if err != nil {
		return err
	}

	if len(parts) > 1 {
		dest.Range = true
		dest.Min = dest.Max
		dest.Max, err = units.RAMInBytes(parts[1])
		if err != nil {
			return err
		}
	}

	if dest.Max < dest.Min {
		return fmt.Errorf("invalid size: %q", strSize)
	}
	if dest.Max > maxSize {
		return fmt.Errorf("max size of %v exceeded", maxSize)
	}

	dest.Valid = true
	return nil
}

type DurationValue struct {
	Valid bool
	Range bool
	Min   time.Duration
	Max   time.Duration
}

func ParseDurationValue(dest *DurationValue, strDelay string, maxDelay time.Duration) error {
	dest.Valid = false
	dest.Range = false
	dest.Min = time.Duration(0)
	dest.Max = time.Duration(0)

	if strDelay == "" {
		return nil
	}

	var err error

	parts := strings.Split(strDelay, ":")

	dest.Max, err = time.ParseDuration(parts[0])
	if err != nil {
		return err
	}

	if len(parts) > 1 {
		dest.Range = true
		dest.Min = dest.Max
		dest.Max, err = time.ParseDuration(parts[1])
		if err != nil {
			return err
		}
	}

	if dest.Max < dest.Min {
		return fmt.Errorf("invalid range: %q", strDelay)
	}
	if dest.Max > maxDelay {
		return fmt.Errorf("max delay of %v exceeded", maxDelay)
	}

	dest.Valid = true
	return nil
}
