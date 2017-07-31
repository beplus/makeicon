class Makeicon < Formula
  desc "Tool for generating icons in all resolutions for both iOS and Android applications"
  homepage "https://github.com/beplus/makeicon"
  url "https://github.com/beplus/makeicon.git",
      :tag => "v0.0.21",
      :revision => "8bf480cd227b69e0cccd6758f26eb0e0422aed8a"
  sha256 "26d5c008d8433faee0d600cd80076b968697369cf9836e0adab9b882e526c49f"

  head "https://github.com/beplus/makeicon.git",
       :shallow => false

  depends_on "go" => :build
  depends_on "govendor" => :build

  def install
    contents = Dir["{*,.git,.gitignore}"]
    gopath = buildpath/"gopath"
    (gopath/"src/github.com/beplus/makeicon").install contents

    ENV["GOPATH"] = gopath
    ENV.prepend_create_path "PATH", gopath/"bin"

    cd gopath/"src/github.com/beplus/makeicon" do
      system "govendor", "sync"
      system "go", "build", "-o", "bin/makeicon"
      bin.install "bin/makeicon"
    end
  end

  test do
    system "#{bin}/makeicon", "--version"
  end
end
