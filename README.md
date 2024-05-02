# DateVer [![Godoc](https://godoc.org/github.com/bschaatsbergen/datever?status.svg)](https://godoc.org/github.com/bschaatsbergen/datever)

DateVer is a library to define, compare and validate date ranges while adhering to [semantic versioning](https://semver.org/). DateVer versions consist of a year, month, day, and an optional patch. The patch can be used to indicate pre-release versions, such as alpha, beta, or release candidates. DateVer helps to provide a clear chronological view of your software's evolution.

## Example DateVer versions

* v2024.1.1
* v2024.1.2-1
* v2024.2.1-rc1
* v2024.2.1-alpha001

## Why DateVer?
Many projects use semantic versioning to manage their releases. However, the version number does not provide any information about the release date. DateVer combines the release date with semantic versioning to provide more context about the release.

## Usage

This section details how to use `datever`.

### Creating a DateVer version

```go
version := &datever.Version{Year: 2024, Month: 2, Day: 1, Patch: "alpha001"}
fmt.Println(version.String()) // Output: v2024.2.1-alpha001
```

### Parsing a DateVer version

```go
version, err := datever.ParseVersion("v2024.1.1-rc1")
```

### Using a Version struct
Once you have a `datever.Version` struct, you can access its components (year, month, day, patch) and use the provided methods:

* `String()`: Returns the string representation of the version.
* `Compare(other *Version)`: Compares the current version with another datever.Version and returns -1, 0, or 1 based on comparison.
* `IsValid()`: Checks if the date components of the version are valid (year > 0, month within 1-12, day within 1-31).
