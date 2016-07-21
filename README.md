# tweetdelete

[![Software
License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/mannkind/tweetdelete/blob/master/LICENSE.md)
[![Travis CI](https://img.shields.io/travis/mannkind/tweetdelete/master.svg?style=flat-square)](https://travis-ci.org/mannkind/tweetdelete)
[![Coverage Status](http://codecov.io/github/mannkind/tweetdelete/coverage.svg?branch=master)](http://codecov.io/github/mannkind/tweetdelete?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mannkind/tweetdelete)](https://goreportcard.com/report/github.com/mannkind/tweetdelete)

Deleting Tweets Automatically

# Installation

* go get github.com/mannkind/tweetdelete
* go intall github.com/mannkind/tweetdelete
* tweetdelete -c */the/path/to/config.yaml*

# Configuration

Configuration happens in the config.yaml file. A full example might look this:

```
consumer_key: "YOUR CONSUMER KEY"
consumer_secret: "YOUR CONSUMER SECRET"
oauth_token: "YOUR OAUTH TOKEN"
oauth_token_secret: "YOUR OAUTH TOKEN SECRET"
timeline_count: 50
max_age: 72
save:
    - "123456789012345678"
```
