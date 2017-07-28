
## how to make release
- https://github.com/goreleaser/goreleaser 
- `curl -sL https://git.io/goreleaser > goreleaser`
- `chmod +x goreleaser`
- `export GITHUB_TOKEN=YOUR_TOKEN`
- `git add .`
- `git commit -m "message"`
- `git tag -a v0.0.1 -m "v0.0.1"`
- `rm -rf dist` - help write use 'use the --rm-dist flag' but it not work for me
- `goreleaser`
