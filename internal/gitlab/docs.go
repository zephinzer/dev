/*
Package gitlab is a wrapper around the gitlab package in the `./pkg` directory.
The difference between this package and that is this package implements a
processing layer so that the view layer can interact with interfaces instead of
the raw API outputs.

This package also implements other app-specific things such as database table
initialisations for custom implemented logic.
*/
package gitlab
