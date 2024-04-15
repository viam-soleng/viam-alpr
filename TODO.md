# Viam-ALPR TODO

Outstanding action items.

### Todo

- [ ] Update binary file name from viam-alpr.AppImage to viam-alpr for the Linux appimage and the mac binary
- [ ] Add additional openalpr configuration parameters -> see config files in the runtime_dir (Pattern is missing but not exposed through the CGO api yet)
- [ ] Automate build process
  - [ ] Add build/bundling command for Mac: `dylibbundler -od -b -x viam-alpr`
  - [ ] Optimize / prepare AppDir and AppImageBuild.yml so it can be used in an easier way with less complicated scripting

### In Progress

- [ ] ...  

### Done ✓

- [x] Implemented getDetectionsFromCamera()
- [x] BBox display incorrect
