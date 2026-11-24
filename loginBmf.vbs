user  = GetEnvValue("C:\senhas\.env", "GP_USER")
senha = GetEnvValue("C:\senhas\.env", "SAP_GP")
entry = GetEnvValue("C:\senhas\.env", "ENTRY_GP")


If Not IsObject(application) Then
   Set SapGuiAuto  = GetObject("SAPGUI")
   Set application = SapGuiAuto.GetScriptingEngine
End If
If Not IsObject(connection) Then
   Set connection = application.OpenConnection(entry, True)
End If
If Not IsObject(session) Then
  Set session    = connection.Children(0)
End If
If IsObject(WScript) Then
   WScript.ConnectObject session,     "on"
   WScript.ConnectObject application, "on"
End If
session.findById("wnd[0]").sendVKey 0
session.findById("wnd[0]").maximize
session.findById("wnd[0]/usr/txtRSYST-BNAME").text = user
session.findById("wnd[0]/usr/pwdRSYST-BCODE").text = senha
session.findById("wnd[0]").sendVKey 0


Function GetEnvValue(envFile, key)

    Set fso = CreateObject("Scripting.FileSystemObject")
    Set file = fso.OpenTextFile(envFile, 1)

    Do Until file.AtEndOfStream

        line = Trim(file.ReadLine)

        If Left(line, Len(key) + 1) = key & "=" Then
            GetEnvValue = Mid(line, Len(key) + 2)
            file.Close
            Exit Function
        End If

    Loop

    file.Close
    GetEnvValue = ""

End Function