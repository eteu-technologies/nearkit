{ lib, buildGoModule, rev ? null }:

buildGoModule rec {
  pname = "nearkit";
  version = if (rev != null) then (builtins.substring 0 7 rev) else "0000000";

  src = lib.cleanSource ./.;

  vendorSha256 = "sha256-PnwprAOOCHX4JMHucW0XLnUFSflwHuXr6qHA2oL8GPY=";
  subPackages = [ "cmd/nearkit" ];
}
