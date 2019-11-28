package mysql

import (
	"encoding/json"
	"log"
	entity4 "tesou.io/platform/foot-parent/foot-api/module/analy/entity"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	entity1 "tesou.io/platform/foot-parent/foot-api/module/match/entity"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/odds/entity"
)

type DBOpsService struct {
}

/**
 * 清空表
 */
func (this *DBOpsService) TruncateTable(tables []string) {
	engine := GetEngine()
	for _, v := range tables {
		_, err := engine.Exec(" TRUNCATE TABLE " + v)
		if nil != err {
			log.Println(err)
		}
	}
}

/**
 * xorm支持获取表结构信息，通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息。
 */
func (this *DBOpsService) DBMetas() string {
	engine := GetEngine()
	dbMetas, err := engine.DBMetas()
	if nil != err {
		log.Println(err.Error())
	}
	bytes, _ := json.Marshal(dbMetas)
	result := string(bytes)
	return result
}

/**
 * 同步生成数据库表
 */
func (this *DBOpsService) SyncTableStruct() {
	engine := GetEngine()
	var err error
	//sync model
	//比赛相关表
	err = engine.Sync2(new(entity1.MatchLast), new(entity1.MatchLastConfig), new(entity1.MatchHis))
	if nil != err {
		log.Println(err.Error())
	}

	//赔率相关表
	err = engine.Sync2(new(entity2.EuroLast), new(entity2.EuroHis), new(entity2.AsiaLast), new(entity2.AsiaHis))
	if nil != err {
		log.Println(err.Error())
	}

	//波菜公司，联赛其他数据表
	err = engine.Sync2(new(entity3.Comp), new(entity3.CompConfig), new(entity3.League), new(entity3.LeagueConfig))
	if nil != err {
		log.Println(err.Error())
	}
	//分析的结果集表
	err = engine.Sync2(new(entity4.AnalyResult))
	if nil != err {
		log.Println(err.Error())
	}

}