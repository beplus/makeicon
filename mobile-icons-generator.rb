class MobileIconsGenerator < Formula
  desc ""
  homepage ""
  version "0.0.6"
  url "https://github.com/beplus/mobile-icons-generator/releases/download/v0.0.6/mobile-icons-generator_0.0.6_macOS_64-bit.tar.gz"
  sha256 "f4bac7a996b3c70977984b206bdb6260523420e86e520e51c8a7b682cdc6e07c"

  def install
    bin.install "makeicon"
  end

  test do
    
  end
end
