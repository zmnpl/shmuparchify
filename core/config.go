package core

import (
	"fmt"
	"path/filepath"
)

type cfgEntry struct {
	option string
	value  string
}

func (e cfgEntry) String() string {
	return fmt.Sprintf("%v = \"%v\"", e.option, e.value)
}

var (
	GlobalSettings         []cfgEntry
	FBNeoCoreSettings      []cfgEntry
	ShmupArchGameSettings  map[string][]cfgEntry
	AdditionalGameSettings map[string][]cfgEntry
	OverlayFBNeoGames      map[string][]cfgEntry
)

const (
	RETROARCH_CFG         = "retroarch.cfg"
	RETROARCH_DEFAULT_CFG = "retroarch.default.cfg"
	FBNEO_CFG             = "fbneo.cfg"
	FBNEO_CFG_DIR         = "config/FinalBurn Neo/"

	OVERLAY_PATH = "/overlay/arcade-overlays/overlays/borders-Various_Creators/"
)

func init() {
	initGlobalSettings()
	initFBNeoCoreSettings()
	initFBNeoGameSettings()
	initFBNeoOverlays()
}

func initGlobalSettings() {
	GlobalSettings = []cfgEntry{
		{option: "video_fullscreen", value: "true"},             // Start in Fullscreen Mode
		{option: "video_windowed_fullscreen", value: "false"},   // Windowed Fullscreen Mode
		{option: "video_hard_sync", value: "true"},              // Hard GPU Sync
		{option: "video_hard_sync_frames", value: "0"},          // Hard GPU Sync Frames
		{option: "video_frame_delay", value: "0"},               // Frame Delay (ms)
		{option: "input_poll_type_behavior", value: "0"},        // Polling Behaviour (0 = Early)
		{option: "run_ahead_enabled", value: "true"},            // Run-Ahead to Reduce Latency
		{option: "run_ahead_frames", value: "2"},                // Number of Frames to Run-Ahead
		{option: "run_ahead_secondary_instance", value: "true"}, // Use Second Instance for Run-Ahead
		{option: "fastforward_ratio", value: "1.000000"},        // Fast Forward Rate
	}
}

func initFBNeoCoreSettings() {
	FBNeoCoreSettings = []cfgEntry{
		//{option: "foo", value: "bar"},
	}
}

// game specific settings for fb neo roms
func initFBNeoGameSettings() {
	ShmupArchGameSettings = make(map[string][]cfgEntry)

	ShmupArchGameSettings["1941j"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["batriderj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["bbakraidj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["bgaregga"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
		{option: "state_slot", value: "0"},
	}

	ShmupArchGameSettings["blazstar"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["esprade"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["gigawingj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["gunbird"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["guwange"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["metalb"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["mmatrixj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["mslug"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["mslug2"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["mslug3"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["mslug4"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["mslugx"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["p47j"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["progearj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["rayforce"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["rtype"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["rtype2"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["rtypeleo"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["s1945"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}
}

// game specific settings for fb neo roms
func initFBNeoGameSettingsAdditional() {
	AdditionalGameSettings = make(map[string][]cfgEntry)

	ShmupArchGameSettings["1941"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["batrider"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["batsugun"] = []cfgEntry{}

	ShmupArchGameSettings["bbakraid"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["darius"] = []cfgEntry{}

	ShmupArchGameSettings["dariusg"] = []cfgEntry{}

	ShmupArchGameSettings["ddonpach"] = []cfgEntry{}

	ShmupArchGameSettings["ddonpacha"] = []cfgEntry{}

	ShmupArchGameSettings["ddp2"] = []cfgEntry{}

	ShmupArchGameSettings["ddp3"] = []cfgEntry{}

	ShmupArchGameSettings["ddpdfk"] = []cfgEntry{}

	ShmupArchGameSettings["ddpdfk10"] = []cfgEntry{}

	ShmupArchGameSettings["ddpdoj"] = []cfgEntry{}

	ShmupArchGameSettings["ddpdojblk"] = []cfgEntry{}

	ShmupArchGameSettings["deathsm2"] = []cfgEntry{}

	ShmupArchGameSettings["deathsml"] = []cfgEntry{}

	ShmupArchGameSettings["donpachi"] = []cfgEntry{}

	ShmupArchGameSettings["donpachij"] = []cfgEntry{}

	ShmupArchGameSettings["dragnblz"] = []cfgEntry{}

	ShmupArchGameSettings["espgal"] = []cfgEntry{}

	ShmupArchGameSettings["espgal2"] = []cfgEntry{}

	ShmupArchGameSettings["futari10"] = []cfgEntry{}

	ShmupArchGameSettings["futari15"] = []cfgEntry{}

	ShmupArchGameSettings["futaribl"] = []cfgEntry{}

	ShmupArchGameSettings["galaga"] = []cfgEntry{}

	ShmupArchGameSettings["galaga3"] = []cfgEntry{}

	ShmupArchGameSettings["gigawing"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["gunbird2"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["ikaruga"] = []cfgEntry{}

	ShmupArchGameSettings["ket"] = []cfgEntry{}

	ShmupArchGameSettings["ketarr"] = []cfgEntry{}

	ShmupArchGameSettings["mmatrix"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["mslug5"] = []cfgEntry{}

	ShmupArchGameSettings["mslug6"] = []cfgEntry{}

	ShmupArchGameSettings["mushisam"] = []cfgEntry{}

	ShmupArchGameSettings["p47"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	ShmupArchGameSettings["progear"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	ShmupArchGameSettings["raiden"] = []cfgEntry{}

	ShmupArchGameSettings["raiden2"] = []cfgEntry{}

	ShmupArchGameSettings["s1945"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	ShmupArchGameSettings["s1945ii"] = []cfgEntry{}

}

func initFBNeoOverlays() {
	// overlays
	OverlayFBNeoGames = make(map[string][]cfgEntry)

	OverlayFBNeoGames["ddonpach"] = []cfgEntry{
		{option: "input_overlay", value: "~/.config/retroarch/overlay/arcade-overlays/overlays/borders-Various_Creators/ddonpach.cfg"},
		{option: "menu_show_advanced_settings", value: "true"},
	}
}

func makeOverlayCfg(raCfgDir string, gameName string) []cfgEntry {
	return []cfgEntry{
		{option: "input_overlay", value: filepath.Join(raCfgDir, OVERLAY_PATH, gameName+".cfg")},
		//{option: "menu_show_advanced_settings", value: "true"},
	}
}
