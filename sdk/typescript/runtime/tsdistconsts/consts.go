package tsdistconsts

const (
	DefaultNodeVersion  = "22.11.0" // LTS version, JOD (https://nodejs.org/en/about/previous-releases)
	nodeImageDigest     = "sha256:b64ced2e7cd0a4816699fe308ce6e8a08ccba463c757c00c14cd372e3d2c763e"
	DefaultNodeImageRef = "node:" + DefaultNodeVersion + "-alpine@" + nodeImageDigest

	DefaultBunVersion  = "1.1.38"
	bunImageDigest     = "sha256:c1cc397e0be452c54f37cbcdfaa747eff93c993723af7d91658764d0fdfe5873"
	DefaultBunImageRef = "oven/bun:" + DefaultBunVersion + "-alpine@" + bunImageDigest

	DefaultDenoVersion  = "2.4.0"
	denoImageDigest     = "sha256:fcf215ca621c2834157dcb8a8c8c48b64d273b542b4fc8baee1b5c6de50b326c"
	DefaultDenoImageRef = "denoland/deno:alpine-" + DefaultDenoVersion + "@" + denoImageDigest
)
