package language

const  CHINESE_SUCCESS = "成功"
const ENGLISH_SUCCESS = "Success"
const FRENCH_SUCCESS = "Succès"
const JAPANESE_SUCCESS = "せいこう"
const GERMAN_SUCCESS = "Erfolg"

const CHINESE_FAIL = "失敗"
const ENGLISH_FAIL = "Fail"
const FRENCH_FAIL = "Échec"
const JAPANESE_FAIL = "に失敗"
const GERMAN_FAIL = "fehlschlagen"

const CHINESE_PARAMS_WRONG = "參數錯誤"
const ENGLISH_PARAMS_WRONG = "Params wrong"
const FRENCH_PARANS_WRONG = "Erreur de paramètre"
const JAPANESE_PARAMS_WRONG = "パラメータエラー"
const GERMAN_PARAMS_WRONG = "Parameterfehler"

const CHINESETIME = "处于申购处理时间不能参与"
const ENGLISHTIME = "Cannot participate during subscription processing time"
const FRENCHTIME = "Impossible de participer pendant le temps de traitement de la demande"
const JAPANESETIME = "発注処理時間中は参加できません"
const GERMANTIME = "Kann während der Bearbeitungszeit des Abonnements nicht teilnehmen"

const CHINESE_ACOUNT = "用户不存在,请创建或转账到指定地址创建"
const ENGLISH_ACOUNT = "User does not exist. Please create or transfer to the specified address to create it\n\n"
const FRENCH_ACOUNT = "L'utilisateur n'existe pas, veuillez créer ou transférer à l'adresse indiquée créer"
const JAPANESE_ACOUNT = "ユーザーは存在しません。指定したアドレスの作成に作成または振り替えてください"
const GERMAN_ACOUNT = "Benutzer existiert nicht. Bitte erstellen oder an die angegebene Adresse übertragen"

const CHINESE_BLANCE = "该地址账户余额不足请充值"
const ENGLISH_BLANCE = "The account balance at this address is insufficient. Please recharge"
const FRENCH_BLANCE = "Le solde du compte à cette adresse est insuffisant Veuillez recharger"
const JAPANESE_BLANCE = "当該住所口座の残高が不足している場合はチャージしてください"
const GERMAN_BLANCE= "Der Kontostand an dieser Adresse ist unzureichend. Bitte laden Sie"

const CHINESE_REPEAT = "本场次已参与，请不要重复参与"
const ENGLISH_REPEAT = "This session has already been participated in, please do not participate again"
const FRENCH_REPEAT = "Ce lieu a déjà participé, ne répétez pas votre participation"
const JAPANESE_REPEAT = "今回は参加しましたので、重複しないでください"
const GERMAN_REPEAT= "Diese Session wurde bereits besucht, bitte nicht mehr teilnehmen"

const CHINESE_FREQUENCY = "提现访问过于频繁，稍后再试"
const ENGLISH_FREQUENCY = "Withdrawal access is too frequent, try again later"
const FRENCH_FREQUENCY = "L'accès au retrait est trop fréquent, réessayez plus tard"
const JAPANESE_FREQUENCY = "出金アクセスが多すぎます。後でもう一度お試しください"
const GERMAN_FREQUENCY= "Auszahlungszugriff ist zu häufig, versuchen Sie es später erneut"

type Language struct {
	Cn CN
	En EN
	Jp JP
	Fr FR
	Gr GR
	country string
}

type CN struct {
	Success string
	Fail string
	ParaWrong string
	TimeWrong string
	AccountWrong string
	BalanceWrong string
	RepeatWrong string
	FrequencyWrong string
}

type EN struct {
	Success string
	Fail string
	ParaWrong string
	TimeWrong string
	AccountWrong string
	BalanceWrong string
	RepeatWrong string
	FrequencyWrong string
}
type JP struct {
	Success string
	Fail string
	ParaWrong string
	TimeWrong string
	AccountWrong string
	BalanceWrong string
	RepeatWrong string
	FrequencyWrong string
}
type FR struct {
	Success string
	Fail string
	ParaWrong string
	TimeWrong string
	AccountWrong string
	BalanceWrong string
	RepeatWrong string
	FrequencyWrong string
}
type GR struct {
	Success string
	Fail string
	ParaWrong string
	TimeWrong string
	AccountWrong string
	BalanceWrong string
	RepeatWrong string
	FrequencyWrong string
}

