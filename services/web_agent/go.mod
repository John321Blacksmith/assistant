module osint_agent/services/web_agent

go 1.24.9

require (
	osint_agent/libs/classifier v0.0.0-00010101000000-000000000000
	osint_agent/libs/logging/local v0.0.0-00010101000000-000000000000
	osint_agent/libs/logging/prometheus v0.0.0-00010101000000-000000000000
)

replace osint_agent/libs/logging/local => ../../libs/logging/local/

replace osint_agent/libs/classifier => ../../libs/classifier/

replace osint_agent/libs/logging/prometheus => ../../libs/logging/prometheus/
