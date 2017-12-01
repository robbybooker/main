unit frmSSSplashUnit;

interface

uses
  Windows, Messages, SysUtils, Classes, Graphics, Controls, Forms, Dialogs,
  StdCtrls, ExtCtrls, lmdctrl, lmdstdcS, lmdcompo, lmdclass, lmdformA,
  lmdcctrl, lmdextcS, lmdnonvS, ComCtrls, jpeg;

type
  TfrmSSSplash = class(TForm)
    Image1: TImage;
    timDisplay: TTimer;
    lblVersion: TLMDSimpleLabel;
    LMDPanelFill: TLMDPanelFill;
    lblAppTitle: TLMDSimpleLabel;
    timCheckForms: TTimer;
    procedure timDisplayTimer(Sender: TObject);
    procedure FormShow(Sender: TObject);
    procedure timCheckFormsTimer(Sender: TObject);
    procedure FormClose(Sender: TObject; var Action: TCloseAction);
  private
    StartDisplayTimeMSecs: comp;
  end;

implementation

{$R *.DFM}

uses SSSplash;

{---------------}
procedure TfrmSSSplash.FormShow(Sender: TObject);
var
  i: integer;
begin
  {Set the application title on the splash screen}
  if Trim((Owner as TSSSplash).AppTitle) = '' then
    lblAppTitle.Caption := Trim(Application.Title)
  else
    lblAppTitle.Caption := Trim((Owner as TSSSplash).AppTitle);

  //Set the font size so that it's appropriate(ish) to the splash form width
  if (Length(lblAppTitle.Caption) > 15) then
    lblAppTitle.Font.Size := 15
  else
    lblAppTitle.Font.Size := 25;

  {Set the version details on the splash screen}
  if (Owner as TSSSplash).Version.ShowVersion then
    lblVersion.Caption := 'Version  ' + Trim((Owner as TSSSplash).Version.Version)
  else
    lblVersion.Hide;
  {Ensure that the hour glass cursor appears all over the splash screen}
  for i := 0 to ControlCount - 1 do
    Controls[i].Cursor := crHourGlass;
  StartDisplayTimeMSecs := TimeStampToMSecs(DateTimeToTimeStamp(Now));
end;

{---------------}
procedure TfrmSSSplash.FormClose(Sender: TObject; var Action: TCloseAction);
begin
  {If the AIS login window exists then DO NOT free the splash screen!!! Reason: The AIS }
  {login window behaves like (but is actually not) a modal child of whatever was the    }
  {active form when the login form was created (but is not available thru TScreen.Forms.}
  {Therefore when the splash screen closes the AIS login window is also closed which is }
  {regarded as a login abortion at which point the user gets all sorts of errors        }      
  if FindWindow(nil, 'Advantage Internet Client Authentication') = 0 then
    Action := caFree
  else
    {If the AIS login window exists then DO NOT free the splash screen. No big deal} 
    Action := caHide;
end;

{Because the application is so pre-occupied with starting up, timCheckForms will not  }
{not start 'ticking' until the first form is displayed (i.e. the app becomes idle) and}
{so it is in here that we want to establish how much longer the splash screen is to be}
{displayed for                                                                        }
procedure TfrmSSSplash.timCheckFormsTimer(Sender: TObject);
var
  TimeSpanMSecs: variant;
begin
  timCheckForms.Enabled := False;
  {Calculate the time elapsed since the splash screen was displayed}
  TimeSpanMSecs := TimeStampToMSecs(DateTimeToTimeStamp(Now)) - StartDisplayTimeMSecs;
  {If not displayed for at least n seconds then set display time remaining}
  if TimeSpanMSecs < (Owner as TSSSplash).MinDisplayMSecs then
    timDisplay.Interval := (Owner as TSSSplash).MinDisplayMSecs - TimeSpanMSecs
  else
    timDisplay.Interval := 1;
end;

{---------------}
procedure TfrmSSSplash.timDisplayTimer(Sender: TObject);
begin
  timDisplay.Enabled := False;
  Close;
end;

end.
