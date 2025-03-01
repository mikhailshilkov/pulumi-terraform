package tfgen

const pyUtilitiesFile = `
import os
import pkg_resources

from semver import VersionInfo as SemverVersion
from parver import Version as PEP440Version

def get_env(*args):
    for v in args:
        value = os.getenv(v)
        if value is not None:
            return value
    return None

def get_env_bool(*args):
    str = get_env(*args)
    if str is not None:
        # NOTE: these values are taken from https://golang.org/src/strconv/atob.go?s=351:391#L1, which is what
        # Terraform uses internally when parsing boolean values.
        if str in ["1", "t", "T", "true", "TRUE", "True"]:
            return True
        if str in ["0", "f", "F", "false", "FALSE", "False"]:
            return False
    return None

def get_env_int(*args):
    str = get_env(*args)
    if str is not None:
        try:
            return int(str)
        except:
            return None
    return None

def get_env_float(*args):
    str = get_env(*args)
    if str is not None:
        try:
            return float(str)
        except:
            return None
    return None

def get_version():
    # __name__ is set to the fully-qualified name of the current module, In our case, it will be
    # <some module>.utilities. <some module> is the module we want to query the version for.
    root_package, *rest = __name__.split('.')

    # pkg_resources uses setuptools to inspect the set of installed packages. We use it here to ask
    # for the currently installed version of the root package (i.e. us) and get its version.

    # Unfortunately, PEP440 and semver differ slightly in incompatible ways. The Pulumi engine expects
    # to receive a valid semver string when receiving requests from the language host, so it's our
    # responsibility as the library to convert our own PEP440 version into a valid semver string.

    pep440_version_string = pkg_resources.require(root_package)[0].version
    pep440_version = PEP440Version.parse(pep440_version_string)
    (major, minor, patch) = pep440_version.release
    prerelease = None
    if pep440_version.dev is not None:
        prerelease = f"dev.{pep440_version.dev}"

    # The only significant difference between PEP440 and semver as it pertains to us is that PEP440 has explicit support
    # for dev builds, while semver encodes them as "prerelease" versions. In order to bridge between the two, we convert
    # our dev build version into a prerelease tag. This matches what all of our other packages do when constructing
    # their own semver string.
    semver_version = SemverVersion(major=major, minor=minor, patch=patch, prerelease=prerelease)
    return str(semver_version)
`
