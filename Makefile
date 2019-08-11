ifeq ($(OS),Windows_NT)
  PLATFORM="Windows"
else
  ifeq ($(shell uname),Darwin)
    PLATFORM="MacOS"
  else
    PLATFORM="Linux"
  endif
endif

HELP:
	@echo "make android or make linux"

android:
	PLATFORM="android" ./build.sh android

linux:
	PLATFORM="linux" ./build.sh linux

clean:
	rm -rf sysup ipfs sysup_linux ipfs_linux ota.zip reboot.flag

.PHONY:HELP
