package backend

type aliasedAsset struct {
	Code   string `json:"code"`
	Issuer string `json:"issuer"`
	Alias  string `json:"alias"`
}

func getAssets() []aliasedAsset {
	return []aliasedAsset{
		aliasedAsset{
			Code:   "MXN",
			Issuer: "GCKIK5F6J4KMKF6MKB5EM67S5CK557EZQ3IAMZM5FFAYST63S3HWXVPE",
			Alias:  "AnchorMXN",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GDUKMGUGDZQK6YHYA5Z6AY2G4XDSZPSZ3SW5UN3ARVMO6QSRDWP5YLEX",
			Alias:  "AnchorUSD",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GCQTGZQQ5G4PTM2GL7CDIFKUBIPEC52BROAQIAPW53XBRJVN6ZJVTG6V",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "BAT",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "BCH",
			Issuer: "GAEGOS7I6TYZSVHPSN76RSPIVL3SELATXZWLFO4DIAFYJF7MPAHFE7H4",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GAUTUYY2THLF7SGITDFMXJVYH3LHDSMGEAKSBU267M2K7A3W543CKUEF",
			Alias:  "Apay",
		},
		aliasedAsset{Code: "ETH",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "KIN",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "LINK",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "LTC",
			Issuer: "GC5LOR3BK6KIOK7GKAUD5EGHQCMFOGHJTC7I3ELB66PTDFXORC2VM5LP",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "MTL",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "OMG",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "REP",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "SALT",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "ZRX",
			Issuer: "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			Alias:  "Apay",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GB2O5PBQJDAFCNM2U2DIMVAEI7ISOYL4UJDTLN42JYYXAENKBWY6OBKZ",
			Alias:  "Billex",
		},
		aliasedAsset{
			Code:   "BB1",
			Issuer: "GD5J6HLF5666X4AZLTFTXLY46J5SW7EXRKBLEYPJP33S33MXZGV6CWFN",
			Alias:  "Bitbond",
		},
		aliasedAsset{
			Code:   "PHP",
			Issuer: "GBUQWP3BOUZX34TOND2QV7QQ7K7VJTG6VSE7WMLBTMDJLLAW7YKGU6EP",
			Alias:  "CoinsPro",
		},
		aliasedAsset{
			Code:   "NGNT",
			Issuer: "GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD",
			Alias:  "Cowrie",
		},
		aliasedAsset{
			Code:   "ZAR",
			Issuer: "GBJHCJH3RKPNNNRJKHQWPAYPJAKBB55CJWCDIM4CQGI2ADQF67RII6Z3",
			Alias:  "Demars",
		},
		aliasedAsset{
			Code:   "AAPL",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "D5BK",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "DBXN",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "FB",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "FTSE",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "GOOG",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "MSCIW",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "SP500",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "SXRL",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "TSLA",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "ZM",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "XCS6",
			Issuer: "GBRDHSZL4ZKOI2PTUMM53N3NICZXC5OX3KPCD4WD4NG4XGCBC2ZA3KAG",
			Alias:  "DSQ AG",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GBETHKBL5TCUTQ3JPDIYOZ5RDARTMHMEKIO2QZQ7IOZ4YC5XV3C2IKYU",
			Alias:  "Firefly",
		},
		aliasedAsset{
			Code:   "XCN",
			Issuer: "GCNY5OXYSY4FKHOPT2SPOQZAOEIGXB5LBYW3HVU3OWSTQITS65M5RCNY",
			Alias:  "Firefly",
		},
		aliasedAsset{
			Code:   "NGN",
			Issuer: "GCC4YLCR7DDWFCIPTROQM7EB2QMFD35XRWEQVIQYJQHVW6VE5MJZXIGW",
			Alias:  "Flutterwave",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GBUYUAI75XXWDZEKLY66CFYKQPET5JR4EENXZBUZ3YXZ7DS56Z4OKOFU",
			Alias:  "Funtracker",
		},
		aliasedAsset{
			Code:   "BCH",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "LTC",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "XRP",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "XAF",
			Issuer: "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
			Alias:  "InterstellarExchange",
		},
		aliasedAsset{
			Code:   "NGN",
			Issuer: "GACA3OWPK26L4SOPKFDSIG2UYCLOYBH32WI3NMEIE6RZ5LRHJMG72ANT",
			Alias:  "KuBitX",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH",
			Alias:  "NaoBTC",
		},
		aliasedAsset{
			Code:   "BRL",
			Issuer: "GDVKY2GU2DRXWTBEYJJWSFXIGBZV6AZNBVVSUHEPZI54LIS6BA7DVVSP",
			Alias:  "nTokens",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GBNLJIYH34UWO5YZFA3A3HD3N76R6DOI33N4JONUOHEEYZYCAYTEJ5AK",
			Alias:  "realio",
		},
		aliasedAsset{
			Code:   "CNY",
			Issuer: "GAREELUB43IRHWEASCFBLKHURCGMHE5IF6XSE7EXDLACYHGRHM43RFOX",
			Alias:  "RippleFox",
		},
		aliasedAsset{
			Code:   "MXN",
			Issuer: "GBUMQHWIQELILQEQ5YEEHUFR6SRLBNRKWHJ3JX7JBRFONG24FWUDG627",
			Alias:  "saldoMX",
		},
		aliasedAsset{
			Code:   "SMX",
			Issuer: "GCDN3VGXZZRCKPG2UEUNR54QDVJRAYINMHBXIT4ZQUFCEQSFN2ZZFSMX",
			Alias:  "saldoMX",
		},
		aliasedAsset{
			Code:   "ARST",
			Issuer: "GCSAZVWXZKWS4XS223M5F54H2B6XPIIXZZGP7KEAIU6YSL5HDRGCI3DG",
			Alias:  "StableX",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GAC63S3QEA5TKKGPLSLGCUSAFASDNO44XDWDNGLPM6UZSBLXA7W47JGZ",
			Alias:  "XCM",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GAC63S3QEA5TKKGPLSLGCUSAFASDNO44XDWDNGLPM6UZSBLXA7W47JGZ",
			Alias:  "XCM",
		},
		aliasedAsset{
			Code:   "GOLD",
			Issuer: "GAJSVHNJ7WNU3M2URGC76NJLL6FDEC5Z2DUJZO24J6MVWLZGE6XIV4PF",
			Alias:  "StellarMetals",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GBVOL67TMUQBGL4TZYNMY3ZQ5WGQYFPFD5VJRWXR72VA33VFNL225PL5",
			Alias:  "Stellarport",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GBVOL67TMUQBGL4TZYNMY3ZQ5WGQYFPFD5VJRWXR72VA33VFNL225PL5",
			Alias:  "Stellarport",
		},
		aliasedAsset{
			Code:   "LTC",
			Issuer: "GBVOL67TMUQBGL4TZYNMY3ZQ5WGQYFPFD5VJRWXR72VA33VFNL225PL5",
			Alias:  "Stellarport",
		},
		aliasedAsset{
			Code:   "XRP",
			Issuer: "GBVOL67TMUQBGL4TZYNMY3ZQ5WGQYFPFD5VJRWXR72VA33VFNL225PL5",
			Alias:  "Stellarport",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GBSTRUSD7IRX73RQZBL3RQUH6KS3O4NYFY3QCALDLZD77XMZOPWAVTUK",
			Alias:  "Stronghold",
		},
		aliasedAsset{
			Code:   "EURT",
			Issuer: "GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S",
			Alias:  "Tempo",
		},
		aliasedAsset{
			Code:   "USD",
			Issuer: "GDSRCV5VTM3U7Y3L6DFRP3PEGBNQMGOWSRTGSBWX6Z3H6C7JHRI4XFJP",
			Alias:  "TokenIO",
		},
		aliasedAsset{
			Code:   "ARS",
			Issuer: "GCYE7C77EB5AWAA25R5XMWNI2EDOKTTFTTPZKM2SR5DI4B4WFD52DARS",
			Alias:  "Anclap",
		},
		aliasedAsset{
			Code:   "JPY",
			Issuer: "GBVAOIACNSB7OVUXJYC5UE2D4YK2F7A24T7EE5YOMN4CE6GCHUTOUQXM",
			Alias:  "VCBear",
		},
		aliasedAsset{
			Code:   "BSV",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "BCH",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "WSEUR",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "WSGBP",
			Issuer: "GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			Alias:  "WhiteStandard",
		},
		aliasedAsset{
			Code:   "BTC",
			Issuer: "GAC63S3QEA5TKKGPLSLGCUSAFASDNO44XDWDNGLPM6UZSBLXA7W47JGZ",
			Alias:  "XCM",
		},
		aliasedAsset{
			Code:   "ETH",
			Issuer: "GAC63S3QEA5TKKGPLSLGCUSAFASDNO44XDWDNGLPM6UZSBLXA7W47JGZ",
			Alias:  "XCM",
		},
		aliasedAsset{
			Code:   "TZS",
			Issuer: "GA2MSSZKJOU6RNL3EJKH3S5TB5CDYTFQFWRYFGUJVIN5I6AOIRTLUHTO",
			Alias:  "ClickPesa",
		},
		aliasedAsset{
			Code:   "ZAR",
			Issuer: "GDYG7OEXT7GO2WOYJKRFMYK6PXQTPFRKO4JSNRRZWE4JM2V6QWQR2QZD",
			Alias:  "Uhuruwallet",
		},
	}
}
