{
  description = "GoStatusBar";
  outputs = {self, nixpkgs, ...}:

  {
    defaultPakage = self.statusbar;
    statusbar = nixpkgs.pkgs.buildGoModule {
      pname = "statusbar";
      version = "0.0.2";
      src = builtins.fetchGit {
        url = "https://github.com/mamaart/dwm-status.git";
        ref = "main";
        rev = "746e9dff3bcf39e12f4ed399f9c3199be1ed9d9f";
      };
      vendorHash = "sha256-bZ8BbYgebatTQh4KVv2J0hBLwPuOHZaQAQX3o63R4HU=";
    };
  };
}
