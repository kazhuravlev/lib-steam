package meta

// AppID - идентфиикатор приложения steam
type AppID uint

const (
	// AppSteam - Steam
	AppSteam AppID = iota
	// AppCSGO - Counter-Strike: Global Offensive
	AppCSGO
	// AppDota2 - Dota 2
	AppDota2
	// AppTF2 - Team Fortress 2
	AppTF2
	// AppAltitude0 - Altitude0: Lower &amp; Faster
	AppAltitude0
	// AppArmello - Armello
	AppArmello
	// AppBattleBlock - BattleBlock Theater
	AppBattleBlock
	// AppDontStarveTogether - Don't Starve Together Beta
	AppDontStarveTogether
	// AppH1Z1JustSurvive - H1Z1: Just Survive
	AppH1Z1JustSurvive
	// AppH1Z1KingOfTheKill - H1Z1: King of the Kill
	AppH1Z1KingOfTheKill
	// AppImmune - Immune
	AppImmune
	// AppKillingFloor2 - Killing Floor 2
	AppKillingFloor2
	// AppMinimum - Minimum
	AppMinimum
	// AppMoveOrDie - Move or Die
	AppMoveOrDie
	// AppNaturalSelection2 - Natural Selection 2
	AppNaturalSelection2
	// AppPayDay2 - PAYDAY 2
	AppPayDay2
	// AppPathOfExile - Path of Exile
	AppPathOfExile
	// AppPrimalCarnageExtinction - Primal Carnage: Extinction
	AppPrimalCarnageExtinction
	// AppRatzInstagib2 - Ratz Instagib 2.0
	AppRatzInstagib2
	// AppReflex - Reflex
	AppReflex
	// AppRobotRollerDerby - Robot Roller-Derby Disco Dodgeball
	AppRobotRollerDerby
	// AppRust - Rust
	AppRust
	// AppSinsOfADarkAge - Sins of a Dark Age
	AppSinsOfADarkAge
	// AppSubnautica - Subnautica
	AppSubnautica
	// AppMightyQuestForEpicLoot - The Mighty Quest For Epic Loot
	AppMightyQuestForEpicLoot
	// AppUnturned - Unturned
	AppUnturned
	// AppWarframe - Warframe
	AppWarframe
	// AppZombieGrinder - Zombie Grinder
	AppZombieGrinder
)

// Apps - список всех приложений steam, доступных в маркете
var Apps = map[AppID]int{
	AppSteam:                   753,
	AppCSGO:                    730,
	AppDota2:                   570,
	AppTF2:                     440,
	AppAltitude0:               308080,
	AppArmello:                 290340,
	AppBattleBlock:             238460,
	AppDontStarveTogether:      322330,
	AppH1Z1JustSurvive:         295110,
	AppH1Z1KingOfTheKill:       433850,
	AppImmune:                  348670,
	AppKillingFloor2:           232090,
	AppMinimum:                 214190,
	AppMoveOrDie:               323850,
	AppNaturalSelection2:       4920,
	AppPayDay2:                 218620,
	AppPathOfExile:             238960,
	AppPrimalCarnageExtinction: 321360,
	AppRatzInstagib2:           338170,
	AppReflex:                  328070,
	AppRobotRollerDerby:        270450,
	AppRust:                    252490,
	AppSinsOfADarkAge:          251970,
	AppSubnautica:              264710,
	AppMightyQuestForEpicLoot:  239220,
	AppUnturned:                304930,
	AppWarframe:                230410,
	AppZombieGrinder:           263920,
}

// ReverseApps - список всех приложений steam, доступных в маркете
var ReverseApps = map[int]AppID{
	753:    AppSteam,
	730:    AppCSGO,
	570:    AppDota2,
	440:    AppTF2,
	308080: AppAltitude0,
	290340: AppArmello,
	238460: AppBattleBlock,
	322330: AppDontStarveTogether,
	295110: AppH1Z1JustSurvive,
	433850: AppH1Z1KingOfTheKill,
	348670: AppImmune,
	232090: AppKillingFloor2,
	214190: AppMinimum,
	323850: AppMoveOrDie,
	4920:   AppNaturalSelection2,
	218620: AppPayDay2,
	238960: AppPathOfExile,
	321360: AppPrimalCarnageExtinction,
	338170: AppRatzInstagib2,
	328070: AppReflex,
	270450: AppRobotRollerDerby,
	252490: AppRust,
	251970: AppSinsOfADarkAge,
	264710: AppSubnautica,
	239220: AppMightyQuestForEpicLoot,
	304930: AppUnturned,
	230410: AppWarframe,
	263920: AppZombieGrinder,
}
