param(
    [string]$VariableName = "DRESSY_LOCAL_PROTO_PATH"
)

# Check admin rights
if (-not ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Error "Run only with administrator rights!"
    exit 1
}

# Path definition
$ScriptPath = Split-Path -Parent $MyInvocation.MyCommand.Definition

# Variable set
$VariableValue = $ScriptPath
Write-Host "Adding/updating environment variable $VariableName = $VariableValue..."
[Environment]::SetEnvironmentVariable($VariableName, $VariableValue, "Machine")

# Check
$updatedValue = [Environment]::GetEnvironmentVariable($VariableName, "Machine")
if ($updatedValue -eq $VariableValue) {
    Write-Host "Variable $VariableName = $VariableValue added/updated succesfully. Don't forget to restart your IDE!"
} else {
    Write-Error "Adding/updating variable $VariableName failed."
}