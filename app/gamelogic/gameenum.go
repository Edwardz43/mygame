package gamelogic

// Cards ...
type Cards map[uint8]string

// Poker contains 52 poker cards with map[int]string form.
var Poker = Cards{
	// Spad
	0:  "s_1",
	1:  "s_2",
	2:  "s_3",
	3:  "s_4",
	4:  "s_5",
	5:  "s_6",
	6:  "s_7",
	7:  "s_8",
	8:  "s_9",
	9:  "s_10",
	10: "s_j",
	11: "s_q",
	12: "s_k",

	// Heart
	13: "h_1",
	14: "h_2",
	15: "h_3",
	16: "h_4",
	17: "h_5",
	18: "h_6",
	19: "h_7",
	20: "h_8",
	21: "h_9",
	22: "h_10",
	23: "h_j",
	24: "h_q",
	25: "h_k",

	// Diamond
	26: "d_1",
	27: "d_2",
	28: "d_3",
	29: "d_4",
	30: "d_5",
	31: "d_6",
	32: "d_7",
	33: "d_8",
	34: "d_9",
	35: "d_10",
	36: "d_j",
	37: "d_q",
	38: "d_k",

	// Club
	39: "c_1",
	40: "c_2",
	41: "c_3",
	42: "c_4",
	43: "c_5",
	44: "c_6",
	45: "c_7",
	46: "c_8",
	47: "c_9",
	48: "c_10",
	49: "c_j",
	50: "c_q",
	51: "c_k",
}
