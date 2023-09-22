package dns

import (
	"fmt"
	"log"
	"net"
)

func DNSFind() {
	// 要查询的域名
	domainName := "www.google.com"

	// 使用 LookupHost 函数查询 IP 地址
	ips, err := net.LookupHost(domainName)
	if err != nil {
		log.Fatal(err)
	}

	// 打印查询到的所有 IP 地址
	fmt.Printf("IP addresses for %s:\n", domainName)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
