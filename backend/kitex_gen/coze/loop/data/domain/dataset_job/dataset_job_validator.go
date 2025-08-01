// Code generated by Validator v0.2.6. DO NOT EDIT.

package dataset_job

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *JobLog) IsValid() error {
	return nil
}
func (p *DatasetIOFile) IsValid() error {
	if p.Provider.String() == "<UNSET>" {
		return fmt.Errorf("field Provider defined_only rule failed")
	}
	if len(p.Path) < int(1) {
		return fmt.Errorf("field Path min_len rule failed, current value: %d", len(p.Path))
	}
	return nil
}
func (p *DatasetIODataset) IsValid() error {
	return nil
}
func (p *DatasetIOEndpoint) IsValid() error {
	if p.File != nil {
		if err := p.File.IsValid(); err != nil {
			return fmt.Errorf("field File not valid, %w", err)
		}
	}
	if p.Dataset != nil {
		if err := p.Dataset.IsValid(); err != nil {
			return fmt.Errorf("field Dataset not valid, %w", err)
		}
	}
	return nil
}
func (p *DatasetIOJob) IsValid() error {
	if p.Source != nil {
		if err := p.Source.IsValid(); err != nil {
			return fmt.Errorf("field Source not valid, %w", err)
		}
	}
	if p.Target != nil {
		if err := p.Target.IsValid(); err != nil {
			return fmt.Errorf("field Target not valid, %w", err)
		}
	}
	if p.Option != nil {
		if err := p.Option.IsValid(); err != nil {
			return fmt.Errorf("field Option not valid, %w", err)
		}
	}
	if p.Progress != nil {
		if err := p.Progress.IsValid(); err != nil {
			return fmt.Errorf("field Progress not valid, %w", err)
		}
	}
	return nil
}
func (p *DatasetIOJobOption) IsValid() error {
	return nil
}
func (p *DatasetIOJobProgress) IsValid() error {
	return nil
}
func (p *FieldMapping) IsValid() error {
	if len(p.Source) < int(1) {
		return fmt.Errorf("field Source min_len rule failed, current value: %d", len(p.Source))
	}
	if len(p.Target) < int(1) {
		return fmt.Errorf("field Target min_len rule failed, current value: %d", len(p.Target))
	}
	return nil
}
