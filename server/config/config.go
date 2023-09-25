package config
// 系统配置

type NEBULA struct{
	SERVER		SERVER 	// 系统配置 
	ZAP 		ZAP		// zap配置
	MySQL		MySQL	// MySQL配置
	PGSQL		Pgsql	// postgresql配置
	SQLITE		Sqlite	// sqlite配置
	REDIS 		Redis	// Redis 配置
	JWT			JWT		// JWT配置
}