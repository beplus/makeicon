class MobileIconsGenerator < Formula
  desc ""
  homepage ""
  url "https://github.com/beplus/mobile-icons-generator/releases/download/v0.0.4/mobile-icons-generator_0.0.4_macOS_64-bit.tar.gz"
  version "0.0.4"
  sha256 "bc15bc5488401f5b18dea9ab84b0dca736a5f6fd06dca70dbe9e22696e696558"

  depends_on "git"depends_on "zsh"
  conflicts_with "svn"

  def install
    bin.install "program"
  end

  test do
    
  end
end
