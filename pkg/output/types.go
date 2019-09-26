package output

// Groups is a collection of groups.
type Groups []*Group

// Append simply appends a Group if it doesn't exist.
func (g *Groups) Append(val Group) *Group {
	for _, group := range *g {
		if group.Name == val.Name {
			return group
		}
	}

	*g = append(*g, &val)
	return &val
}

// Len implements sort.Interface
func (g Groups) Len() int {
	return len(g)
}

// Less implements sort.Interface
func (g Groups) Less(i, j int) bool {
	return g[i].Name < g[j].Name
}

// Swap implements sort.Interface
func (g Groups) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// Group defines a single group record.
type Group struct {
	Name  string
	Hosts Hosts
}

// Hosts is a collection of hosts.
type Hosts []*Host

// Append simply appends a Host if it doesn't exist.
func (h *Hosts) Append(val Host) *Host {
	for _, host := range *h {
		if host.Name == val.Name {
			return host
		}
	}

	*h = append(*h, &val)
	return &val
}

// Len implements sort.Interface
func (h Hosts) Len() int {
	return len(h)
}

// Less implements sort.Interface
func (h Hosts) Less(i, j int) bool {
	return h[i].Name < h[j].Name
}

// Swap implements sort.Interface
func (h Hosts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Host defines a single host record.
type Host struct {
	Name    string
	User    string
	Address string
	Port    string
}

// DefaultTemplate defines the standard hosts config.
var DefaultTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE hosts SYSTEM "file:///usr/share/xml/schema/apt-dater/hosts.dtd">
<hosts>
	<default ssh-user="root"/>

{{- range .Groups }}
	<group name="{{ .Name }}">
{{- range .Hosts }}
		<host name="{{ .Name }}" ssh-user="{{ .User }}" ssh-host="{{ .Address }}" ssh-port="{{ .Port }}" />
{{- end }}
	</group>
{{- end }}
</hosts>`
