FROM scratch

WORKDIR /Volumes/2Tmac/github/mine/blogie
COPY . /Volumes/2Tmac/github/mine/blogie

EXPOSE 8080
CMD ["./blogie"]
