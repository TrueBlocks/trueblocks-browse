// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type WizardContainer struct {
	Chain   string       `json:"chain"`
	Items   []WizError   `json:"items"`
	NItems  uint64       `json:"nItems"`
	State   WizState     `json:"state"`
	Updater sdk.Updater  `json:"updater"`
	Sorts   sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
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

func NewWizardUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Duration: 2 * time.Minute, Type: sdk.Timer},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("wizard", items)
	if reset {
		u.Reset()
	}
	return u
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
		State:   s.State,
		Updater: s.Updater,
		// EXISTING_CODE
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

func (s *WizardContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

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
		if s.passesFilter(item, filter) {
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

func (s *WizardContainer) Sort() (err error) {
	// EXISTING_CODE
	// TODO: Sorting?
	// EXISTING_CODE
	return
}

// EXISTING_CODE
func (s *WizardContainer) GetWizChain() string {
	return s.Chain
}

func (s *WizardContainer) GetWizState() WizState {
	return s.State
}

func (s *WizardContainer) SetWizChain(chain string) {
	s.Chain = chain
}

func (s *WizardContainer) SetWizState(state WizState) {
	s.State = state
}

// EXISTING_CODE
