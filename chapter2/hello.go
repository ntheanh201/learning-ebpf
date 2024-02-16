package main

import "github.com/iovisor/gobpf/bcc"

const source string = `
int hello(void *ctx) {
	bpf_trace_printk("Hello World!");
	return 0;
}`

func main() {
	m := bcc.NewModule(source, []string{})
	defer m.Close()
}
