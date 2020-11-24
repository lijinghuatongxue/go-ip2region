package util

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/sirupsen/logrus"
)

func Ip2Area(Ip string) (string, bool) {
	var IpArea string
	region, err := ip2region.New("./IpLib/ip2region.db")
	defer region.Close()
	if err != nil {
		logrus.Error("[util - ip2region] | ❌ false｜IP数据文件读取失败")
		return "Null", false
	}
	ip, err := region.MemorySearch(Ip)
	if err == nil {
		//Info, _ := fmt.Printf("%s-%s-%s", ip.Country, ip.Province, ip.City)
		IpArea := ip.Country + "-" + ip.Province + "-" + ip.City + "-" + ip.ISP
		logrus.Infof("[util - ip2region] | ✅ true｜IP ===> %s | Area ===> %s", Ip, IpArea)
		return IpArea, true
	} else {
		logrus.Errorf("[util - ip2region] | ❌ false| IP ===> %s | Area ===> %s", Ip, IpArea)
		return "Null", false
	}
	return IpArea, true
}
