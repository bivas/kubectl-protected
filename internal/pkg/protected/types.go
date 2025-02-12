package protected

const (
	DefaultProtectedFilePath = "$HOME/.kube/protected.yaml"
)

type Options struct {
	SilenceOnProtected bool
	ProtectedFilePath  string
}
