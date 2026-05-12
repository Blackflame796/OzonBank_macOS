#!/bin/bash

APP_NAME="OzonBank"
BUNDLE="${APP_NAME}.app"
MAC_OS_DIR="${BUNDLE}/Contents/MacOS"
RESOURCES_DIR="${BUNDLE}/Contents/Resources"

echo "Building ${APP_NAME}..."
mkdir -p "${MAC_OS_DIR}"
mkdir -p "${RESOURCES_DIR}"

# Build the Go app with optimizations disabled to prevent darwinkit crash
go build -gcflags "all=-N" -o "${MAC_OS_DIR}/${APP_NAME}" main.go

# Create icon
ICON_PNG="Logo.png"
ICONSET_DIR="MyIcon.iconset"
if [ -f "${ICON_PNG}" ]; then
	echo "Generating icns from ${ICON_PNG}..."
	mkdir -p "${ICONSET_DIR}"
	sips -z 16 16     "${ICON_PNG}" --out "${ICONSET_DIR}/icon_16x16.png" > /dev/null
	sips -z 32 32     "${ICON_PNG}" --out "${ICONSET_DIR}/icon_16x16@2x.png" > /dev/null
	sips -z 32 32     "${ICON_PNG}" --out "${ICONSET_DIR}/icon_32x32.png" > /dev/null
	sips -z 64 64     "${ICON_PNG}" --out "${ICONSET_DIR}/icon_32x32@2x.png" > /dev/null
	sips -z 128 128   "${ICON_PNG}" --out "${ICONSET_DIR}/icon_128x128.png" > /dev/null
	sips -z 256 256   "${ICON_PNG}" --out "${ICONSET_DIR}/icon_128x128@2x.png" > /dev/null
	sips -z 256 256   "${ICON_PNG}" --out "${ICONSET_DIR}/icon_256x256.png" > /dev/null
	sips -z 512 512   "${ICON_PNG}" --out "${ICONSET_DIR}/icon_256x256@2x.png" > /dev/null
	sips -z 512 512   "${ICON_PNG}" --out "${ICONSET_DIR}/icon_512x512.png" > /dev/null
	sips -z 1024 1024 "${ICON_PNG}" --out "${ICONSET_DIR}/icon_512x512@2x.png" > /dev/null
	iconutil -c icns "${ICONSET_DIR}" -o "${RESOURCES_DIR}/AppIcon.icns"
	rm -R "${ICONSET_DIR}"
fi

# Create Info.plist
cat <<EOF > "${BUNDLE}/Contents/Info.plist"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>${APP_NAME}</string>
	<key>CFBundleIdentifier</key>
	<string>com.ozon.bankapp</string>
	<key>CFBundleName</key>
	<string>${APP_NAME}</string>
	<key>CFBundleVersion</key>
	<string>1.0</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>LSMinimumSystemVersion</key>
	<string>10.13</string>
	<key>CFBundleIconFile</key>
	<string>AppIcon</string>
</dict>
</plist>
EOF

echo "App Bundle created at ${BUNDLE}"
