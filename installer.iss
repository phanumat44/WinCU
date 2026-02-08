; Inno Setup Script for Window Clean-up CLI (wincu)
; Requires Inno Setup to compile: https://jrsoftware.org/isdl.php

[Setup]
AppName=WinCU
AppVersion=1.0.1
DefaultDirName={sd}\wincu
DefaultGroupName=WinCU
UninstallDisplayIcon={app}\wincu.exe
Compression=lzma2
SolidCompression=yes
OutputDir=.
OutputBaseFilename=wincu_installer
 SetupIconFile=assets\wincu.ico  
 WizardImageFile=assets\sideicon.bmp 

[Files]
Source: "wincu.exe"; DestDir: "{app}"; Flags: ignoreversion
; Source: "User-manual.md"; DestDir: "{app}"; Flags: isreadme
Source: "Readme.txt"; DestDir: "{app}"; Flags: isreadme

[Icons]
Name: "{group}\WinCU"; Filename: "{app}\wincu.exe"
Name: "{group}\Uninstall WinCU"; Filename: "{uninstallexe}"

[Tasks]
Name: "modifypath"; Description: "Add application directory to your environmental path"; GroupDescription: "Additional icons:";

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; Tasks: modifypath; Check: NeedsAddPath(ExpandConstant('{app}'))

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_LOCAL_MACHINE,
    'SYSTEM\CurrentControlSet\Control\Session Manager\Environment',
    'Path', OrigPath)
  then begin
    Result := True;
    exit;
  end;
  // look for the path with leading and trailing semicolon
  // Pos() returns 0 if not found
  Result := Pos(';' + Param + ';', ';' + OrigPath + ';') = 0;
end;
