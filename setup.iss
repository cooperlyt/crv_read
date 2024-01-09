; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#define MyAppName "身份证读卡器辅助程序"
#define MyAppVersion "1.5"
#define MyAppPublisher "沈阳菲普科技有限公司"
#define MyAppURL "https://www.example.com/"
#define MyAppExeName "card-reading.exe"
#define MyAppAssocName MyAppName + ""
#define MyAppAssocExt ""
#define MyAppAssocKey StringChange(MyAppAssocName, " ", "") + MyAppAssocExt

[Setup]
; NOTE: The value of AppId uniquely identifies this application. Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{E37903B3-D990-4F57-8EDA-52A584C47195}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\card-reader
ChangesAssociations=yes
DisableProgramGroupPage=yes
; Uncomment the following line to run in non administrative install mode (install for current user only.)
;PrivilegesRequired=lowest
OutputBaseFilename=mysetup
Compression=lzma
SolidCompression=yes
WizardStyle=modern
PrivilegesRequired=admin

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "C:\Users\cooper\Documents\go\card-reading\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\card-reading.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\DLL_File.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\license.dat"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\sdtapi.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\Termb.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:\Users\cooper\Documents\go\card-reading\WltRS.dll"; DestDir: "{app}"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Registry]
Root: HKA; Subkey: "Software\Classes\{#MyAppAssocExt}\OpenWithProgids"; ValueType: string; ValueName: "{#MyAppAssocKey}"; ValueData: ""; Flags: uninsdeletevalue
Root: HKA; Subkey: "Software\Classes\{#MyAppAssocKey}"; ValueType: string; ValueName: ""; ValueData: "{#MyAppAssocName}"; Flags: uninsdeletekey
Root: HKA; Subkey: "Software\Classes\{#MyAppAssocKey}\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: "{app}\{#MyAppExeName},0"
Root: HKA; Subkey: "Software\Classes\{#MyAppAssocKey}\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\{#MyAppExeName}"" ""%1"""
Root: HKA; Subkey: "Software\Classes\Applications\{#MyAppExeName}\SupportedTypes"; ValueType: string; ValueName: ".myp"; ValueData: ""
Root: HKLM; Subkey: "Software\Microsoft\Windows\CurrentVersion\Run"; ValueType: string; ValueName: "CardReader"; ValueData: "{app}\{#MyAppExeName}"
Root: HKCR; Subkey: "cardreader"; ValueType: string; ValueName: ""; ValueData: "cardreaderProtocol"; Flags: uninsdeletekey
Root: HKCR; SubKey: "cardreader"; ValueName: "URL Protocol"; ValueData: "{app}\{#MyAppExeName}"; ValueType: string; Flags: createvalueifdoesntexist uninsdeletekey
Root: HKCR; Subkey: "cardreader\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\{#MyAppExeName}"" ""%1"""; Flags: uninsdeletekey
Root: HKCR; SubKey: "cardreader\DefaultIcon"; ValueData: "{app}\{#MyAppExeName}"; ValueType: string; Flags: createvalueifdoesntexist uninsdeletekey;
Root: HKCR; SubKey: "cardreader\shell\open\command"; ValueData: "{app}\{#MyAppExeName} ""%1"""; Flags: createvalueifdoesntexist uninsdeletekey; ValueType: string;


[Icons]
Name: "{autoprograms}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall runascurrentuser skipifsilent

