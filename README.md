#general
DIS is YASB.  This time it's docker-in-dockerized.  Dis gives you the ability to realize your wildest slack dreams just by installing new containers and running them (via dis in slack).

#build
```
make build
make docker
```

#usage
You need a slack token, so get one of those.
You may need to set the bot's name in the src.

```
docker run --privileged -e SLACK_TOKEN='' dis
```

#examples
> Want to query wolframalpha in slack?
```
@dis docker pull -t dancompton/wolfram:latest
@dis docker run -e WOLFRAM_API_KEY=YOURWOLFRAMKEY dancompton/wolfram cat images
```

#TODO
Everything's just a command right now.  Might be better to limit what you can do.  Might be good to display images by themeselves. etc.

