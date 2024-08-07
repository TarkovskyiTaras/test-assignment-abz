package guimanager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"test-assignment-abz/config"
	"test-assignment-abz/datamanager"

	"github.com/sirupsen/logrus"
)

type GUIManager struct {
	app      fyne.App
	window   fyne.Window
	trayIcon *fyne.StaticResource
	fs       *datamanager.FileStorage
	cfg      *config.Config
	logger   *logrus.Logger
}

func NewGUIManager(fs *datamanager.FileStorage, cfg *config.Config) *GUIManager {
	a := app.New()
	w := a.NewWindow("SysTray")
	svgContent := `<svg width="800px" height="800px" viewBox="0 0 1024 1024" class="icon"  version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M384 384m-320 0a320 320 0 1 0 640 0 320 320 0 1 0-640 0Z" fill="#3F51B5" /><path d="M433.066667 341.333333v36.266667h-81.066667v29.866667h81.066667v36.266666h-81.066667c0 12.8 2.133333 25.6 6.4 34.133334 4.266667 8.533333 8.533333 17.066667 14.933333 21.333333 6.4 6.4 14.933333 8.533333 23.466667 12.8 8.533333 2.133333 19.2 4.266667 29.866667 4.266667 8.533333 0 14.933333 0 23.466666-2.133334 8.533333-2.133333 14.933333-2.133333 21.333334-6.4l8.533333 57.6c-8.533333 2.133333-19.2 4.266667-29.866667 4.266667-10.666667 2.133333-21.333333 2.133333-32 2.133333-19.2 0-38.4-2.133333-55.466666-8.533333-17.066667-4.266667-32-12.8-42.666667-23.466667-12.8-10.666667-21.333333-23.466667-29.866667-40.533333-6.4-14.933333-10.666667-34.133333-10.666666-55.466667h-40.533334v-36.266666h40.533334v-29.866667h-40.533334V341.333333h40.533334c2.133333-21.333333 6.4-38.4 12.8-55.466666 8.533333-14.933333 17.066667-29.866667 29.866666-40.533334 12.8-10.666667 27.733333-19.2 44.8-23.466666 17.066667-6.4 36.266667-8.533333 55.466667-8.533334 8.533333 0 19.2 0 27.733333 2.133334s19.2 2.133333 27.733334 6.4l-8.533334 57.6c-6.4-2.133333-12.8-4.266667-21.333333-6.4-8.533333-2.133333-14.933333-2.133333-23.466667-2.133334-10.666667 0-21.333333 2.133333-29.866666 4.266667-8.533333 2.133333-17.066667 6.4-21.333334 12.8-6.4 6.4-10.666667 12.8-14.933333 21.333333s-6.4 19.2-6.4 32h81.066667z" fill="#FFF59D" /><path d="M640 640m-320 0a320 320 0 1 0 640 0 320 320 0 1 0-640 0Z" fill="#4CAF50" /><path d="M605.866667 576c2.133333 4.266667 4.266667 8.533333 8.533333 12.8 4.266667 4.266667 8.533333 8.533333 14.933333 10.666667 6.4 4.266667 14.933333 6.4 23.466667 10.666666 14.933333 6.4 29.866667 12.8 42.666667 19.2 12.8 6.4 23.466667 14.933333 32 23.466667 8.533333 8.533333 17.066667 19.2 21.333333 29.866667 4.266667 10.666667 8.533333 25.6 8.533333 40.533333s-2.133333 27.733333-6.4 38.4c-4.266667 10.666667-10.666667 21.333333-19.2 29.866667s-19.2 14.933333-29.866666 19.2c-12.8 4.266667-25.6 8.533333-38.4 10.666666v46.933334h-38.4v-46.933334c-12.8-2.133333-25.6-4.266667-38.4-8.533333s-23.466667-10.666667-32-21.333333c-10.666667-8.533333-17.066667-21.333333-23.466667-34.133334-6.4-12.8-8.533333-29.866667-8.533333-49.066666h70.4c0 10.666667 2.133333 21.333333 4.266666 27.733333 2.133333 8.533333 6.4 12.8 12.8 19.2 4.266667 4.266667 10.666667 8.533333 17.066667 10.666667 6.4 2.133333 12.8 2.133333 19.2 2.133333 8.533333 0 14.933333 0 19.2-2.133333 6.4-2.133333 10.666667-4.266667 14.933333-8.533334 4.266667-4.266667 6.4-8.533333 8.533334-12.8 2.133333-4.266667 2.133333-10.666667 2.133333-17.066666 0-6.4 0-12.8-2.133333-17.066667-2.133333-4.266667-4.266667-10.666667-8.533334-14.933333s-8.533333-8.533333-14.933333-10.666667c-6.4-4.266667-14.933333-6.4-23.466667-10.666667-14.933333-6.4-29.866667-12.8-42.666666-19.2-12.8-6.4-23.466667-14.933333-32-23.466666-8.533333-8.533333-17.066667-19.2-21.333334-29.866667-4.266667-10.666667-8.533333-25.6-8.533333-40.533333 0-12.8 2.133333-25.6 6.4-36.266667 4.266667-10.666667 10.666667-21.333333 19.2-29.866667 8.533333-8.533333 19.2-14.933333 29.866667-21.333333 10.666667-4.266667 25.6-8.533333 38.4-10.666667v-51.2h38.4v51.2c12.8 2.133333 25.6 6.4 38.4 12.8 10.666667 6.4 21.333333 12.8 27.733333 23.466667 8.533333 8.533333 14.933333 21.333333 19.2 34.133333 4.266667 12.8 6.4 27.733333 6.4 42.666667h-70.4c0-19.2-4.266667-34.133333-12.8-42.666667-8.533333-8.533333-19.2-14.933333-32-14.933333-6.4 0-12.8 2.133333-19.2 4.266667-4.266667 2.133333-8.533333 4.266667-12.8 8.533333-4.266667 4.266667-6.4 8.533333-6.4 12.8-2.133333 4.266667-2.133333 10.666667-2.133333 17.066667-2.133333 4.266667 0 10.666667 0 14.933333z" fill="#FFFFFF" /></svg>`
	iconResource := fyne.NewStaticResource("icon.svg", []byte(svgContent))

	return &GUIManager{
		app:      a,
		window:   w,
		trayIcon: iconResource,
		fs:       fs,
		cfg:      cfg,
		logger:   logrus.New(),
	}
}

