/*
Package config defines the structure of the configuration
for the `dev` tool. I expect it to be YAML, but am providing
support for JSON via the struct tags to keep the options open.

There's a global configuration instance which should be populated
by the controller and consumed from all other parts of the view
or controller layers.
*/
package config
