{
  description = "DWM statusbar service";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {self, nixpkgs, flake-utils, ...}:
  flake-utils.lib.eachDefaultSystem (system: {
    packages.default = nixpkgs.legacyPackages.${system}.buildGoModule {
      pname = "dwm-status";
      version = "0.0.1";
      src = ./.;
      vendorHash = "";
    };

    apps.default = {
      type = "app";
      program = "${self.packages.${system}.default}/bin/dwm-status";
    };
  }) // {
    nixosModules.default = {config, lib, pkgs, ...}: {
      options.services.dwm-status = {
        enable = lib.mkEnableOption "Enable dwm-statusbar";
      };

      config = lib.mkIf config.services.dwm-status.enable {
        systemd.user.services.dwm-status = {
          description = "DWM statusbar";
          wantedBy = ["default.target"];
          after = ["graphical-session.target"];
          serviceConfig = {
            ExecStart = "${self.packages.${pkgs.system}.default}/bin/statusbar";
            Restart = "always";
            Type = "simple";
          };
          environment = {
          };
        };
      };
    };
  };
}

