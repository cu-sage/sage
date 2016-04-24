package models

var typeMappings map[string][]string

func init() {
	typeMappings = make(map[string][]string)

	typeMappings["sensing"] = []string{
		"touching:",
		"touchingColor:",
		"color:sees:",
		"distanceTo:",
		"doAsk",
		"answer",
		"keyPressed:",
		"mousePressed",
		"mouseX",
		"mouseY",
		"soundLevel",
		"senseVideoMotion",
		"setVideoState",
		"setVideoTransparency",
		"timer",
		"timerReset",
		"getAttribute:of:",
		"timeAndDate",
		"timestamp",
		"getUserName",
	}
}

// GetBlocksForType returns block names that are part of a
// given block type category
func GetBlocksForType(blockType string) ([]string, bool) {
	blocks, prs := typeMappings[blockType]
	return blocks, prs
}
