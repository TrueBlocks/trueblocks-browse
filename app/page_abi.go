// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package app

// EXISTING_CODE
import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// type OperationFunc func(types.Containerer) types.Containerer

// func Pagination[C *types.Containerer, T any](count int, offset, pageSize int) OperationFunc {
// 	return func(container types.Containerer) types.Containerer {
// 		first := base.Max(0, base.Min(offset, count-1))
// 		last := base.Min(count, first+pageSize)
// 		copy := container.ShallowCopy().(C)
// 		theSlice := container.GetItems().([]T)[first:last]
// 		copy.SetItems(theSlice)
// 		return copy
// 	}
// }

// func Organize[T any](container types.Containerer, operations ...OperationFunc) *types.Containerer {
// 	copy := (container).ShallowCopy()
// 	for _, operation := range operations {
// 		copy = operation(copy)
// 	}
// 	return &copy
// }

// EXISTING_CODE

func (a *App) AbiPage(first, pageSize int) *types.AbiContainer {
	// EXISTING_CODE
	// EXISTING_CODE

	a.abis.CollateAndFilter()
	first = base.Max(0, base.Min(first, len(a.abis.Items)-1))
	last := base.Min(len(a.abis.Items), first+pageSize)
	copy, _ := a.abis.ShallowCopy().(*types.AbiContainer)
	copy.Items = a.abis.Items[first:last]
	return copy
}
