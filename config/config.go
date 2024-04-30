package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"kk_cookie"`b-user-id=97011768-bd83-a9a6-ffbe-16f9f7ff3309; __wpkreporterwid_=c8aa6073-92be-4691-a0c6-0c3de76ac383; __itrace_wid=6020202f-a564-44e3-1f4d-13c48cccee58; user_device_id=d5332d3f78ca4256a26a5976dca05d34; user_device_id_timestamp=1712209372330; _UP_A4A_11_=wb9641d358984cfb9fe29240aaa50bdb; _UP_D_=pc; _UP_30C_6A_=st9646201992vovnkt4nso0pig3ga4bk; _UP_TS_=sg1307817fb39a5d5896a3c4504ead8c203; _UP_E37_B7_=sg1307817fb39a5d5896a3c4504ead8c203; _UP_TG_=st9646201992vovnkt4nso0pig3ga4bk; _UP_335_2B_=1; __pus=c1decf90ae309d4b5211118caa6ed815AARbcA6141ELsH788tvTO2n4/zFEq95x4BCp5n1YKXOYNBMHH3ozIh3aRphPNT3Gc5mAw6CeTzcmxoFp5jw9Si2i; __kp=8f3b5f80-0609-11ef-ab72-01dce42b0b3b; __kps=AAS6jYFBQVdLfn6Zlm5QmQcY; __ktd=oxUytwohjP4pIrZho18AZA==; __uid=AAS6jYFBQVdLfn6Zlm5QmQcY; tfstk=fjvSNEsnoz4SuaKf-_nVhQFbjMXIbUMZVkspjHezvTB8pethugPrZ6uBGeLVyLllzE_BbH2Fz3W-AeT9P68Bx0jdJeLC8QlqQ3xlK93wRAkwqyJdimL5w9ExvMsQJMvhwx-lK93qrd8Ca3Y1oBWKd9nfHMS_99epJZKAoMjLyMe-cnQcv9IdygQYMgj1JyF-1v0fmb_JV2z7nA6SjxK5lJe_jG5-K3NUCROVV_tW2Z15NKs5NN5VLVopeE-pULfm9W6H0CTB9F34aad1fE1wNVwW5UjpkiLrYWjXOHpGE_ZaQntXvC5X3cHGFTQJTTxUf4-XFU-khHrLMT6yG195BA4R8I-6dM9n-xT6vhAdGpa54uyN5dcYOo1gpi_ZcmN3TFFAHfdMr4i5wij2QmibPWfRmi_ZcmN3t_Ic0FnjczNh.; ctoken=oBIDa2YMYd75kVwbUatIIwyA; grey-id=369d316c-1b47-b2f8-91bc-3bf2b63aa653; grey-id.sig=tB1X8f2ZafjtmCKMsxzC7ujszFFX-mxqCja9BAlM64Y; isQuark=false; isQuark.sig=DWPHMZYiiwQ-v58AbcP-rBdSIpzO8ZnrD67BdJuPatU; __puus=60b29f3a26adf00a7fabd4b1ae8e71f5AARvjH7XNv7AN8qJ01msFrCRGqInti8lr/9Wfdnx0vlps79Wm9rI4U08eRc+uBJsrzZGfy2qkZyTzsfDA1bOIw6WZW85nasXj6bRUYRiZu0sXBdrRpe0qaW5Co+i8UI6Jm1gOdktT05dEHIKHvKXk2PAJa9D/14WA7Q+Yg8d3bVNy8OeKX+AkPm1MDwQbqgAuFFU/31sv+3kgT0bKeiRAus5
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
