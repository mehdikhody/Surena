package config

import (
	"fmt"
	"github.com/xtls/xray-core/common/errors"
	"surena/node/utils"
)

func (c *Config) GetAPIAddress() (string, error) {
	if c.Config.API == nil {
		c.Logger.Warn("API is not enabled")
		return "", errors.New("API is not enabled")
	}

	if c.Config.API.Tag == "" {
		c.Logger.Warn("API tag is not set")
		return "", errors.New("API tag is not set")
	}

	router, err := c.Config.RouterConfig.Build()
	if err != nil {
		c.Logger.Error("failed to build router config")
		return "", err
	}

	var inboundsTag []string
	for _, rule := range router.Rule {
		if rule.GetTag() != c.Config.API.Tag {
			continue
		}

		inboundsTag = rule.GetInboundTag()
	}

	for _, inbound := range c.Config.InboundConfigs {
		isDokodemo := inbound.Protocol == "dokodemo-door"
		isIncluded := utils.Include(inboundsTag, inbound.Tag)

		c.Logger.Tracef("Protocol: %s", inbound.Protocol)
		c.Logger.Tracef("Tag: %s", inbound.Tag)
		c.Logger.Tracef("Tags: %t", inboundsTag)
		c.Logger.Tracef("Included: %t", isIncluded)

		if isIncluded && isDokodemo {
			host := inbound.ListenOn.IP().String()
			port := inbound.PortList.Build().GetRange()[0].GetFrom()

			c.Logger.Debugf("API address: %s:%d", host, port)
			return fmt.Sprintf("%s:%d", host, port), nil
		}
	}

	c.Logger.Warn("API address not found")
	return "", errors.New("API address not found")
}
