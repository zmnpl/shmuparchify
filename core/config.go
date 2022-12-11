package core

import "fmt"

type cfgEntry struct {
	option string
	value  string
}

func (e cfgEntry) String() string {
	return fmt.Sprintf("\"%v\" = \"%v\"", e.option, e.value)
}

var (
	GlobalSettings    []cfgEntry
	FBNeoCoreSettings []cfgEntry
	GameSettings      map[string][]cfgEntry
	OverlayFBNeoGames map[string][]cfgEntry
)

var (
	RETROARCH_CFG = "retroarch.cfg"
	FBNEO_CFG     = "fbneo.cfg"
	FBNEO_CFG_DIR = "config/FinalBurn Neo/"
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
	GameSettings = make(map[string][]cfgEntry)

	GameSettings["1941j"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["batriderj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["bbakraidj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["bgaregga"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
		{option: "state_slot", value: "0"},
	}

	GameSettings["blazstar"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["esprade"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["gigawingj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["gunbird"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	GameSettings["guwange"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["metalb"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["mmatrixj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	GameSettings["mslug"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}

	GameSettings["mslug2"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["mslug3"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["mslug4"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["mslugx"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["p47j"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["progearj"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["rayforce"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["rtype"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["rtype2"] = []cfgEntry{
		{option: "run_ahead_frames", value: "2"},
	}

	GameSettings["rtypeleo"] = []cfgEntry{
		{option: "run_ahead_frames", value: "3"},
	}

	GameSettings["s1945"] = []cfgEntry{
		{option: "run_ahead_frames", value: "4"},
	}
}

func initFBNeoOverlays() {
	// overlays
	OverlayFBNeoGames = make(map[string][]cfgEntry)

	OverlayFBNeoGames["ddonpach"] = []cfgEntry{
		{option: "input_overlay", value: "~/.config/retroarch/overlay/arcade-overlays/overlays/borders-Various_Creators/ddonpach.cfg"},
		{option: "menu_show_advanced_settings", value: "true"},
	}
}
