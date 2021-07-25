{ lib, buildGoModule, rev ? null }:

buildGoModule rec {
  pname = "nearkit";
  version = if (rev != null) then (builtins.substring 0 7 rev) else "0000000";

  src = lib.cleanSource ./.;

  vendorSha256 = "sha256-V/gOoskUT0mD2vZLuGN7uMHvR8Q5bvA/QnQtn+SEOwE=";
  subPackages = [ "cmd/nearkit" ];
}
