package ip

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

// GetLocation 根据传入的 IP 地址查询归属地（省份-城市）。
//
// 处理规则：
//  1. 当 ip 为 "127.0.0.1" 或 "localhost" 时，直接返回 "内部IP"。
//  2. 通过高德 IP 定位接口查询公网 IP 对应地理位置。
//  3. 若查询结果中省份为空，则返回 "未知位置"。
//
// 返回值示例："广东省-深圳市"。
func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=21d225447dc9bf4ff040218eefcef219")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}

// GetLocaHonst 获取当前机器可用的局域网 IPv4 地址。
//
// 实现逻辑：
//  1. 遍历所有网络接口，仅处理状态为 up 的接口。
//  2. 遍历接口地址，过滤掉回环地址（如 127.0.0.1）。
//  3. 返回第一个匹配到的 IPv4 字符串。
//  4. 若未找到可用 IPv4，则返回空字符串。
func GetLocaHonst() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
