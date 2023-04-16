var Layouts = {

	_meta: {
		"shift": { func: "shift", text: "Shift", size: 2, classes: "control" },
		"enter": { func: "enter", text: "", size: 1, classes: "control icons featured" },
		"backspace": { func: "backspace", classes: "control icons" },
		"clr": { func: "clr", text: "CLR", size: 2, classes: "control " },
		"esc": { func: "esc", text: "", size: 2, classes: "control icons" },
		"space": { key: " ", text: "Space", size: 2 },
	},

	"normal": {
		default: [
			"esc ` 1 2 3 4 5 6 7 8 9 0 - = backspace",
			"q w e r t y u i o p [ ] \\ clr",
			"a s d f g h j k l ; '",
			"shift z x c v b n m , . /",
			"space enter"
		],
		shifted: [
			"esc ~ ! @ # $ % ^ & * ( ) _ + backspace",
			"Q W E R T Y U I O P { } | clr",
			"A S D F G H J K L : \"",
			"shift Z X C V B N M < > ?",
			"space enter"
		],
	},

	"numeric": {
		default: [
			"1 2 3 backspace",
			"4 5 6 clr",
			"7 8 9 esc",
			". 0 - enter"
		]
	}
};

export default Layouts