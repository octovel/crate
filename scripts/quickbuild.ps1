param ()

$source = "./cmd/crate/main.go"

function QuickBuild-Project {
	[CmdletBinding()]
	param ()

	if (-not (Test-Path -Path "build")) {
		Write-Host "Output directory not found."
		New-Item -ItemType Directory -Path "./" -Name "build" | Out-Null
		Write-Host "Created a new output directoy: ./build/"
	}

	if (-not (Test-Path -Path $source)) {
		Write-Host "Source file not found, could not quick-build project."
	}

	& go build -o ./build/main.exe $source
}

QuickBuild-Project
