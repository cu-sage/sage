FROM golang
ADD bin/sage /
CMD ["/sage"]
