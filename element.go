package gopq

type Element interface {
	CompareTo(interface{}) int
}