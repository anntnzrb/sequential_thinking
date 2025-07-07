{
  description = "Go development environment";

  inputs = {
    nixpkgs.url = "https://channels.nixos.org/nixos-25.05/nixexprs.tar.xz";
    flake-parts.url = "github:hercules-ci/flake-parts/main";
    flake-parts.inputs.nixpkgs-lib.follows = "nixpkgs";
  };

  outputs =
    inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = inputs.nixpkgs.lib.systems.flakeExposed;

      perSystem =
        { pkgs, ... }:
        let
          buildGoModule = pkgs.buildGoModule.override { go = pkgs.go; };
          buildWithSpecificGo = pkg: pkg.override { inherit buildGoModule; };
        in
        {
          devShells.default = pkgs.mkShell {
            packages = [
              pkgs.go
              (buildWithSpecificGo pkgs.gopls)
            ];
          };
        };
    };
}
