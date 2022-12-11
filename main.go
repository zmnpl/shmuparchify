package main

import "github.com/zmnpl/shmuparchify/bubl"

var (
	retroarchCfgDir = "/home/simon/.config/retroarch"
)

func main() {
	// TODO get params / ask user for dir

	// checkInstall()
	// // global retroarch settings
	// core.Update_cfg(filepath.Join(retroarchCfgDir, core.RETROARCH_CFG), core.GlobalSettings)
	// // fb neo core settings
	// // core.Update_cfg(filepath.Join(retroarchCfgDir, core.FBNEO_CFG_DIR, core.FBNEO_CFG), core.FBNeoCoreSettings)
	// // fb neo game settings
	// for g, settings := range core.GameSettings {
	// 	core.Update_cfg(filepath.Join(retroarchCfgDir, core.FBNEO_CFG_DIR, g)+".cfg", settings)
	// }

	bubl.Run()

	//update_cfg("/home/simon/.config/retroarch/config/FinalBurn Neo/ddonpach.cfg", gameSettings["ddonpach"])

}

func checkInstall() {
	// TODO
	// supposed to do some checks here before screwing with configs
	// is config folder empty? -> yolo
	// all readable?
	// all writeable?
	// ..
}
