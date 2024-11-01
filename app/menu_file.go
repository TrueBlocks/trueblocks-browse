package app

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (a *App) FileNew(cd *menu.CallbackData) {
	// a.projects = types.NewProjectContainer("Untitled.tbx", []base.Address{})
	// messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
	// 	String1: "/",
	// })
	// messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
	// 	String1: a.Filename,
	// 	String2: "Created",
	// })
}

func (a *App) FileOpen(cd *menu.CallbackData) {
	// fn, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
	// 	DefaultDirectory:           "/Users/jrush/Documents/",
	// 	Title:                      "Open File",
	// 	CanCreateDirectories:       true,
	// 	ShowHiddenFiles:            false,
	// 	ResolvesAliases:            false,
	// 	TreatPackagesAsDirectories: false,
	// 	Filters: []runtime.FileFilter{
	// 		{DisplayName: "Monitor Groups", Pattern: "*.tbx"},
	// 	},
	// })

	// if len(fn) > 0 {
	// 	save := a.FreshenController.Sleep
	// 	defer func() { a.FreshenController.Sleep = save }()
	// 	a.FreshenController.Sleep = 1000
	// 	a.session.LastFile = fn
	// 	a.saveSession()

	// 	a.CancelAllContexts()
	// 	a.projects = types.NewProjectContainer(a.session.LastChain, []base.Address{})
	// 	newProject := types.ProjectContainer{}
	// 	if err := newProject.Load(fn); err != nil {
	// 		logger.Error("Error loading project", err.Error(), a.session.Session.LastFile)
	// 		newProject.Session = a.session.Session // backup plan
	// 		newProject.Session.Wizard.State = coreTypes.Error
	// 	}
	// 	copyFields := func(dest, src *coreTypes.Session) {
	// 		dest.LastChain = src.LastChain
	// 		dest.LastFile = src.LastFile
	// 		dest.LastFolder = src.LastFolder
	// 		dest.LastRoute = src.LastRoute
	// 		dest.LastSub = src.LastSub
	// 		dest.Toggles = src.Toggles
	// 		dest.Window = src.Window
	// 		dest.Wizard = src.Wizard
	// 	}
	// 	copyFields(&a.session.Session, &newProject.Session)

	// 	a.session = types.NewSessionContainer(a.Chain, &a.session.Session)
	// 	var wg sync.WaitGroup
	// 	for _, address := range newProject.Items {
	// 		wg.Add(1)
	// 		go a.loadHistory(address, &wg, nil)
	// 	}
	// 	wg.Wait()

	// 	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
	// 		String1: "/",
	// 	})
	// 	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
	// 		String1: a.Filename,
	// 		String2: "Opened",
	// 	})
	// } else {
	// 	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
	// 		String1: a.Filename,
	// 		String2: "Not opened",
	// 	})
	// }
}

func (a *App) FileSave(cd *menu.CallbackData) {
	// a.saveSession()
	// fn, _ := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
	// 	DefaultDirectory:           a.session.Session.LastFolder,
	// 	DefaultFilename:            a.Filename,
	// 	Title:                      "Save File",
	// 	CanCreateDirectories:       true,
	// 	ShowHiddenFiles:            false,
	// 	TreatPackagesAsDirectories: false,
	// 	Filters: []runtime.FileFilter{
	// 		{DisplayName: "Projects", Pattern: "*.tbx"},
	// 	},
	// })
	// if len(fn) > 0 {
	// 	a.session.Session.LastFolder, a.session.Session.LastFile = filepath.Split(fn)
	// 	a.session.Session.LastRoute = "/"
	// 	a.projects.Session = a.session.Session
	// 	a.Filename = fn
	// 	a.projects.Save(fn)
	// 	messages.EmitMessage(a.ctx, messages.Document, &messages.MessageMsg{
	// 		String1: a.Filename,
	// 		String2: "Saved",
	// 	})
	// } else {
	// }
}

func (a *App) FileSaveAs(cd *menu.CallbackData) {
	// a.saveSession()
}
