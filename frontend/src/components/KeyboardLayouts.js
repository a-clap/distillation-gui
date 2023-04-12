var Layouts = {

	"normal": {

		_meta: {
			"tab": { key: "\t", text: "Tab", size: 2, classes: "control" },
			"shiftl": { keySet: "shifted", text: "Shift", size: 2, classes: "control" },
			"shiftr": { keySet: "shifted", text: "Shift", size: 2, classes: "control" },
			"caps": { keySet: "capsed", text: "Caps lock", size: 2, classes: "control" },
			"space": { key: " ", text: "Space", size: 5 },
			"enter": { key: "\r\n", text: "Enter", size: 2, classes: "control" },
			"backspace": { func: "backspace", classes: "control", size: 2 },
			"close": { func: "close", text: "Close", size: 2, classes: "control featured" },
			"next": { func: "next", text: "Next", size: 2, classes: "control featured" }
		},

		default: [
			"` 1 2 3 4 5 6 7 8 9 0 - = {backspace}",
			"{tab} q w e r t y u i o p [ ] \\",
			"{caps} a s d f g h j k l ; ' {enter}",
			"{shiftl} z x c v b n m , . / {shiftr}",
			"{next} {space} {close}"
		],
		shifted: [
			"~ ! @ # $ % ^ & * ( ) _ + {backspace}",
			"{tab} Q W E R T Y U I O P { } |",
			"{caps} A S D F G H J K L : \" {enter}",
			"{shiftl} Z X C V B N M < > ? {shiftr}",
			"{next} {space} {close}"
		],

		capsed: [
			"` 1 2 3 4 5 6 7 8 9 0 - = {backspace}",
			"{tab} Q W E R T Y U I O P [ ] \\",
			"{caps} A S D F G H J K L ; ' {enter}",
			"{shiftl} Z X C V B N M , . / {shiftr}",
			"{next} {space} {close}"
		]
	},

	"compact": {

		_meta: {
			"default": { keySet: "default", text: "abc", classes: "control" },
			"alpha": { keySet: "default", text: "Abc", classes: "control" },
			"shift": { keySet: "shifted", text: "ABC", classes: "control" },
			"numbers": { keySet: "numbers", text: "123", classes: "control" },
			"space": { key: " ", text: "Space", size: 200 },
			"backspace": { func: "backspace", classes: "control" },
			"close": { func: "close", text: "Close", classes: "control featured" },
			"next": { func: "next", text: "Next", classes: "featured" },
			"zero": { key: "0", size: 130 }
		},

		default: [
			"q w e r t y u i o p",
			" a s d f g h j k l ",
			"{shift} z x c v b n m {backspace}",
			"{numbers} , {space} . {next} {close}"
		],

		shifted: [
			"Q W E R T Y U I O P",
			" A S D F G H J K L ",
			"{default} Z X C V B N M ",
			"{numbers} _ {space} {backspace} {next} {close}"
		],

		numbers: [
			"1 2 3",
			"4 5 6",
			"7 8 9",
			"  {alpha} . {zero} {backspace} {next} {close}"
		]
	},

	"numeric": {

		_meta: {
			"backspace": { func: "backspace", classes: "control" },
			"close": { func: "close", text: "Close", classes: "control featured" },
			"next": { func: "next", text: "Next", classes: "control featured" },
			"zero": { key: "0", size: 130 }
		},

		default: [
			"1 2 3",
			"4 5 6",
			"7 8 9",
			"_ - . {zero} {backspace} {next} {close}"
		]
	}
};

export var Layouts