package hpack


type HuffmanCode struct {
	code string
	bitLength int
}

var HuffmanTable = map[byte]HuffmanCode{
	byte(0): HuffmanCode{  "3ffffba",  26 },
	byte(1): HuffmanCode{  "3ffffbb",  26 },
	byte(2): HuffmanCode{  "3ffffbc",  26 },
	byte(3): HuffmanCode{  "3ffffbd",  26 },
	byte(4): HuffmanCode{  "3ffffbe",  26 },
	byte(5): HuffmanCode{  "3ffffbf",  26 },
	byte(6): HuffmanCode{  "3ffffc0",  26 },
	byte(7): HuffmanCode{  "3ffffc1",  26 },
	byte(8): HuffmanCode{  "3ffffc2",  26 },
	byte(9): HuffmanCode{  "3ffffc3",  26 },
	byte(10): HuffmanCode{  "3ffffc4",  26 },
	byte(11): HuffmanCode{  "3ffffc5",  26 },
	byte(12): HuffmanCode{  "3ffffc6",  26 },
	byte(13): HuffmanCode{  "3ffffc7",  26 },
	byte(14): HuffmanCode{  "3ffffc8",  26 },
	byte(15): HuffmanCode{  "3ffffc9",  26 },
	byte(16): HuffmanCode{  "3ffffca",  26 },
	byte(17): HuffmanCode{  "3ffffcb",  26 },
	byte(18): HuffmanCode{  "3ffffcc",  26 },
	byte(19): HuffmanCode{  "3ffffcd",  26 },
	byte(20): HuffmanCode{  "3ffffce",  26 },
	byte(21): HuffmanCode{  "3ffffcf",  26 },
	byte(22): HuffmanCode{  "3ffffd0",  26 },
	byte(23): HuffmanCode{  "3ffffd1",  26 },
	byte(24): HuffmanCode{  "3ffffd2",  26 },
	byte(25): HuffmanCode{  "3ffffd3",  26 },
	byte(26): HuffmanCode{  "3ffffd4",  26 },
	byte(27): HuffmanCode{  "3ffffd5",  26 },
	byte(28): HuffmanCode{  "3ffffd6",  26 },
	byte(29): HuffmanCode{  "3ffffd7",  26 },
	byte(30): HuffmanCode{  "3ffffd8",  26 },
	byte(31): HuffmanCode{  "3ffffd9",  26 },
	byte(32): HuffmanCode{  "6",   5 },
	byte(33): HuffmanCode{  "1ffc",  13 },
	byte(34): HuffmanCode{  "1f0",   9 },
	byte(35): HuffmanCode{  "3ffc",  14 },
	byte(36): HuffmanCode{  "7ffc",  15 },
	byte(37): HuffmanCode{  "1e",   6 },
	byte(38): HuffmanCode{  "64",   7 },
	byte(39): HuffmanCode{  "1ffd",  13 },
	byte(40): HuffmanCode{  "3fa",  10 },
	byte(41): HuffmanCode{  "1f1",   9 },
	byte(42): HuffmanCode{  "3fb",  10 },
	byte(43): HuffmanCode{  "3fc",  10 },
	byte(44): HuffmanCode{  "65",   7 },
	byte(45): HuffmanCode{  "66",   7 },
	byte(46): HuffmanCode{  "1f",   6 },
	byte(47): HuffmanCode{  "7",   5 },
	byte(48): HuffmanCode{  "0",   4 },
	byte(49): HuffmanCode{  "1",   4 },
	byte(50): HuffmanCode{  "2",   4 },
	byte(51): HuffmanCode{  "8",   5 },
	byte(52): HuffmanCode{  "20",   6 },
	byte(53): HuffmanCode{  "21",   6 },
	byte(54): HuffmanCode{  "22",   6 },
	byte(55): HuffmanCode{  "23",   6 },
	byte(56): HuffmanCode{  "24",   6 },
	byte(57): HuffmanCode{  "25",   6 },
	byte(58): HuffmanCode{  "26",   6 },
	byte(59): HuffmanCode{  "ec",   8 },
	byte(60): HuffmanCode{  "1fffc",  17 },
	byte(61): HuffmanCode{  "27",   6 },
	byte(62): HuffmanCode{  "7ffd",  15 },
	byte(63): HuffmanCode{  "3fd",  10 },
	byte(64): HuffmanCode{  "7ffe",  15 },
	byte(65): HuffmanCode{  "67",   7 },
	byte(66): HuffmanCode{  "ed",   8 },
	byte(67): HuffmanCode{  "ee",   8 },
	byte(68): HuffmanCode{  "68",   7 },
	byte(69): HuffmanCode{  "ef",   8 },
	byte(70): HuffmanCode{  "69",   7 },
	byte(71): HuffmanCode{  "6a",   7 },
	byte(72): HuffmanCode{  "1f2",   9 },
	byte(73): HuffmanCode{  "f0",   8 },
	byte(74): HuffmanCode{  "1f3",   9 },
	byte(75): HuffmanCode{  "1f4",   9 },
	byte(76): HuffmanCode{  "1f5",   9 },
	byte(77): HuffmanCode{  "6b",   7 },
	byte(78): HuffmanCode{  "6c",   7 },
	byte(79): HuffmanCode{  "f1",   8 },
	byte(80): HuffmanCode{  "f2",   8 },
	byte(81): HuffmanCode{  "1f6",   9 },
	byte(82): HuffmanCode{  "1f7",   9 },
	byte(83): HuffmanCode{  "6d",   7 },
	byte(84): HuffmanCode{  "28",   6 },
	byte(85): HuffmanCode{  "f3",   8 },
	byte(86): HuffmanCode{  "1f8",   9 },
	byte(87): HuffmanCode{  "1f9",   9 },
	byte(88): HuffmanCode{  "f4",   8 },
	byte(89): HuffmanCode{  "1fa",   9 },
	byte(90): HuffmanCode{  "1fb",   9 },
	byte(91): HuffmanCode{  "7fc",  11 },
	byte(92): HuffmanCode{  "3ffffda",  26 },
	byte(93): HuffmanCode{  "7fd",  11 },
	byte(94): HuffmanCode{  "3ffd",  14 },
	byte(95): HuffmanCode{  "6e",   7 },
	byte(96): HuffmanCode{  "3fffe",  18 },
	byte(97): HuffmanCode{  "9",   5 },
	byte(98): HuffmanCode{  "6f",   7 },
	byte(99): HuffmanCode{  "a",   5 },
	byte(100): HuffmanCode{  "29",   6 },
	byte(101): HuffmanCode{  "b",   5 },
	byte(102): HuffmanCode{  "70",   7 },
	byte(103): HuffmanCode{  "2a",   6 },
	byte(104): HuffmanCode{  "2b",   6 },
	byte(105): HuffmanCode{  "c",   5 },
	byte(106): HuffmanCode{  "f5",   8 },
	byte(107): HuffmanCode{  "f6",   8 },
	byte(108): HuffmanCode{  "2c",   6 },
	byte(109): HuffmanCode{  "2d",   6 },
	byte(110): HuffmanCode{  "2e",   6 },
	byte(111): HuffmanCode{  "d",   5 },
	byte(112): HuffmanCode{ "2f",   6 },
	byte(113): HuffmanCode{  "1fc",   9 },
	byte(114): HuffmanCode{ "30",   6 },
	byte(115): HuffmanCode{  "31",   6 },
	byte(116): HuffmanCode{  "e",   5 },
	byte(117): HuffmanCode{  "71",   7 },
	byte(118): HuffmanCode{  "72",   7 },
	byte(119): HuffmanCode{  "73",   7 },
	byte(120): HuffmanCode{  "74",   7 },
	byte(121): HuffmanCode{  "75",   7 },
	byte(122): HuffmanCode{  "f7",   8 },
	byte(123): HuffmanCode{  "1fffd",  17 },
	byte(124): HuffmanCode{  "ffc",  12 },
	byte(125): HuffmanCode{  "1fffe",  17 },
	byte(126): HuffmanCode{  "ffd",  12 },
	byte(127): HuffmanCode{  "3ffffdb",  26 },
	byte(128): HuffmanCode{  "3ffffdc",  26 },
	byte(129): HuffmanCode{  "3ffffdd",  26 },
	byte(130): HuffmanCode{  "3ffffde",  26 },
	byte(131): HuffmanCode{  "3ffffdf",  26 },
	byte(132): HuffmanCode{  "3ffffe0",  26 },
	byte(133): HuffmanCode{  "3ffffe1",  26 },
	byte(134): HuffmanCode{  "3ffffe2",  26 },
	byte(135): HuffmanCode{  "3ffffe3",  26 },
	byte(136): HuffmanCode{  "3ffffe4",  26 },
	byte(137): HuffmanCode{  "3ffffe5",  26 },
	byte(138): HuffmanCode{  "3ffffe6",  26 },
	byte(139): HuffmanCode{  "3ffffe7",  26 },
	byte(140): HuffmanCode{  "3ffffe8",  26 },
	byte(141): HuffmanCode{  "3ffffe9",  26 },
	byte(142): HuffmanCode{  "3ffffea",  26 },
	byte(143): HuffmanCode{  "3ffffeb",  26 },
	byte(144): HuffmanCode{  "3ffffec",  26 },
	byte(145): HuffmanCode{  "3ffffed",  26 },
	byte(146): HuffmanCode{  "3ffffee",  26 },
	byte(147): HuffmanCode{  "3ffffef",  26 },
	byte(148): HuffmanCode{  "3fffff0",  26 },
	byte(149): HuffmanCode{  "3fffff1",  26 },
	byte(150): HuffmanCode{  "3fffff2",  26 },
	byte(151): HuffmanCode{  "3fffff3",  26 },
	byte(152): HuffmanCode{  "3fffff4",  26 },
	byte(153): HuffmanCode{  "3fffff5",  26 },
	byte(154): HuffmanCode{  "3fffff6",  26 },
	byte(155): HuffmanCode{  "3fffff7",  26 },
	byte(156): HuffmanCode{  "3fffff8",  26 },
	byte(157): HuffmanCode{  "3fffff9",  26 },
	byte(158): HuffmanCode{  "3fffffa",  26 },
	byte(159): HuffmanCode{  "3fffffb",  26 },
	byte(160): HuffmanCode{  "3fffffc",  26 },
	byte(161): HuffmanCode{  "3fffffd",  26 },
	byte(162): HuffmanCode{  "3fffffe",  26 },
	byte(163): HuffmanCode{  "3ffffff",  26 },
	byte(164): HuffmanCode{  "1ffff80",  25 },
	byte(165): HuffmanCode{  "1ffff81",  25 },
	byte(166): HuffmanCode{  "1ffff82",  25 },
	byte(167): HuffmanCode{  "1ffff83",  25 },
	byte(168): HuffmanCode{  "1ffff84",  25 },
	byte(169): HuffmanCode{  "1ffff85",  25 },
	byte(170): HuffmanCode{  "1ffff86",  25 },
	byte(171): HuffmanCode{  "1ffff87",  25 },
	byte(172): HuffmanCode{  "1ffff88",  25 },
	byte(173): HuffmanCode{  "1ffff89",  25 },
	byte(174): HuffmanCode{  "1ffff8a",  25 },
	byte(175): HuffmanCode{  "1ffff8b",  25 },
	byte(176): HuffmanCode{  "1ffff8c",  25 },
	byte(177): HuffmanCode{  "1ffff8d",  25 },
	byte(178): HuffmanCode{  "1ffff8e",  25 },
	byte(179): HuffmanCode{  "1ffff8f",  25 },
	byte(180): HuffmanCode{  "1ffff90",  25 },
	byte(181): HuffmanCode{  "1ffff91",  25 },
	byte(182): HuffmanCode{  "1ffff92",  25 },
	byte(183): HuffmanCode{  "1ffff93",  25 },
	byte(184): HuffmanCode{  "1ffff94",  25 },
	byte(185): HuffmanCode{  "1ffff95",  25 },
	byte(186): HuffmanCode{  "1ffff96",  25 },
	byte(187): HuffmanCode{  "1ffff97",  25 },
	byte(188): HuffmanCode{  "1ffff98",  25 },
	byte(189): HuffmanCode{  "1ffff99",  25 },
	byte(190): HuffmanCode{  "1ffff9a",  25 },
	byte(191): HuffmanCode{  "1ffff9b",  25 },
	byte(192): HuffmanCode{  "1ffff9c",  25 },
	byte(193): HuffmanCode{  "1ffff9d",  25 },
	byte(194): HuffmanCode{  "1ffff9e",  25 },
	byte(195): HuffmanCode{  "1ffff9f",  25 },
	byte(196): HuffmanCode{  "1ffffa0",  25 },
	byte(197): HuffmanCode{  "1ffffa1",  25 },
	byte(198): HuffmanCode{  "1ffffa2",  25 },
	byte(199): HuffmanCode{  "1ffffa3",  25 },
	byte(200): HuffmanCode{  "1ffffa4",  25 },
	byte(201): HuffmanCode{  "1ffffa5",  25 },
	byte(202): HuffmanCode{  "1ffffa6",  25 },
	byte(203): HuffmanCode{  "1ffffa7",  25 },
	byte(204): HuffmanCode{  "1ffffa8",  25 },
	byte(205): HuffmanCode{  "1ffffa9",  25 },
	byte(206): HuffmanCode{  "1ffffaa",  25 },
	byte(207): HuffmanCode{  "1ffffab",  25 },
	byte(208): HuffmanCode{  "1ffffac",  25 },
	byte(209): HuffmanCode{  "1ffffad",  25 },
	byte(210): HuffmanCode{  "1ffffae",  25 },
	byte(211): HuffmanCode{  "1ffffaf",  25 },
	byte(212): HuffmanCode{  "1ffffb0",  25 },
	byte(213): HuffmanCode{  "1ffffb1",  25 },
	byte(214): HuffmanCode{  "1ffffb2",  25 },
	byte(215): HuffmanCode{  "1ffffb3",  25 },
	byte(216): HuffmanCode{  "1ffffb4",  25 },
	byte(217): HuffmanCode{  "1ffffb5",  25 },
	byte(218): HuffmanCode{  "1ffffb6",  25 },
	byte(219): HuffmanCode{  "1ffffb7",  25 },
	byte(220): HuffmanCode{  "1ffffb8",  25 },
	byte(221): HuffmanCode{  "1ffffb9",  25 },
	byte(222): HuffmanCode{  "1ffffba",  25 },
	byte(223): HuffmanCode{  "1ffffbb",  25 },
	byte(224): HuffmanCode{  "1ffffbc",  25 },
	byte(225): HuffmanCode{  "1ffffbd",  25 },
	byte(226): HuffmanCode{  "1ffffbe",  25 },
	byte(227): HuffmanCode{  "1ffffbf",  25 },
	byte(228): HuffmanCode{  "1ffffc0",  25 },
	byte(229): HuffmanCode{  "1ffffc1",  25 },
	byte(230): HuffmanCode{  "1ffffc2",  25 },
	byte(231): HuffmanCode{  "1ffffc3",  25 },
	byte(232): HuffmanCode{  "1ffffc4",  25 },
	byte(233): HuffmanCode{  "1ffffc5",  25 },
	byte(234): HuffmanCode{  "1ffffc6",  25 },
	byte(235): HuffmanCode{  "1ffffc7",  25 },
	byte(236): HuffmanCode{  "1ffffc8",  25 },
	byte(237): HuffmanCode{  "1ffffc9",  25 },
	byte(238): HuffmanCode{  "1ffffca",  25 },
	byte(239): HuffmanCode{  "1ffffcb",  25 },
	byte(240): HuffmanCode{  "1ffffcc",  25 },
	byte(241): HuffmanCode{  "1ffffcd",  25 },
	byte(242): HuffmanCode{  "1ffffce",  25 },
	byte(243): HuffmanCode{  "1ffffcf",  25 },
	byte(244): HuffmanCode{  "1ffffd0",  25 },
	byte(245): HuffmanCode{  "1ffffd1",  25 },
	byte(246): HuffmanCode{  "1ffffd2",  25 },
	byte(247): HuffmanCode{  "1ffffd3",  25 },
	byte(248): HuffmanCode{  "1ffffd4",  25 },
	byte(249): HuffmanCode{  "1ffffd5",  25 },
	byte(250): HuffmanCode{  "1ffffd6",  25 },
	byte(251): HuffmanCode{  "1ffffd7",  25 },
	byte(252): HuffmanCode{  "1ffffd8",  25 },
	byte(253): HuffmanCode{  "1ffffd9",  25 },
	byte(254): HuffmanCode{  "1ffffda",  25 },
	byte(255): HuffmanCode{  "1ffffdb",  25 },
}

var HuffmanEOS = HuffmanCode{ "1ffffdc", 25 }
