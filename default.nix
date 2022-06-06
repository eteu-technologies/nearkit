{ lib, buildGoModule, rev ? null }:

buildGoModule rec {
  pname = "nearkit";
  version = if (rev != null) then (builtins.substring 0 7 rev) else "0000000";

  src = lib.cleanSource ./.;

  vendorSha256 = "sha256-tGW6HvZPfzx4sQugRlq3d3cBMNMr0ghuqgJo+xSEOxE=";
  subPackages = [ "cmd/nearkit" ];
}
