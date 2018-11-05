package jb

/*
func idSentinel(s string) (string, []string) {
	mstr := "vortex"
	hosts := []string{
		"atl-vps-red01.nix.corp.pps.io:26379",
		"atl-vps-red02.nix.corp.pps.io:26379",
		"atl-vps-red03.nix.corp.pps.io:26379",
	}
	if strings.Contains(s, "vss") {
		if strings.Contains(s, "natl") {
			return "vsec", []string{
				"natl-vss-red01.nix.corp.pps.io:26379",
				"natl-vss-red02.nix.corp.pps.io:26379",
				"natl-vss-red03.nix.corp.pps.io:26379",
			}
		}
		return "vsec", []string{
			"atl-vss-red01.nix.corp.pps.io:26379",
			"atl-vss-red02.nix.corp.pps.io:26379",
			"atl-vss-red03.nix.corp.pps.io:26379",
		}
	}
	if strings.Contains(s, "natl") {
		return mstr, []string{
			"atl-vps-red01.nix.corp.pps.io:26379",
			"atl-vps-red02.nix.corp.pps.io:26379",
			"atl-vps-red03.nix.corp.pps.io:26379",
		}
	}
	return mstr, hosts
}

func discoverLocal() string {
	localRegex := regexp.MustCompile("^([na].*-v[ps][cs])")
	host, _ := os.Hostname()
	e := localRegex.FindStringSubmatch(host)
	switch l := len(e); l {
	case 0:
		return host
	case 1:
		host = e[0]
	case 2:
		host = e[1]
	default:
		host = fmt.Sprintf("%s", "ERR")
	}
	return host
}

func discoverRemote(host string) string {
	regex := regexp.MustCompile("^([na].*-v[ps][cs])")
	e := regex.FindStringSubmatch(host)
	switch l := len(e); l {
	case 0:
		return host
	case 1:
		host = e[0]
	case 2:
		host = e[1]
	default:
		host = fmt.Sprintf("%s", "ERR")
	}
	return host
}
*/
