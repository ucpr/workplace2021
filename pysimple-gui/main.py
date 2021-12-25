import PySimpleGUIQt as sg

sg.theme('SystemDefault')

value = sg.popup_get_file('get file multi', multiple_files=True)
print(value)
