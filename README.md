# 📈 Statusbar for DWM

![GitHub](https://img.shields.io/github/license/mamaart/dwm-status)

This is actually two programs in one. Its a service that is meant to be run in the background, and it is a client as well to add or delete todo items from the service.

The service can show some standard information about the computer, but it also has an infinite rolling text field which is used to show todo items.

## ⚡️Some features

- [x] Show ip of default route
- [x] Volume from pulseaudio
- [x] Brightness
- [x] Date and time
- [x] Ask LLM AI

## 🔧 TODO's

- [ ] Try to not use library for pulseaudio (make it yourself)
- [ ] Add command to change speed
- [ ] Test in on something else than DWM
- [ ] window width should be UTF8 or 16 strings not bytes 
- [ ] add stop command for AI stream (maybe)
- [ ] Take AI out of statusbar and make a stream thingy that accepts a pipe
- [ ] Write tests (find out what it the best way to write tests when you are interacting with the os)

