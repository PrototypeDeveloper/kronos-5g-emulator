module kronos

go 1.22.0

require (
	gopkg.in/yaml.v3 v3.0.1
	libgtp v0.0.0-00010101000000-000000000000
	libnas v0.0.0
	libngap v0.0.0
)

replace (
	libgtp => ./lib/gtp
	libnas => ./lib/nas
	libngap => ./lib/ngap
)
