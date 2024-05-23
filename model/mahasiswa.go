package model

type DataMahasiswa struct {
	ID          string `json:"userid"`
	NamaLengkap string `json:"namalengkap"`
	Prodi       string `json:"prodi"`
	Alamat      string `json:"alamat"`
	AsalSekolah string `json:"asalsekolah"`
}
