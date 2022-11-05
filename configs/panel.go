/*
 * @Author: Charley
 * @Date: 2022-11-04 14:19:58
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-04 14:22:09
 * @FilePath: /mospanel/configs/panel.go
 * @Description: Panel Config
 */
package configs

type PanelConfig struct {
	Bind          string `yaml:"bind"`
	Port          int    `yaml:"port"`
	SessionSecret string `yaml:"sessionSecret"`
}
