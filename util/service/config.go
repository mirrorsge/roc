package rocserv

// BaseConfig 基础配置
type BaseConfig struct {
	Base struct {
		crossRegisterCenters []string `sep:"," sconf:"crossRegisterCenters"`
	}
}
