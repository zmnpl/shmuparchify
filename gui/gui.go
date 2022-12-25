package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zmnpl/shmuparchify/core"
)

// blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

func Run() {
	//os.Setenv("FYNE_THEME", "light")

	a := app.New()
	a.Settings().SetTheme(newCustomTheme())
	w := a.NewWindow("ShmupArchify - Make your RetroArch Shmup Ready")
	w.Resize(fyne.NewSize(600, 500))

	hello := widget.NewLabel("Enter your RetroArch config dir below:")

	// path checks
	testRACfgExists := widget.NewCheck("retroarch.cfg exists", nil)
	testRACfgExists.Disable()
	testCanRW := widget.NewCheck("Can read/write", nil)
	testCanRW.Disable()
	testRACfgExists.Checked = true
	testCanRW.Checked = true
	// input field for path
	pathEntry := widget.NewEntry()
	pathEntry.SetText(core.TryFindRetroarchCFGDir())
	pathEntry.Validator = func(string) error {
		sai := core.NewRATransformer(pathEntry.Text)

		testRACfgExists.Enable()
		testCanRW.Enable()
		testRACfgExists.Checked = sai.CheckRetroarchCfgExists()
		testCanRW.Checked = sai.CheckCanReadWrite()
		testRACfgExists.Disable()
		testCanRW.Disable()

		if testRACfgExists.Checked && testCanRW.Checked {
			return nil
		}
		return fmt.Errorf("Yeah")
	}

	// create report view
	reportMD := ""
	reportRichText := widget.NewRichTextFromMarkdown(reportMD)
	reportRichText.Wrapping = fyne.TextTruncate
	reportScroll := container.NewVScroll(reportRichText)
	reportScroll.SetMinSize(fyne.NewSize(200, 200))

	// options
	coreOptionsCheck := widget.NewCheck("ShmupArch Core Settings", func(value bool) {})
	coreOptionsCheck.Checked = true
	coreOptionsCheck.Disable()
	cosmeticsCheck := widget.NewCheck("RetroArch Cosmetics", func(value bool) {})

	// progress bar
	dlProgress := widget.NewProgressBar()
	dlProgress.Hidden = true

	//button to run everything
	okButton := widget.NewButton("Ok", func() {
		opts := make([]func(core.RetroArchChanger) core.RetroArchChanger, 0)

		r := core.NewRATransformer(pathEntry.Text, opts...)
		report, _ := r.SetShmupArchCoreSettings()
		reportMD += messagesToMD(report)
		reportRichText.ParseMarkdown(reportMD)

		// download bezels
		if false {
			dlProgress.Hidden = false
			bezelJobs := r.GetBezelDownloadJobs()
			dlProgress.Min = 0
			dlProgress.Max = float64(len(bezelJobs))

			progress := 0.0
			go func() {
				for _, j := range bezelJobs {
					j()
					progress += 1
					dlProgress.SetValue(progress)
				}
			}()
		}

		reportScroll.ScrollToBottom()
	})

	shmupArchOptionsLayout := container.NewVBox(
		widget.NewLabel("Options"),
		container.NewHBox(coreOptionsCheck, cosmeticsCheck),
		widget.NewSeparator(),
		container.NewHBox(layout.NewSpacer(), okButton),
	)

	helpLayout := container.NewVBox(
		widget.NewLabel("foo"),
		widget.NewLabel("foo"),
		widget.NewLabel("foo"),
		widget.NewLabel("foo"),
	)

	openFolderButton := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {})

	retroArchPathLayout := container.NewVBox(
		widget.NewToolbar(
			widget.NewToolbarAction(theme.FileIcon(), nil),
			widget.NewToolbarAction(theme.CancelIcon(), func() {}),
		),
		hello,
		container.New(layout.NewBorderLayout(nil, nil, openFolderButton, nil), openFolderButton, pathEntry),
		container.NewHBox(testRACfgExists, testCanRW),
		widget.NewSeparator(),
	)

	// tabs holding the app functions
	tabs := container.NewAppTabs(
		container.NewTabItem("ShmupArch Config", shmupArchOptionsLayout),
		container.NewTabItem("Arcade Overlays", widget.NewLabel("foo")),
		container.NewTabItem("Help", helpLayout),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	//mainLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(foo, nil, nil, nil), foo, reportScroll)
	mainLayout := container.New(layout.NewBorderLayout(retroArchPathLayout, reportScroll, nil, nil), retroArchPathLayout, reportScroll, tabs)
	w.SetContent(mainLayout)

	w.ShowAndRun()
}

func messagesToMD(messages []core.Message) string {
	result := ""

	for _, m := range messages {
		status := "FAILED"
		if m.Success {
			status = "SUCCESS"
		}
		result += fmt.Sprintf("***%s*** | %s\n\n", status, m.Text)
	}

	return result
}
