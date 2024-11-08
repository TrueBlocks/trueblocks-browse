package app

func (a *App) GetDeferredErrors() []string {
	var errs []string
	for _, err := range a.deferredErrors {
		errs = append(errs, err.Error())
	}
	return errs
}
