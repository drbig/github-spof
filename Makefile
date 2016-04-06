.PHONY: release

clean:
	rm -rf github_spof_*

compile:
	env GOOS=darwin GOARCH=386 go build
	mv github-spof github_spof_$(VERSION)_darwin_386
	env GOOS=linux GOARCH=arm GOARM=7 go build
	mv github-spof github_spof_$(VERSION)_linux_arm
	env GOOS=linux GOARCH=386 go build
	mv github-spof github_spof_$(VERSION)_linux_386

release: clean compile
