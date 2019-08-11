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
	PLATFORM="linux" ./build.sh run_linux

clean:

.PHONY:HELP
