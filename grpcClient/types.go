package grpc


// holds data of each command
// package private
// each command has elements such as cmd name, path where the command binary will be dropped.
// it'll also have access (rwx) permissions. rwx for root user for now to keep it simple
type cmd struct {
	name string    `json:"name"`
	absPath string `json:"absPath"`
}
