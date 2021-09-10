{ lib, buildGoModule, rev ? null }:

buildGoModule rec {
  pname = "nearkit";
  version = if (rev != null) then (builtins.substring 0 7 rev) else "0000000";

  src = lib.cleanSource ./.;

  vendorSha256 = "sha256-gmbdbmMr5Thr1JqTA8OuZKl8U8oWtbzqTk3XbBLBuXs=";
  subPackages = [ "cmd/nearkit" ];
}
