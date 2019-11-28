package launch

import (
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//遍历十次
	count := 10
	for i := 0; i < count; i++ {
		spider_euroHis()
	}
}
*/
func Before_spider_euroHis(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_euro_his"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroHis() {
	matchLastConfigService := new(service2.MatchLastConfigService)
	config := &entity.MatchLastConfig{}
	config.S = win007.MODULE_FLAG
	config.EuroSpided = false
	matchLastConfigs := matchLastConfigService.Query(config)

	//设置要抓取的波菜公司id
	betCompWin007Ids := []string{"616"}
	//betCompWin007Ids := new(entity2.Comp).FindAllIds()

	processer := proc.GetEuroHisProcesser()
	processer.BetCompWin007Ids = betCompWin007Ids
	processer.MatchLastConfig_list = matchLastConfigs
	processer.Startup()

}