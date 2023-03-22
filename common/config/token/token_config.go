package token

var TokenCfg *TokenConfig // token配置

type TokenConfig struct {
	IgnorePaths []string
}

// AddIgnorePath 增加token不校验路径
func (config *TokenConfig) AddIgnorePath(ignorePath string) *TokenConfig {
	config.IgnorePaths = append(config.IgnorePaths, ignorePath)
	return config
}

// TokenIgnorePath token不校验路径集
func (config *TokenConfig) TokenIgnorePath() {
	config.AddIgnorePath("/token/*").
		AddIgnorePath("/ping").AddIgnorePath("/user/*")
}

// InitTokenConfig 初始化token配置
func InitTokenConfig() {
	TokenCfg = &TokenConfig{
		IgnorePaths: make([]string, 0),
	}
	TokenCfg.TokenIgnorePath()
}
