init:
	go build
	./weekly-report generate -t=template/cn.md --token=b93f8ca80e43c1e3553a806f347ac1d81760979c --owner=kubernetes --repo=kubernetes > weekly.md
