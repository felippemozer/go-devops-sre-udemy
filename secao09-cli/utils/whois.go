package utils

import "github.com/likexian/whois"

func GetWhoIs(host string) string {
	result, _ := whois.Whois(host)
	return result
}
