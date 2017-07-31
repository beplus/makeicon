class Makeicon < Formula
  desc "Generate mobile app icons in all resolutions for both iOS and Android"
  homepage "https://github.com/beplus/makeicon"
  url "https://github.com/beplus/makeicon/releases/download/v0.0.27/makeicon_0.0.27_darwin_amd64.tar.gz"
  version "0.0.27"
  sha256 "c929c24d5a9ada71e45361db483ed0daacdd3e559ac17103a599881e8bc36dea"

  def install
    bin.install "makeicon"
  end

  test do
    
  end
end
