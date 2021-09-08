{ lib, buildGoModule, rev ? null }:

buildGoModule rec {
  pname = "nearkit";
  version = if (rev != null) then (builtins.substring 0 7 rev) else "0000000";

  src = lib.cleanSource ./.;

  vendorSha256 = "sha256-s2g1stDCtj66KFs6lnZPwPTXFU1Mcn7vf8XEnUXlPyQ=";
  subPackages = [ "cmd/nearkit" ];
}
