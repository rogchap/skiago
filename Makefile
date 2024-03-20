build-darwin_arm64:
	cd external/skia; \
	rm -rf out/mac-apple; \
	bin/gn gen out/mac-apple --args='\
	target_cpu="arm64" \
	is_official_build=true \
	skia_use_harfbuzz=true \
	skia_pdf_subset_harfbuzz=true \
	skia_use_icu=false \
	skia_use_system_harfbuzz=false \
	skia_use_system_libjpeg_turbo=false \
	skia_use_system_libpng=false \
	skia_use_system_libwebp=false \
	skia_use_system_expat=false \
	skia_use_system_zlib=false \
	'; \
	../depot_tools/ninja -C out/mac-apple

copy:
	cp external/skia/out/mac-apple/libskia.a lib/darwin_arm64/libskia.a
	find external/skia/include/core -name \*.h -exec cp {} lib/include/core \;
	find external/skia/include/private -name \*.h -exec cp {} lib/include/private \;
	find external/skia/include/private/base -name \*.h -exec cp {} lib/include/private/base \;
	find external/skia/include/config -name \*.h -exec cp {} lib/include/config \;
	find external/skia/include/docs -name \*.h -exec cp {} lib/include/docs \;
	find external/skia/include/ports -name \*.h -exec cp {} lib/include/ports \;
	cp external/skia/src/base/SkTime.h lib/src/base/SkTime.h


sync:
	cd external/skia; python3 tools/git-sync-deps; bin/fetch-ninja

