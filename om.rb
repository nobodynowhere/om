# This file was generated by GoReleaser. DO NOT EDIT.
class Om < Formula
  desc ""
  homepage ""
  version "6.3.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal-cf/om/releases/download/6.3.0/om-darwin-6.3.0.tar.gz"
    sha256 "322a1d8d5014c43d09c735d273b03d501f2d11559d0393cc2e7000d94fb3a51d"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/om/releases/download/6.3.0/om-linux-6.3.0.tar.gz"
      sha256 "fc9a49ba986a11c61836218cd36824faa5155449a5552f171ea0275f3a5c1fa6"
    end
  end

  def install
    bin.install "om"
  end

  test do
    system "#{bin}/om --version"
  end
end
