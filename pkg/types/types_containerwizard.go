// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"errors"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type WizardContainer struct {
	Chain   string          `json:"chain"`
	Items   []WizError      `json:"items"`
	NItems  uint64          `json:"nItems"`
	Updater updater.Updater `json:"updater"`
	Sorts   sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	State WizState `json:"state"`
	// EXISTING_CODE
}

func NewWizardContainer(chain string, itemsIn []WizError) WizardContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := WizardContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewWizardUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	ret.State = WizWelcome
	// EXISTING_CODE
	return ret
}

func NewWizardUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Duration: 2 * time.Minute, Type: updater.Timer},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("wizard", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *WizardContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *WizardContainer) GetItems() interface{} {
	return s.Items
}

func (s *WizardContainer) SetItems(items interface{}) {
	s.Items = items.([]WizError)
}

func (s *WizardContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *WizardContainer) ShallowCopy() Containerer {
	ret := &WizardContainer{
		Chain:   s.Chain,
		NItems:  s.NItems,
		Updater: s.Updater,
		// EXISTING_CODE
		State: s.State,
		// EXISTING_CODE
	}
	return ret
}

func (s *WizardContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *WizardContainer) passesFilter(item *WizError, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		_ = item
		// EXISTING_CODE
	}
	return
}

func (s *WizardContainer) Accumulate(item *WizError) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *WizardContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *WizardContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("wizard") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *WizError, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []WizError{}
	s.ForEveryItem(func(item *WizError, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *WizardContainer) ForEveryItem(process EveryWizErrorFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