func (gui *GUIManager) SetupTray() {
	if desk, ok := gui.app.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				gui.window.Show()
			}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(gui.trayIcon)
	}
}

func (gui *GUIManager) SetupWindows() {
	var checkBoxList []fyne.CanvasObject
	selected := make(map[string]bool)
	for _, currency := range gui.cfg.Currencies {
		check := widget.NewCheck(currency, func(checked bool) {
			if checked {
				selected[currency] = true
			} else {
				delete(selected, currency)
			}
		})
		checkBoxList = append(checkBoxList, check)

		getTodaysRatesButton := widget.NewButton("Get Today's Rates", func() {
			gui.showTodaysRates(selected)
		})

		getThisMonthRatesButton := widget.NewButton("Get Rates for This Month", func() {
			gui.showThisMonthRates(selected)
		})

		getLastMonthRatesButton := widget.NewButton("Get Rates for Last Month", func() {
			gui.showLastMonthRates(selected)
		})

		buttons := container.NewVBox(
			getTodaysRatesButton,
			getThisMonthRatesButton,
			getLastMonthRatesButton,
		)

		checkBoxContainer := container.NewVBox(checkBoxList...)
		allContainers := container.NewVBox(checkBoxContainer, buttons)

		gui.window.SetContent(allContainers)

		gui.window.SetCloseIntercept(func() {
			gui.window.Hide()
		})

		gui.window.SetIcon(gui.trayIcon)
	}
}

func (gui *GUIManager) showThisMonthRates(selected map[string]bool) {
	thisMonth, err := gui.fs.GetThisMonthData(selected)
	if err != nil {
		gui.logger.Info(err)
	}

	win := gui.app.NewWindow("Rates This Month")
	jsonLabel := widget.NewLabel(string(thisMonth))
	jsonLabel.Wrapping = fyne.TextWrapBreak

	win.SetContent(container.NewScroll(jsonLabel))
	win.Resize(fyne.NewSize(400, 300))
	win.Show()
}

func (gui *GUIManager) showTodaysRates(selected map[string]bool) {
	todaysRates, err := gui.fs.GetTodaysData(selected)
	if err != nil {
		gui.logger.Info(err)
	}

	win := gui.app.NewWindow("Rates This Month")
	jsonLabel := widget.NewLabel(string(todaysRates))
	jsonLabel.Wrapping = fyne.TextWrapBreak

	win.SetContent(container.NewScroll(jsonLabel))
	win.Resize(fyne.NewSize(400, 300))
	win.Show()
}

func (gui *GUIManager) showLastMonthRates(selected map[string]bool) {
	thisMonth, err := gui.fs.GetLastMonthData(selected)
	if err != nil {
		gui.logger.Info(err)
	}

	win := gui.app.NewWindow("Rates This Month")
	jsonLabel := widget.NewLabel(string(thisMonth))
	jsonLabel.Wrapping = fyne.TextWrapBreak

	win.SetContent(container.NewScroll(jsonLabel))
	win.Resize(fyne.NewSize(400, 300))
	win.Show()
}

func (gui *GUIManager) Run() {
	gui.SetupTray()
	gui.SetupWindows()
	gui.window.ShowAndRun()
}
