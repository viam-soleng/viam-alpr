# VIAMALPR Build Instructions

## Mac / OSX

OSX uses @rpath to reference dynamic libraries. After you have built the binary, you therefore have to set the @rpath with the following command. Otherwise the module will complain about missing libraries!

`install_name_tool -add_rpath @executable_path/./libs viamalpr`

You also have to make sure you copy the required libraries in the ./libs folder!

