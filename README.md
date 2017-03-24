# Stamper

Stamper is a simple web service listening for incoming web hook requests from
GitHub for opened issues/pull requests, and comments on issues/pull requests.

The idea is that Stamper will assign labels to the issue or pull request based
on whether the issue/pull request body or issue/pull request comment contains a
certain string.
