{
  description = "A flake that installs youtube-dl";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = [
          pkgs.cobra-cli
          pkgs.go
          pkgs.yt-dlp
        ];

        apps.default = {
          type = "app";
          program = "${pkgs.yt-dlp}/bin/yt-dlp";
        };

        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.cobra-cli
            pkgs.go
            pkgs.yt-dlp
          ];
        };
      });
}
