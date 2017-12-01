unit SSSplash; x

interface

uses
  Windows, Messages, SysUtils, Classes, Graphics, Controls, Forms, Dialogs;

type

  TSemaphoreNames = (snUseExeName, snUseAppTitle);

  TSSVersionDetails = class(TPersistent)
  private
    FShowVersion: boolean;
    FVersion: string;
  published
    property ShowVersion: boolean read FShowVersion write FShowVersion default True;
    property Version: string read FVersion write FVersion;
  end;

  TSSSplash = class(TComponent)
  private
    FAppTitle: string;
    FAutoRun: boolean;
    FMaxInstances: integer;
    FMinDisplayMSecs: integer;
    FVersion: TSSVersionDetails;
    FSemaphoreName: TSemaphoreNames;
    SemHandle: THandle;
    SemAquired: Boolean;
    function CheckInstance: Boolean;
    procedure AbandonApplication;
    procedure SetMaxInstances(Value: integer);
  public
    constructor Create(aOwner: TComponent); override;
    destructor Destroy; override;
    procedure Loaded; override;
  published
    property AppTitle: string read FAppTitle write FAppTitle;
    property AutoRun: boolean read FAutoRun write FAutoRun default True;
    property MaxInstances: integer read FMaxInstances write SetMaxInstances default 1;
    property MinDisplayMSecs: integer read FMinDisplayMSecs write FMinDisplayMSecs default 3000;
    property SemaphoreName: TSemaphoreNames read FSemaphoreName write FSemaphoreName default snUseExeName;
    property Version: TSSVersionDetails read FVersion write FVersion;
  end;

procedure Register;

implementation

uses frmSSSplashUnit;

{---------------}
procedure Register;
begin
  RegisterComponents('Synstar', [TSSSplash]);
end;

{---------------}
constructor TSSSplash.Create(aOwner: TComponent);
begin
  inherited Create(aOwner);
  FAutoRun := True;
  FMaxInstances := 1;
  FMinDisplayMSecs := 3000;
  FVersion := TSSVersionDetails.Create;
  FVersion.ShowVersion := True;
end;

{---------------}
destructor TSSSplash.Destroy;
begin
  FVersion.Free;
  if SemAquired then ReleaseSemaphore(SemHandle, 1, nil);
  inherited Destroy;
end;

{---------------}
procedure TSSSplash.SetMaxInstances(Value: integer);
begin
  if Value < -1 then
    FMaxInstances := -1
  else
    if Value = 0 then
      FMaxInstances := 1
    else
      FMaxInstances := Value;   
end;

{---------------}
procedure TSSSplash.Loaded;
begin
  inherited Loaded;
  if not(csDesigning in ComponentState) and AutoRun and ((FMaxInstances = -1) or CheckInstance) then
    with TfrmSSSplash.Create(Self) do
    begin
      Show;
      Update;
    end;
end;

{---------------}
function TSSSplash.CheckInstance: boolean;
var
  AppHandle: THandle;
  SemName: string;
begin
  {Set the name to give to the semaphore. The option of using the app title is required by xolis}
  if (FSemaphoreName = snUseExeName) or (Trim(Application.Title) = '') then
  begin
    SemName := ExtractFileName(Application.ExeName);
    if Pos('.', SemName) > 0 then SemName := Copy(SemName, 0, Pos('.', SemName)-1);
  end
    else
      SemName := Trim(Application.Title);
  {If only permitting a single instance of the app then hide the application handle}
  {in the semaphore as its 'count' value so that duplicate instances can retrieve  }
  {it and use it to bring the existing application instance to the fore            }
  if FMaxInstances = 1 then
    SemHandle := CreateSemaphore(nil, Application.Handle+2, Application.Handle+2, PChar(SemName))
  else
    SemHandle := CreateSemaphore(nil, MaxInstances, MaxInstances, PChar(SemName));
  {If an attempt is being made to start multiple instances of a single instance}
  {app then bring the existing instance to the fore. The existing app handle is}
  {saved in the semaphore as its 'count' value                                 }
  if (GetLastError = ERROR_ALREADY_EXISTS) and (FMaxInstances = 1) then
  begin
    WaitForSingleObject(SemHandle, 0);
    ReleaseSemaphore(SemHandle, 1, @AppHandle);
    SetForegroundWindow(AppHandle);
    ShowWindow(AppHandle, SW_SHOW);
    AbandonApplication;
  end
    else
      {Else if we have a semaphore handle then decrement its count}
      if (SemHandle <> 0) and (WaitForSingleObject(SemHandle, 0) = WAIT_OBJECT_0) then
        SemAquired := True
      else
        AbandonApplication;
  Result := SemAquired;      
end;

{---------------}
procedure TSSSplash.AbandonApplication;
begin
  Application.ShowMainForm := False;
  Application.Terminate;
  Application.ProcessMessages;
  Application.Terminate;
end;

end.
