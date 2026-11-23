
On Error Resume Next  ' Ignora erros (para tratar manualmente)

Set SapGuiAuto  = GetObject("SAPGUI")
If Err.Number <> 0 Then
    MsgBox "SAP GUI não está em execução.", vbExclamation
    WScript.Quit
End If

Set application = SapGuiAuto.GetScriptingEngine
If application.Children.Count = 0 Then
    MsgBox "Nenhuma conexão SAP aberta encontrada.", vbExclamation
    WScript.Quit
End If

Set connection = application.Children(0)
If connection.Children.Count = 0 Then
    MsgBox "Nenhuma sessão SAP ativa na conexão.", vbExclamation
    WScript.Quit
End If

Set session = connection.Children(0)


session.findById("wnd[0]").maximize
session.findById("wnd[0]/tbar[0]/okcd").text = "/nse16n"
session.findById("wnd[0]").sendVKey 0

