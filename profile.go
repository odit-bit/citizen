package citizen

//example of profile JSON

// {
// 	"key": "29",
// 	"alamat": "BINASISWA 1 BLOK D No 19",
// 	"kk": "3275080809090008",
// 	"nama": "DEWI RAMADANI FEBRIANA",
// 	"kelamin": "P",
// 	"ktp": "JATICEMPAKA",
// 	"domisili": "RT05",
// 	"tanggal_lahir": "6 Feb 1994",
// 	"keterangan": "ANAK",
// 	"usia": "28"
// },

type Profile struct {
	Key        string `json:"key"`
	Alamat     string `json:"alamat"`
	KK         string `json:"kk"`
	Nama       string `json:"nama"`
	Kelamin    string `json:"kelamin"`
	KTP        string `json:"ktp"`
	Domisili   string `json:"domisili"`
	Lahir      string `json:"tanggal_lahir" bson:"tanggal_lahir"`
	Keterangan string `json:"keterangan"`
	Usia       string `json:"usia"`
}
