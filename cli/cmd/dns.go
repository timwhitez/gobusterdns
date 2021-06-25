package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"github.com/timwhitez/gobusterdns/cli"
	"github.com/timwhitez/gobusterdns/gobusterdns"
	"github.com/timwhitez/gobusterdns/libgobuster"
	"github.com/spf13/cobra"
)

// nolint:gochecknoglobals
var cmdDNS *cobra.Command

func runDNS(cmd *cobra.Command, args []string) error {
	globalopts, pluginopts, err := parseDNSOptions()
	if err != nil {
		return fmt.Errorf("error on parsing arguments: %w", err)
	}

	if pluginopts.Domainlist != "" {
		//fmt.Println(pluginopts.Domainlist)
		lines := getContent(pluginopts.Domainlist)
		for _, value := range lines{
			pluginopts.Domain = value
			fmt.Println(value+"::")
			plugin, err := gobusterdns.NewGobusterDNS(globalopts, pluginopts)
			if err != nil {
				return fmt.Errorf("error on creating gobusterdns: %w", err)
			}
			if err := cli.Gobuster(mainContext, globalopts, plugin); err != nil {
				var wErr *gobusterdns.ErrWildcard
				fmt.Println(pluginopts.Domain)
				if errors.As(err, &wErr) {
					fmt.Println("%w. To force processing of Wildcard DNS, specify the '--wildcard' switch", wErr)
					continue
				}
				fmt.Println("error on running gobuster: %w", err)
				continue
			}else{
				fmt.Errorf("error on running gobuster: %w", err)
			}

		}
	}else {
		fmt.Println(pluginopts.Domain)
		plugin, err := gobusterdns.NewGobusterDNS(globalopts, pluginopts)
		if err != nil {
			return fmt.Errorf("error on creating gobusterdns: %w", err)
		}
		if err := cli.Gobuster(mainContext, globalopts, plugin); err != nil {
			var wErr *gobusterdns.ErrWildcard
			if errors.As(err, &wErr) {
				return fmt.Errorf("%w. To force processing of Wildcard DNS, specify the '--wildcard' switch", wErr)
			}
			return fmt.Errorf("error on running gobuster: %w", err)
		}
	}
	return nil
}

func parseDNSOptions() (*libgobuster.Options, *gobusterdns.OptionsDNS, error) {
	globalopts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	plugin := gobusterdns.NewOptionsDNS()

	plugin.Domainlist, err = cmdDNS.Flags().GetString("domainlist")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for domainlist: %w", err)
	}

	plugin.Domain, err = cmdDNS.Flags().GetString("domain")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for domain: %w", err)
	}

	plugin.ShowIPs, err = cmdDNS.Flags().GetBool("show-ips")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for show-ips: %w", err)
	}

	plugin.ShowCNAME, err = cmdDNS.Flags().GetBool("show-cname")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for show-cname: %w", err)
	}

	plugin.WildcardForced, err = cmdDNS.Flags().GetBool("wildcard")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for wildcard: %w", err)
	}

	plugin.Timeout, err = cmdDNS.Flags().GetDuration("timeout")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for timeout: %w", err)
	}

	plugin.Resolver, err = cmdDNS.Flags().GetString("resolver")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for resolver: %w", err)
	}

	if plugin.Resolver != "" && runtime.GOOS == "windows" {
		return nil, nil, fmt.Errorf("currently can not set custom dns resolver on windows. See https://golang.org/pkg/net/#hdr-Name_Resolution")
	}

	return globalopts, plugin, nil
}

// nolint:gochecknoinits
func init() {
	cmdDNS = &cobra.Command{
		Use:   "dns",
		Short: "Uses DNS subdomain enumeration mode",
		RunE:  runDNS,
	}
	cmdDNS.Flags().StringP("domainlist", "l", "", "The target list file")
	cmdDNS.Flags().StringP("domain", "d", "", "The target domain")
	cmdDNS.Flags().BoolP("show-ips", "i", false, "Show IP addresses")
	cmdDNS.Flags().BoolP("show-cname", "c", false, "Show CNAME records (cannot be used with '-i' option)")
	cmdDNS.Flags().DurationP("timeout", "", time.Second, "DNS resolver timeout")
	cmdDNS.Flags().BoolP("wildcard", "", false, "Force continued operation when wildcard found")
	cmdDNS.Flags().StringP("resolver", "r", "", "Use custom DNS server (format server.com or server.com:port)")
	/*if err1 := cmdDNS.MarkFlagRequired("domain"); err1 != nil {
		log.Fatalf("error on marking flag as required: %v", err1)
	}*/
	cmdDNS.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		configureGlobalOptions()
	}

	rootCmd.AddCommand(cmdDNS)
}


func getContent (filename string) []string{
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines :=strings.Split(string(content), "\r\n")
	return lines
}