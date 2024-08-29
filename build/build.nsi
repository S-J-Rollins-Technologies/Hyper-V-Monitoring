!include MUI2.nsh

Name "Hyper-V Monitoring API"

Outfile "hvmon_installer_win-x64.exe"

InstallDir $PROGRAMFILES64\hvmon

!define MUI_ABORTWARNING

!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

!insertmacro MUI_UNPAGE_INSTFILES

!insertmacro MUI_LANGUAGE "English"

Section "Check Elevation"

    UserInfo::GetAccountType
    Pop $0
    StrCmp $0 "Admin" +3
    MessageBox MB_OK "Please re-launch the installer as an administrator: $0"
    Quit

SectionEnd

Section "Install"

    SetOutPath $INSTDIR

    File hvmon.exe

    WriteUninstaller $INSTDIR\uninstall.exe

    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\HyperVMonitoringAPI" \
            "DisplayName" "Hyper-V Monitoring API"

    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\HyperVMonitoringAPI" \
            "UninstallString" "$\"$INSTDIR\uninstall.exe$\""

    Exec '"$INSTDIR\hvmon.exe" install'
    Sleep 3000
    Exec '"$INSTDIR\hvmon.exe" start'
SectionEnd

Section "Uninstall"

    Exec '"$INSTDIR\hvmon.exe" stop'
    Exec '"$INSTDIR\hvmon.exe" uninstall'
    Sleep 5000
    DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\HyperVMonitoringAPI"
    RMDir /r $INSTDIR

SectionEnd