package baidufanyi

type Language string

const (
	LanguageAch Language = "ach" // 亚齐
	LanguageAfr Language = "afr" // 南非荷兰
	LanguageAka Language = "aka" // 阿肯
	LanguageAlb Language = "alb" // 阿尔巴尼亚
	LanguageAmh Language = "amh" // 阿姆哈拉
	LanguageAra Language = "ara" // 阿拉伯语
	LanguageArg Language = "arg" // 阿拉贡
	LanguageArm Language = "arm" // 亚美尼亚
	LanguageArq Language = "arq" // 阿尔及利亚阿拉伯语
	LanguageAsm Language = "asm" // 阿萨姆
	LanguageAst Language = "ast" // 阿斯图里亚斯
	LanguageAym Language = "aym" // 艾马拉
	LanguageAze Language = "aze" // 阿塞拜疆
	LanguageBak Language = "bak" // 巴什基尔
	LanguageBal Language = "bal" // 俾路支
	LanguageBaq Language = "baq" // 巴斯克
	LanguageBel Language = "bel" // 白俄罗斯
	LanguageBem Language = "bem" // 本巴
	LanguageBen Language = "ben" // 孟加拉
	LanguageBer Language = "ber" // 柏柏尔
	LanguageBho Language = "bho" // 博杰普尔
	LanguageBis Language = "bis" // 比斯拉马
	LanguageBli Language = "bli" // 比林
	LanguageBos Language = "bos" // 波斯尼亚
	LanguageBre Language = "bre" // 布列塔尼
	LanguageBul Language = "bul" // 保加利亚语
	LanguageBur Language = "bur" // 缅甸
	LanguageCat Language = "cat" // 加泰罗尼亚
	LanguageCeb Language = "ceb" // 宿务
	LanguageChr Language = "chr" // 切罗基
	LanguageCht Language = "cht" // 中文(繁体)
	LanguageChv Language = "chv" // 楚瓦什
	LanguageCor Language = "cor" // 康瓦尔
	LanguageCos Language = "cos" // 科西嘉
	LanguageCre Language = "cre" // 克里克
	LanguageCri Language = "cri" // 克里米亚鞑靼
	LanguageCs  Language = "cs"  // 捷克
	LanguageDan Language = "dan" // 丹麦
	LanguageDe  Language = "de"  // 德
	LanguageDiv Language = "div" // 迪维希
	LanguageEl  Language = "el"  // 希腊
	LanguageEn  Language = "en"  // 英
	LanguageEno Language = "eno" // 古英
	LanguageEpo Language = "epo" // 世界
	LanguageEst Language = "est" // 爱沙尼亚
	LanguageFao Language = "fao" // 法罗
	LanguageFil Language = "fil" // 菲律宾
	LanguageFin Language = "fin" // 芬兰
	LanguageFra Language = "fra" // 法
	LanguageFri Language = "fri" // 弗留利
	LanguageFrm Language = "frm" // 中古法
	LanguageFrn Language = "frn" // 加拿大法
	LanguageFry Language = "fry" // 西弗里斯
	LanguageFul Language = "ful" // 富拉尼
	LanguageGeo Language = "geo" // 格鲁吉亚
	LanguageGla Language = "gla" // 盖尔
	LanguageGle Language = "gle" // 爱尔兰
	LanguageGlg Language = "glg" // 加利西亚
	LanguageGlv Language = "glv" // 曼克斯
	LanguageGra Language = "gra" // 古希腊
	LanguageGrn Language = "grn" // 瓜拉尼
	LanguageGuj Language = "guj" // 古吉拉特
	LanguageHak Language = "hak" // 哈卡钦
	LanguageHau Language = "hau" // 豪萨
	LanguageHaw Language = "haw" // 夏威夷
	LanguageHeb Language = "heb" // 希伯来
	LanguageHi  Language = "hi"  // 印地
	LanguageHil Language = "hil" // 希利盖农
	LanguageHkm Language = "hkm" // 高棉
	LanguageHmn Language = "hmn" // 苗
	LanguageHrv Language = "hrv" // 克罗地亚
	LanguageHt  Language = "ht"  // 海地
	LanguageHu  Language = "hu"  // 匈牙利
	LanguageHup Language = "hup" // 胡帕
	LanguageIbo Language = "ibo" // 伊博
	LanguageIce Language = "ice" // 冰岛
	LanguageId  Language = "id"  // 印尼
	LanguageIdo Language = "ido" // 伊多
	LanguageIku Language = "iku" // 伊努克提图特
	LanguageIna Language = "ina" // 因特
	LanguageIng Language = "ing" // 印古什
	LanguageIr  Language = "ir"  // 伊朗
	LanguageIt  Language = "it"  // 意大利
	LanguageJav Language = "jav" // 爪哇
	LanguageJp  Language = "jp"  // 日
	LanguageKab Language = "kab" // 卡拜尔
	LanguageKah Language = "kah" // 卡舒比
	LanguageKal Language = "kal" // 格陵兰
	LanguageKan Language = "kan" // 卡纳达
	LanguageKas Language = "kas" // 克什米尔
	LanguageKau Language = "kau" // 卡努里
	LanguageKin Language = "kin" // 卢旺达
	LanguageKir Language = "kir" // 吉尔吉斯
	LanguageKli Language = "kli" // 克林贡
	LanguageKok Language = "kok" // 孔卡尼
	LanguageKon Language = "kon" // 刚果
	LanguageKor Language = "kor" // 韩
	LanguageKur Language = "kur" // 库尔德
	LanguageLag Language = "lag" // 拉特加莱
	LanguageLao Language = "lao" // 老挝
	LanguageLat Language = "lat" // 拉丁
	LanguageLav Language = "lav" // 拉脱维亚
	LanguageLim Language = "lim" // 林堡
	LanguageLin Language = "lin" // 林加拉
	LanguageLit Language = "lit" // 立陶宛
	LanguageLog Language = "log" // 低地德
	LanguageLoj Language = "loj" // 逻辑
	LanguageLos Language = "los" // 下索布
	LanguageLtz Language = "ltz" // 卢森堡
	LanguageLug Language = "lug" // 卢干达
	LanguageMac Language = "mac" // 马其顿
	LanguageMah Language = "mah" // 马绍尔
	LanguageMai Language = "mai" // 迈蒂利
	LanguageMal Language = "mal" // 马拉雅拉姆
	LanguageMao Language = "mao" // 毛利
	LanguageMar Language = "mar" // 马拉地
	LanguageMau Language = "mau" // 毛里求斯克里奥尔
	LanguageMay Language = "may" // 马来
	LanguageMg  Language = "mg"  // 马拉加斯
	LanguageMlt Language = "mlt" // 马耳他
	LanguageMot Language = "mot" // 黑山
	LanguageNbl Language = "nbl" // 南恩德贝莱
	LanguageNea Language = "nea" // 那不勒斯
	LanguageNep Language = "nep" // 尼泊尔
	LanguageNl  Language = "nl"  // 荷兰
	LanguageNno Language = "nno" // 新挪威
	LanguageNob Language = "nob" // 书面挪威
	LanguageNor Language = "nor" // 挪威
	LanguageNqo Language = "nqo" // 西非书面
	LanguageNya Language = "nya" // 齐切瓦
	LanguageOci Language = "oci" // 奥克
	LanguageOji Language = "oji" // 奥杰布瓦
	LanguageOri Language = "ori" // 奥里亚
	LanguageOrm Language = "orm" // 奥罗莫
	LanguageOss Language = "oss" // 奥塞梯
	LanguagePam Language = "pam" // 邦板牙
	LanguagePan Language = "pan" // 旁遮普
	LanguagePap Language = "pap" // 帕皮阿门托
	LanguagePed Language = "ped" // 北索托
	LanguagePer Language = "per" // 波斯
	LanguagePl  Language = "pl"  // 波兰
	LanguagePot Language = "pot" // 巴西葡萄牙
	LanguagePt  Language = "pt"  // 葡萄牙
	LanguagePus Language = "pus" // 普什图
	LanguageQue Language = "que" // 克丘亚
	LanguageRo  Language = "ro"  // 罗姆
	LanguageRoh Language = "roh" // 罗曼什
	LanguageRom Language = "rom" // 罗马尼亚
	LanguageRu  Language = "ru"  // 俄
	LanguageRuy Language = "ruy" // 卢森尼亚
	LanguageSan Language = "san" // 梵
	LanguageSec Language = "sec" // 塞尔维亚-克罗地亚
	LanguageSha Language = "sha" // 掸
	LanguageSil Language = "sil" // 西里西亚
	LanguageSin Language = "sin" // 僧伽罗
	LanguageSk  Language = "sk"  // 斯洛伐克
	LanguageSlo Language = "slo" // 斯洛文尼亚
	LanguageSm  Language = "sm"  // 萨摩亚
	LanguageSme Language = "sme" // 北方萨米
	LanguageSna Language = "sna" // 修纳
	LanguageSnd Language = "snd" // 信德
	LanguageSol Language = "sol" // 桑海
	LanguageSom Language = "som" // 索马里
	LanguageSot Language = "sot" // 南索托
	LanguageSpa Language = "spa" // 西班牙
	LanguageSrc Language = "src" // 塞尔维亚语（西里尔）
	LanguageSrd Language = "srd" // 萨丁尼亚
	LanguageSrp Language = "srp" // 塞尔维亚
	LanguageSun Language = "sun" // 巽他
	LanguageSwa Language = "swa" // 斯瓦希里
	LanguageSwe Language = "swe" // 瑞典
	LanguageSyr Language = "syr" // 叙利亚
	LanguageTam Language = "tam" // 泰米尔
	LanguageTat Language = "tat" // 鞑靼
	LanguageTel Language = "tel" // 泰卢固
	LanguageTet Language = "tet" // 德顿
	LanguageTgk Language = "tgk" // 塔吉克
	LanguageTgl Language = "tgl" // 他加禄
	LanguageTh  Language = "th"  // 泰
	LanguageTir Language = "tir" // 提格利尼亚
	LanguageTr  Language = "tr"  // 土耳其
	LanguageTso Language = "tso" // 聪加
	LanguageTua Language = "tua" // 突尼斯阿拉伯
	LanguageTuk Language = "tuk" // 土库曼
	LanguageTwi Language = "twi" // 契维
	LanguageUkr Language = "ukr" // 乌克兰
	LanguageUps Language = "ups" // 高地索布
	LanguageUrd Language = "urd" // 乌尔都
	LanguageVen Language = "ven" // 文达
	LanguageVie Language = "vie" // 越南
	LanguageWel Language = "wel" // 威尔士
	LanguageWln Language = "wln" // 瓦隆
	LanguageWol Language = "wol" // 沃洛夫
	LanguageWyw Language = "wyw" // 中文(文言文)
	LanguageXho Language = "xho" // 科萨
	LanguageYid Language = "yid" // 意第绪
	LanguageYor Language = "yor" // 约鲁巴
	LanguageYue Language = "yue" // 中文(粤语)
	LanguageZaz Language = "zaz" // 扎扎其
	LanguageZh  Language = "zh"  // 中文(简体)
	LanguageZul Language = "zul" // 祖鲁

)