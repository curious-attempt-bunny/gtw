# GTW - Guess the Word

This module implements a word guessing game. The game is similar
to a game hosted on the web site of a major metropolitan newpaper.

## To run

```bash
go run cmd/cli/main.go cmd/cli/wikipedia-top-five-letter-frequency.corpus
```

Example output:
```
New goal word selected
guess> which
       ##### (0 letters in the correct place)
Guess which filtered corpus from 3000 words down to 2351 words
guess> first
       ####* (0 letters in the correct place)
Guess first filtered corpus from 2351 words down to 428 words
guess> after
       ##+*# (1 letters in the correct place)
Guess after filtered corpus from 428 words down to 44 words
guess> noted
       ##+*# (1 letters in the correct place)
Guess noted filtered corpus from 44 words down to 33 words
guess> metal
       #++## (2 letters in the correct place)
Guess metal filtered corpus from 33 words down to 10 words
guess> betty
       #++++ (4 letters in the correct place)
Guess betty filtered corpus from 10 words down to 3 words
guess> petty
       #++++ (4 letters in the correct place)
Guess petty filtered corpus from 3 words down to 2 words
guess> getty
       #++++ (4 letters in the correct place)
Guess getty filtered corpus from 2 words down to 1 words
guess> jetty

Success!
```
