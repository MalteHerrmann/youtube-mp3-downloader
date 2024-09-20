# Youtube MP3 Downloader

This repository contains my personal setup to download mp3s
of the highest available quality through the CLI.

As a test for me to get used to the Nix ways,
it's using a Nix flake to organize the execution in a controlled environment.

## Usage

To run the tool, just enter the Nix shell:

```bash
nix develop
```

This will build and start the shell in the Nix environment,
which has `yt-dlp` installed.
To download a given track in MP3 format, run:

```bash
yt-dlp -x --audio-format mp3 [URL]
```

