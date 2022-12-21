package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zmnpl/shmuparchify/core"
)

// blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

func Run() {
	//os.Setenv("FYNE_THEME", "light")

	a := app.New()
	a.Settings().SetTheme(newCustomTheme())
	w := a.NewWindow("ShmupArchify - Make your RetroArch Shmup Ready")
	w.Resize(fyne.NewSize(800, 600))

	hello := widget.NewLabel("Enter your RetroArch config dir below:")

	// input field for path
	pathEntry := widget.NewEntry()
	pathEntry.SetText(core.TryFindRetroarchCFGDir())
	pathEntry.Validator = func(string) error {
		sai := core.NewSAIFI(pathEntry.Text)
		if sai.CheckRetroarchCfgExists() {
			return nil
		}
		return fmt.Errorf("Yeah")
	}

	// create report view
	reportMD := "## Report\n\n"
	reportRichText := widget.NewRichTextFromMarkdown(reportMD)
	reportRichText.Wrapping = fyne.TextTruncate
	reportScroll := container.NewVScroll(reportRichText)
	reportScroll.SetMinSize(fyne.NewSize(550, 350))

	// options
	coreOptionsCheck := widget.NewCheck("ShmupArch Core Settings", func(value bool) {})
	coreOptionsCheck.Checked = true
	coreOptionsCheck.Disable()
	bezelsCheck := widget.NewCheck("Download and Setup Arcade Bezels", func(value bool) {})
	cosmeticsCheck := widget.NewCheck("RetroArch Cosmetics", func(value bool) {})

	// button to run everything
	okButton := widget.NewButton("Ok", func() {
		sai := core.NewSAIFI(pathEntry.Text)
		report, _ := sai.SetShmupArchCoreSettings()

		// patch together "markdown" for richtext view
		for _, m := range report {
			status := "FAILED"
			if m.Success {
				status = "SUCCESS"
			}
			reportMD += fmt.Sprintf("***%s*** | %s\n\n", status, m.Text)
		}

		reportRichText.ParseMarkdown(reportMD)
	})

	// build layout
	mainLayout := container.NewVBox(
		hello,
		pathEntry,
		widget.NewSeparator(),
		coreOptionsCheck,
		bezelsCheck,
		cosmeticsCheck,
		okButton,
		widget.NewSeparator(),
		//widget.NewLabel("Report"),
		reportScroll)

	tabs := container.NewAppTabs(
		container.NewTabItem("Run", mainLayout),
		container.NewTabItem("Help", widget.NewLabel("TODO - Help here")),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)

	w.ShowAndRun()
}
