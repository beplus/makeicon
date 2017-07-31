class Makeicon < Formula
  desc "Generate mobile app icons in all resolutions for both iOS and Android"
  homepage "https://github.com/beplus/makeicon"
  url "https://github.com/beplus/makeicon/releases/download/v0.0.26/makeicon_0.0.26_darwin_amd64.tar.gz"
  version "0.0.26"
  sha256 "a39c5d7d68b86e2d4454677af3487292a30ed3edf5fd883fac7c3b50255fe47e"

  def install
    bin.install "makeicon"
  end

  test do
    
  end
end
