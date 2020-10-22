![GitHub Workflow Status](https://img.shields.io/github/workflow/status/alexislours/youtube-ban/Build%20releases) 
![GitHub release (latest by date)](https://img.shields.io/github/v/release/alexislours/youtube-ban) 
![GitHub](https://img.shields.io/github/license/alexislours/youtube-ban)

# youtube-ban

A quick app that generates a video that, once uploaded to YouTube, will terminate the account of the uploader in under 5 minutes. Be it an unlisted, private or public video.

## But why?

This simply serve to prove that the YouTube spam detection algorithm can lead to many false positives, that the review staff is totally unqualified and that you might as well flip a coin for the outcome of a false positive. 

I lost an 11 years old YouTube account after uploading a video similar to this one, with a denied appeal by the very qualified reviewers. The second account I uploaded that on was reinstated after a week of wait but I'm still waiting on my first account to be unbanned after 6 month. 

Get ready to play Russian roulette if your account gets terminated for whatever reason.

## Usage

Build the project by cloning the repository and running `go build` inside it.

Generate the video by running `./youtube-ban`.

Upload the generated `out.mp4` to YouTube.

Prebuilt binaries for Windows, Linux and macOS can be found [here](https://github.com/alexislours/youtube-ban/releases/latest).