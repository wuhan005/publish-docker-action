# Publish Docker Action

Publish Docker Action is used to build, tag and publish docker image to your docker registry.

## Usage

This simple example will use `Dockerfile` in your workspace to build image, attach the `latest`
tag and push to docker default registry (docker.io). Repository name is your GitHub repository
name by default.

```yaml
- uses: wuhan005/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

Use `file` and `path` arguments to set docker build file or build context if they are not placed
in the default workspace direcotry.

### Set up registry and repository name

Registry and repository name can be changed with `registry` and `repository` arguments. For example:

```yaml
- uses: wuhan005/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
```

#### Tag Format

When you set `auto_tag` to `true`, you can customize the image's tag.

Use `tag_format` to format the tag. Default is `"%TIMESTAMP%"`

Format signs:
```
%TIMESTAMP%     Timestamp
%YYYY%          Year
%MM%            Month
%DD%            Day
%H%             Hour
%m%             Minute
%s%             Second
```

```yaml
- uses: wuhan005/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: wuhan005/publish-docker-action
    auto_tag: true
    tag_format: "foo%YYYY%_%MM%"
```

And an image with the tag `foo2019_9` will be published.

### Cache

Cache image can be used to build image. Just provide `with.cache` argument.
