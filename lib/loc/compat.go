package loc

// IsLegacyRuntimePackage returns true if the specified package envelope describes
// a legacy runtime package
// TODO: remove when there's no more need to take legacy runtime packages into account
func IsLegacyRuntimePackage(loc Locator) bool {
	if loc.Repository != LegacyPlanetMaster.Repository {
		// Skip runtime package with a non-standard repository
		return false
	}
	switch loc.Name {
	case LegacyPlanetMaster.Name, LegacyPlanetNode.Name:
		return true
	default:
		return false
	}
}

// IntermediateRuntimePackage identifies the intermediate runtime package
// used to bridge updates from 5.0.
// Note, this version needs to be in sync with the version in the Makefile
var IntermediateRuntimePackage = MustParseLocator("gravitational.io/planet:5.2.23-11109")
