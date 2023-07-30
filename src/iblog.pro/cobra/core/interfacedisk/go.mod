module disk
go 1.20


replace iblog.pro/cobra/logs => ../iblog.pro/cobra/logs
replace iblog.pro/cobra/core/utils => ./src/iblog.pro/cobra/core/utils

require (
        iblog.pro/cobra/logs v0.0.0-00010101000000-000000000000
        iblog.pro/cobra/core/utils v0.0.0-00010101000000-000000000000

)

