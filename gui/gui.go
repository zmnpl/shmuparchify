package gui

import (
	"fmt"
	"net/url"
	"strings"

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
	testPathExists := widget.NewCheck("Exists", nil)
	testPathExists.Disable()
	testRACfgExists := widget.NewCheck("retroarch.cfg exists", nil)
	testRACfgExists.Disable()
	testCanRW := widget.NewCheck("Can read/write", nil)
	testCanRW.Disable()

	// input field for path
	pathEntry := widget.NewEntry()
	pathEntry.SetText(core.TryFindRetroarchCFGDir())
	pathEntry.Validator = func(string) error {
		sai := core.NewRATransformer(pathEntry.Text)

		testPathExists.Enable()
		testRACfgExists.Enable()
		testCanRW.Enable()
		testPathExists.Checked = sai.CheckPathExists()
		testRACfgExists.Checked = sai.CheckRetroarchCfgExists()
		testCanRW.Checked = sai.CheckPermissions()
		testPathExists.Disable()
		testRACfgExists.Disable()
		testCanRW.Disable()

		if testPathExists.Checked && testRACfgExists.Checked && testCanRW.Checked {
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

	//button to run everything
	okButton := widget.NewButton("Apply Settings", func() {
		opts := make([]func(core.RetroArchChanger) core.RetroArchChanger, 0)

		r := core.NewRATransformer(pathEntry.Text, opts...)
		// execute jobs
		for _, j := range r.GetShmupArchJobs() {
			reportMD += messageToMD(j())
		}

		reportRichText.ParseMarkdown(reportMD)
		reportScroll.ScrollToBottom()
	})

	// overlays
	customRomsOverlays := widget.NewEntry() // TODO: Make this non-anonymous and use it
	// button to download overlays
	overlayDLProgress := widget.NewProgressBar()
	dlOverlaysButton := widget.NewButton("Download Overlays", func() {
		opts := make([]func(core.RetroArchChanger) core.RetroArchChanger, 0)
		r := core.NewRATransformer(pathEntry.Text, opts...)

		overlayJobs := r.GetOverlayJobs()
		// add custom jobs from input
		customOverlayRoms := strings.Split(customRomsOverlays.Text, ",")
		for _, game := range customOverlayRoms {
			overlayJobs = append(overlayJobs, r.MakeOverlayJob(game))
		}

		overlayDLProgress.Min = 0
		overlayDLProgress.Max = float64(len(overlayJobs))

		progress := 0.0
		go func() {
			for _, j := range overlayJobs {
				reportMD += messageToMD(j())

				// update progress bar
				progress += 1
				overlayDLProgress.SetValue(progress)

				// update report view
				reportRichText.ParseMarkdown(reportMD)
				reportScroll.ScrollToBottom()
			}
		}()
	})

	// tab: ShmupArch Options
	shmupArchOptionsLayout := container.NewVBox(
		widget.NewLabel("Options"),
		container.NewHBox(coreOptionsCheck, cosmeticsCheck),
		widget.NewSeparator(),
		container.NewHBox(layout.NewSpacer(), okButton),
	)

	// tab: Overlays Download
	downloadOverlaysLayout := container.NewVBox(
		widget.NewLabel("Hit the button to download overlays for alle the Shmups known to ShmupArch."),
		widget.NewLabel("Optional: Enter a list of rom names ,-separated to download overlays for them"),
		customRomsOverlays,
		widget.NewSeparator(),
		container.NewHBox(layout.NewSpacer(), dlOverlaysButton),
		overlayDLProgress,
	)

	// tab: help
	githubUrl, _ := url.Parse("https://github.com/zmnpl/shmuparchify")
	whatAndWhyUrl, _ := url.Parse("https://www.patreon.com/posts/article-what-is-57566721?l=fr")
	videoUrl, _ := url.Parse("https://www.youtube.com/watch?v=Sec3r6RKAPg")
	helpLayout := container.NewVBox(
		widget.NewLabel("Find more information here:"),
		widget.NewHyperlink("github.com/zmnpl/shmuparchify", githubUrl),
		widget.NewLabel("About the origins of ShmupArch:"),
		widget.NewHyperlink("What is ShmupArch? Why Does it Matter?", whatAndWhyUrl),
		widget.NewHyperlink("Video on YouTube", videoUrl),
	)

	// top section with ra config path entry
	openFolderButton := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {})
	retroArchPathLayout := container.NewVBox(
		widget.NewToolbar(
			widget.NewToolbarAction(theme.FileIcon(), nil),
			widget.NewToolbarAction(theme.CancelIcon(), func() {}),
		),
		hello,
		container.New(layout.NewBorderLayout(nil, nil, openFolderButton, nil), openFolderButton, pathEntry),
		container.NewHBox(testPathExists, testRACfgExists, testCanRW),
		widget.NewSeparator(),
	)

	// tabs holding the app functions
	tabs := container.NewAppTabs(
		container.NewTabItem("ShmupArch Config", shmupArchOptionsLayout),
		container.NewTabItem("Arcade Overlays", downloadOverlaysLayout),
		container.NewTabItem("Help", helpLayout),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	//mainLayout := fyne.NewContainerWithLayout(layout.NewBorderLayout(foo, nil, nil, nil), foo, reportScroll)
	mainLayout := container.New(layout.NewBorderLayout(retroArchPathLayout, reportScroll, nil, nil), retroArchPathLayout, reportScroll, tabs)
	w.SetContent(mainLayout)

	w.ShowAndRun()
}

func messageToMD(m core.Message) string {
	result := "FAILED"
	if m.Success {
		result = "SUCCESS"
	}

	return fmt.Sprintf("***%s*** | %s\n\n", result, m.Text)
}
