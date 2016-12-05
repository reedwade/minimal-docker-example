
# Tiny Docker Examples

`example1-bad` shows what you might first attempt. It doesn't work because it requires more libraries than you've provided for. It's why you normally do start from a larger base docker image.

`example2-good` shows how you can work with nearly nothing. The executable has all the libraries it needs statically linked.

`example3-awesome` shows an example of a little standalone web app with embedded content. From there it's easy to see how you could expand that into a small service that did more interesting things.

# See also

https://docs.docker.com/engine/userguide/eng-image/baseimages/