func GetLangPackage(country string) *Language {
	l := Language{
		Cn: CN{Success: CHINESE_SUCCESS,
			Fail: CHINESE_FAIL,
			ParaWrong: CHINESE_PARAMS_WRONG,
			TimeWrong: CHINESETIME,
			AccountWrong: CHINESE_ACOUNT,
			BalanceWrong: CHINESE_BLANCE,
			RepeatWrong: CHINESE_REPEAT,
			FrequencyWrong: CHINESE_FREQUENCY,
			},

		En: EN{Success: ENGLISH_SUCCESS,
			Fail: ENGLISH_FAIL,
			ParaWrong: ENGLISH_PARAMS_WRONG,
			TimeWrong: ENGLISHTIME,
			AccountWrong: ENGLISH_ACOUNT,
			BalanceWrong: ENGLISH_BLANCE,
			RepeatWrong: ENGLISH_REPEAT,
			FrequencyWrong: ENGLISH_FREQUENCY,
			},

		Jp: JP{Success: JAPANESE_SUCCESS,
			Fail: JAPANESE_FAIL,
			ParaWrong: JAPANESE_PARAMS_WRONG,
			TimeWrong: JAPANESETIME,
			AccountWrong: JAPANESE_ACOUNT,
			BalanceWrong: JAPANESE_BLANCE,
			RepeatWrong: JAPANESE_REPEAT,
			FrequencyWrong: JAPANESE_FREQUENCY},

		Fr: FR{Success: FRENCH_SUCCESS,
			Fail: FRENCH_FAIL,
			ParaWrong: FRENCH_PARANS_WRONG,
			TimeWrong: FRENCHTIME,
			AccountWrong: FRENCH_ACOUNT,
			BalanceWrong: FRENCH_BLANCE,
			RepeatWrong: FRENCH_REPEAT,
			FrequencyWrong: FRENCH_FREQUENCY},


		Gr: GR{Success: GERMAN_SUCCESS,
			Fail: GERMAN_FAIL,
			ParaWrong: GERMAN_PARAMS_WRONG,
			TimeWrong: GERMANTIME,
			AccountWrong: GERMAN_ACOUNT,
			BalanceWrong: GERMAN_BLANCE,
			RepeatWrong: GERMAN_REPEAT,
			FrequencyWrong: GERMAN_FREQUENCY},
	}
	l.country = country
	return &l
}

func (lang *Language) GetSuccess ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.Success
		break
	case "en-us":
		x = lang.En.Success
		break
	case "jp":
		x = lang.Jp.Success
		break
	case "fr":
		x = lang.Fr.Success
		break
	case "gr":
		x = lang.Gr.Success
		break
	default:
		x = lang.En.Success
	}
	return x
}

func (lang *Language) GetParaWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.ParaWrong
		break
	case "en-us":
		x = lang.En.ParaWrong
		break
	case "jp":
		x = lang.Jp.ParaWrong
		break
	case "fr":
		x = lang.Fr.ParaWrong
		break
	case "gr":
		x = lang.Gr.ParaWrong
		break
	default:
		x = lang.En.ParaWrong
	}
	return x
}

func (lang *Language) GetFail ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.Fail
		break
	case "en-us":
		x = lang.En.Fail
		break
	case "jp":
		x = lang.Jp.Fail
		break
	case "fr":
		x = lang.Fr.Fail
		break
	case "gr":
		x = lang.Gr.Fail
		break
	default:
		x = lang.En.Fail
	}
	return x
}

func (lang *Language) GetTimeWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.TimeWrong
		break
	case "en-us":
		x = lang.En.TimeWrong
		break
	case "jp":
		x = lang.Jp.TimeWrong
		break
	case "fr":
		x = lang.Fr.TimeWrong
		break
	case "gr":
		x = lang.Gr.TimeWrong
		break
	default:
		x = lang.En.TimeWrong
	}
	return x
}

func (lang *Language) GetAccountWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.AccountWrong
		break
	case "en-us":
		x = lang.En.AccountWrong
		break
	case "jp":
		x = lang.Jp.AccountWrong
		break
	case "fr":
		x = lang.Fr.AccountWrong
		break
	case "gr":
		x = lang.Gr.AccountWrong
		break
	default:
		x = lang.En.AccountWrong
	}
	return x
}

func (lang *Language) GetBalanceWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.BalanceWrong
		break
	case "en-us":
		x = lang.En.BalanceWrong
		break
	case "jp":
		x = lang.Jp.BalanceWrong
		break
	case "fr":
		x = lang.Fr.BalanceWrong
		break
	case "gr":
		x = lang.Gr.BalanceWrong
		break
	default:
		x = lang.En.BalanceWrong
	}
	return x
}
func (lang *Language) GetRepeatWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.RepeatWrong
		break
	case "en-us":
		x = lang.En.RepeatWrong
		break
	case "jp":
		x = lang.Jp.RepeatWrong
		break
	case "fr":
		x = lang.Fr.RepeatWrong
		break
	case "gr":
		x = lang.Gr.RepeatWrong
		break
	default:
		x = lang.En.RepeatWrong
	}
	return x
}
func (lang *Language) GetFreWrong ()string {
	var x =""
	switch lang.country {
	case "zh-cn":
		x=lang.Cn.FrequencyWrong
		break
	case "en-us":
		x = lang.En.FrequencyWrong
		break
	case "jp":
		x = lang.Jp.FrequencyWrong
		break
	case "fr":
		x = lang.Fr.FrequencyWrong
		break
	case "gr":
		x = lang.Gr.FrequencyWrong
		break
	default:
		x = lang.En.FrequencyWrong
	}
	return x
}
