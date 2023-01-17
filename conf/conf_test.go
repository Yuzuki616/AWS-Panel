package conf

import "testing"

func TestConf_SaveConfig(t *testing.T) {
	c := New("./config.json")
	err := c.SaveConfig()
	if err != nil {
		t.Errorf("save config error: %s", err)
	}
}
